package main

import (
	"bytes"
	"errors"
	"mime"
	"path/filepath"

	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/types"
	"github.com/gotd/td/telegram/uploader"
	"github.com/gotd/td/tg"
)

func GetFilename(link, disposition, ctype string) string {
	_, params, _ := mime.ParseMediaType(disposition)

	if filename, ok := params["filename"]; ok {
		return filename
	}

	filename := filepath.Base(link)
	if filename == "" {
		exts, _ := mime.ExtensionsByType(ctype)
		if len(exts) != 0 {
			filename = "file" + exts[0]
		} else {
			filename = "file.bin"
		}
	}

	return filename
}

func GetMessageReplyMedia(ctx *ext.Context, msg *types.Message) (tg.InputFileClass, error) {
	rt, ok := msg.ReplyTo.(*tg.MessageReplyHeader)
	if !ok {
		return nil, errors.New("reply_to is not MessageReplyHeader")
	}

	msgs, err := ctx.GetMessages(msg.PeerID.(*tg.PeerUser).UserID, []tg.InputMessageClass{&tg.InputMessageID{ID: rt.ReplyToMsgID}})
	if err != nil {
		return nil, err
	}

	icon := &bytes.Buffer{}
	if _, err := ctx.DownloadMedia(msgs[0].(*tg.Message).Media, ext.DownloadOutputStream{Writer: icon}, nil); err != nil {
		return nil, err
	}

	// thumbnails can't be reused and must be reuploaded
	file, err := uploader.NewUploader(ctx.Raw).FromBytes(ctx, "icon.jpg", icon.Bytes())
	if err != nil {
		return nil, err
	}

	return file, nil
}
