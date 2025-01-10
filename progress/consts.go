package progress

import (
	"errors"
	"time"
)

const (
	MB           = 1024 * 1024
	FloatMB      = float64(MB)
	MaxFileSize  = 2097152000 // 4000 parts of 512 KB each (4000 * 512 * 1024)
	ThreeSeconds = 3 * time.Second
)

var (
	ErrFileTooBig = errors.New("file size exceeds 2gb")
)
