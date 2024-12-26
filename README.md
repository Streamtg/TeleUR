# TeleURLUploader
the most demure url uploader written in go

works nice, shows progress, eta, and speed, supports uploading the file with an icon, and overall works very nicely with nice libraries

[![Go Reference](https://pkg.go.dev/badge/github.com/asdfzxcvbn/TeleURLUploader.svg)](https://pkg.go.dev/github.com/asdfzxcvbn/TeleURLUploader)

# building
there are prebuilt binaries in releases. but you can follow the steps below to build it yourself:

1. install [go](https://go.dev)
2. clone the repo and cd
3. `go get`
4. `go build -ldflags="-s -w"`

# setup
```bash
$ ./TeleURLUploader -help
Usage of ./TeleURLUploader:
  -env string
    	the path to the env file to use (default ".env")
```

there's a `.env.sample` file you can fill out. the only key that might require explanation is `AUTHORIZED`, it's a comma-seperated list of user ids that will be able to use the bot. you can also make `SESSION_PATH` a full path in case you'll be moving the binary around.

# usage
after starting the bot, just send an http url to download. you can use the syntax `<link> <filename>` to upload the file with a custom filename and even upload the file with a custom icon by replying to a jpg image. example: `https://atl.speedtest.clouvider.net/1g.bin 1gigabyte.bin`

![normal download](/imgs/normal-download.png)
![normal upload](/imgs/normal-upload.png)
![normal received](/imgs/normal-received.png)

the bot can also intelligently download files from servers that dont report their Content-Length. in this case, the total file size, percentage, and eta will not be available during the download phase. all this information will become available when uploading, though.

![unknown download](/imgs/unknown-download.png)
![unknown upload](/imgs/unknown-upload.png)
![unknown received](/imgs/unknown-received.png)
