package download

import (
	"KHytaleLauncher/internal/progress"
)

type WriteCounter struct {
	FileSize uint64
	Total    uint64
	Initial  uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()

	return n, nil
}

func (wc *WriteCounter) PrintProgress() {
	if wc.FileSize > 0 {
		percentage := float64(wc.Total) / float64(wc.FileSize) * 100

		progress.SetProgressPresent(int32(percentage))
	} else {
		progress.SetProgressPresent(int32(100))
	}
}
