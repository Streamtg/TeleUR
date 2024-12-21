package main

import (
	"mime"
	"path/filepath"
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
