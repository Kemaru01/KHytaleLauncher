package extract

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(source, destination string) error {
	r, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer r.Close()

	commonPrefix := ""
	if len(r.File) > 0 {
		firstPath := r.File[0].Name
		parts := strings.Split(strings.Trim(firstPath, "/"), "/")

		if len(parts) > 1 || (len(parts) == 1 && r.File[0].FileInfo().IsDir()) {
			candidate := parts[0] + "/"
			isCommon := true

			for _, f := range r.File {
				if !strings.HasPrefix(f.Name, candidate) {
					isCommon = false
					break
				}
			}
			if isCommon {
				commonPrefix = candidate
			}
		}
	}

	for _, f := range r.File {
		relPath := strings.TrimPrefix(f.Name, commonPrefix)
		if relPath == "" {
			continue
		}

		fpath := filepath.Join(destination, relPath)

		if !strings.HasPrefix(fpath, filepath.Clean(destination)+string(os.PathSeparator)) {
			continue
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
