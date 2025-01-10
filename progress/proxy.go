package progress

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/types"
)

type Proxy struct {
	lastUpdate time.Time
	startTime  time.Time

	ctx    *ext.Context
	update *ext.Update
	sent   *types.Message

	pxCtx context.Context
	pxUid string

	rw *os.File

	sizeStr   string
	upCurrent float64
	current   float64
	size      float64
}

func NewProxy(ctx *ext.Context, update *ext.Update, sent *types.Message, size float64, pxCtx context.Context, pxUid string) (*Proxy, error) {
	tmpfile, err := os.CreateTemp(os.TempDir(), "urluploaded")
	if err != nil {
		return nil, err
	}

	if size > MaxFileSize {
		return nil, ErrFileTooBig
	}

	sizeStr := fmt.Sprintf("%.2f", size/FloatMB)
	if sizeStr == "-0.00" {
		sizeStr = "?"
	}

	rn := time.Now()
	return &Proxy{
		ctx:        ctx,
		pxCtx:      pxCtx,
		pxUid:      pxUid,
		update:     update,
		sent:       sent,
		size:       size,
		sizeStr:    sizeStr,
		rw:         tmpfile,
		lastUpdate: rn,
		startTime:  rn,
	}, nil
}

// PrepareToUpload seeks the underlying fd to the beginning of the file to allow uploading.
// it also resets the times used for progress and sets the size in case it was unknown while downloading.
func (px *Proxy) PrepareToUpload() error {
	if _, err := px.rw.Seek(0, 0); err != nil { // idk why this would fail but OK!
		return err
	}

	rn := time.Now()
	px.lastUpdate = rn
	px.startTime = rn

	px.size = px.current // `current` was last set when we finished downloading, so we can use it for size in case we didnt know it when downloading
	px.sizeStr = fmt.Sprintf("%.2f", px.size/FloatMB)
	return nil
}

func (px *Proxy) DeleteTemp() {
	if err := os.Remove(px.rw.Name()); err != nil {
		panic("failed to delete tempfile: " + err.Error())
	}
	log.Println("deleted", px.rw.Name())
}
