package progress

import (
	"fmt"
	"strings"
	"time"

	"github.com/gotd/td/tg"
)

// Read proxies calls to the underlying file's Read method used in the uploading stage.
// it will check the time on every call and edit the message with progress as necessary.
func (px *Proxy) Read(p []byte) (n int, err error) {
	n, err = px.rw.Read(p)
	if err != nil {
		return
	}

	px.upCurrent += float64(n)

	// we always know what the size of the file is (when uploading), so no need to check if we should calculate %
	if rn := time.Now(); px.lastUpdate.Add(ThreeSeconds).Compare(rn) == -1 && px.upCurrent != px.size {
		b := strings.Builder{}
		b.WriteString(fmt.Sprintf("uploading - %.2f / %sMB - %.2f%%", px.upCurrent/FloatMB, px.sizeStr, px.upCurrent/px.size*100))

		bytesPerSec := px.upCurrent / float64(rn.Unix()-px.startTime.Unix())
		b.WriteString(fmt.Sprintf("\nspeed: ~%.2fMB/s", bytesPerSec/FloatMB))
		b.WriteString(fmt.Sprintf("\neta: ~%.2fs", (px.size-px.upCurrent)/bytesPerSec))

		px.ctx.EditMessage(px.update.EffectiveChat().GetID(), &tg.MessagesEditMessageRequest{
			ID:      px.sent.ID,
			Message: b.String(),
		})
		px.lastUpdate = rn
	}

	return
}
