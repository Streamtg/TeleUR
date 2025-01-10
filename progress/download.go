package progress

import (
	"fmt"
	"strings"
	"time"

	"github.com/gotd/td/tg"
)

// Write proxies calls to the underlying file's Write method used in the downloading stage.
// it will check the time on every call and edit the message with progress as necessary.
func (px *Proxy) Write(p []byte) (n int, err error) {
	if err = px.pxCtx.Err(); err != nil {
		return
	}

	n, err = px.rw.Write(p)
	if err != nil {
		return
	}

	px.current += float64(n)

	// in case content-length wasnt reported, we still need to stop at 2gb
	if px.current > MaxFileSize {
		return n, ErrFileTooBig
	}

	if rn := time.Now(); px.lastUpdate.Add(ThreeSeconds).Compare(rn) == -1 && px.current != px.size {
		b := strings.Builder{}
		b.WriteString(fmt.Sprintf("downloading - %.2f / %sMB", px.current/FloatMB, px.sizeStr))

		if px.sizeStr != "?" {
			b.WriteString(fmt.Sprintf(" - %.2f%%", px.current/px.size*100))
		}

		bytesPerSec := px.current / float64(rn.Unix()-px.startTime.Unix())
		b.WriteString(fmt.Sprintf("\nspeed: ~%.2fMB/s", bytesPerSec/FloatMB))

		if px.sizeStr != "?" {
			b.WriteString(fmt.Sprintf("\neta: ~%.2fs", (px.size-px.current)/bytesPerSec))
		}

		px.ctx.EditMessage(px.update.EffectiveChat().GetID(), &tg.MessagesEditMessageRequest{
			ID:      px.sent.ID,
			Message: b.String(),
			ReplyMarkup: &tg.ReplyInlineMarkup{
				Rows: []tg.KeyboardButtonRow{
					{
						Buttons: []tg.KeyboardButtonClass{
							&tg.KeyboardButtonCallback{
								Text: "cancel",
								Data: []byte("cancel-" + px.pxUid),
							},
						},
					},
				},
			},
		})
		px.lastUpdate = rn
	}

	return
}
