package download

import (
	"KHytaleLauncher/internal/download/extract"
	"KHytaleLauncher/internal/env"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func FromUrlAndExtract(
	extractDir, fileName,
	fileURL string) (string, error) {

	cacheDir := env.GetDefaultCacheDir()
	destPath := filepath.Join(cacheDir, fileName)

	filePath, err := FromUrl(destPath, fileURL, false)

	if strings.HasSuffix(fileName, ".zip") {
		return extractDir, extract.Unzip(destPath, extractDir)
	}

	return filePath, err
}

func FromUrl(
	fileDownloadPath,
	fileURL string,
	isOverride bool) (string, error) {

	// Client: NO overall timeout (for large downloads). Transport tuned.
	client := &http.Client{
		Timeout: 0, // no global timeout for large streams
		Transport: &http.Transport{
			DisableCompression:  true,
			ForceAttemptHTTP2:   false, // important per önceki tespit
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			DialContext:         (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}

	// Try HEAD to get size and Accept-Ranges, but don't fail download if HEAD fails.
	var serverFileSize int64 = 0
	var acceptRanges string
	headResp, err := client.Head(fileURL)
	if err == nil {
		if headResp.Header.Get("Content-Length") != "" {
			serverFileSize, _ = strconv.ParseInt(headResp.Header.Get("Content-Length"), 10, 64)
		}
		acceptRanges = headResp.Header.Get("Accept-Ranges")
		_ = headResp.Body.Close()
	} else {
		// fallback: continue without head info (some servers block HEAD). Log as fmt.Printf to avoid changing I/O.
		fmt.Printf("warning: HEAD failed, continuing with GET: %v\n", err)
		serverFileSize = 0
		acceptRanges = ""
	}

	if headResp.StatusCode == http.StatusNotFound {
		// 404
		return "", fmt.Errorf("file not found: %s", fileURL)
	}

	var startByte int64 = 0
	if !isOverride {
		if fi, err := os.Stat(fileDownloadPath); err == nil {
			startByte = fi.Size()
		}
	}

	if serverFileSize > 0 && startByte >= serverFileSize {
		fmt.Printf("File already fully downloaded: %s\n", fileDownloadPath)
		return fileDownloadPath, nil
	}

	req, err := http.NewRequest("GET", fileURL, nil)
	if err != nil {
		return "", err
	}

	// set Range if we have partial file; even if Accept-Ranges missing, we try — server may still honor it.
	if startByte > 0 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", startByte))
	}

	// simple retry loop for transient network errors
	var resp *http.Response
	retries := 3
	backoff := 1 * time.Second
	for attempt := 0; attempt < retries; attempt++ {
		resp, err = client.Do(req)
		if err == nil {
			break
		}
		// last attempt -> return error
		if attempt == retries-1 {
			return "", fmt.Errorf("failed to start download after %d attempts: %v", retries, err)
		}
		// retry for transient errors
		fmt.Printf("download attempt %d failed: %v — retrying in %s\n", attempt+1, err, backoff)
		time.Sleep(backoff)
		backoff *= 2
	}
	if resp == nil {
		return "", fmt.Errorf("no response received")
	}
	defer resp.Body.Close()

	// If HEAD didn't provide Content-Length, try to read from GET response.
	if serverFileSize == 0 {
		if resp.Header.Get("Content-Length") != "" {
			serverFileSize, _ = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		}
	}

	// Normalize acceptRanges check (case-insensitive)
	acceptRanges = strings.ToLower(acceptRanges)

	// Handle cases where we attempted a Range but server returned 200 (ignored Range).
	needTruncate := false
	if startByte > 0 {
		if resp.StatusCode == http.StatusPartialContent {
			// good, server supports range and will send remaining bytes
		} else if resp.StatusCode == http.StatusOK {
			// server ignored Range: we must restart download from 0 (truncate file)
			fmt.Println("server ignored Range header (returned 200). Restarting download from beginning.")
			needTruncate = true
			startByte = 0
		} else if resp.StatusCode == http.StatusRequestedRangeNotSatisfiable {
			// The requested range is not satisfiable — probably file on server is smaller
			return "", fmt.Errorf("requested range not satisfiable (server returned 416)")
		} else {
			// other status codes — treat non-2xx as error
			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				return "", fmt.Errorf("bad status: %s", resp.Status)
			}
		}
	} else {
		// no resume requested: ensure status is OK
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return "", fmt.Errorf("bad status: %s", resp.Status)
		}
	}

	// Open file for writing. Use SEEK/TRUNCATE instead of O_APPEND to control position.
	f, err := os.OpenFile(fileDownloadPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	// make sure file is closed later
	defer f.Close()

	if needTruncate {
		if err := f.Truncate(0); err != nil {
			return "", fmt.Errorf("failed to truncate file: %v", err)
		}
		if _, err := f.Seek(0, 0); err != nil {
			return "", fmt.Errorf("failed to seek: %v", err)
		}
	} else if startByte > 0 {
		if _, err := f.Seek(startByte, 0); err != nil {
			return "", fmt.Errorf("failed to seek to resume position: %v", err)
		}
	} else {
		// fresh download: ensure file is truncated
		if err := f.Truncate(0); err != nil {
			// non-fatal, but warn
			fmt.Printf("warning: truncate failed: %v\n", err)
		}
		if _, err := f.Seek(0, 0); err != nil {
			fmt.Printf("warning: seek failed: %v\n", err)
		}
	}

	// Update serverFileSize from response if still unknown
	if serverFileSize == 0 && resp.Header.Get("Content-Length") != "" {
		serverFileSize, _ = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	}

	counter := &WriteCounter{
		Total:    uint64(startByte),
		FileSize: uint64(serverFileSize),
	}

	// Use a large buffer to reduce syscall pressure and avoid micro-idle behavior
	buf := make([]byte, 1024*1024) // 1MB

	// Copy with buffer and TeeReader for progress counting
	_, err = io.CopyBuffer(f, io.TeeReader(resp.Body, counter), buf)
	if err != nil {
		return "", err
	}

	fmt.Printf("\nDownload file (%s) completed successfully!\n", fileDownloadPath)
	return fileDownloadPath, nil
}
