# TeleURLUploader
the most demure url uploader written in go

works nice, shows progress, eta, and speed, and overall works very nicely with nice libraries

[![Go Reference](https://pkg.go.dev/badge/github.com/asdfzxcvbn/TeleURLUploader.svg)](https://pkg.go.dev/github.com/asdfzxcvbn/TeleURLUploader)

# building
there are prebuilt binaries in releases. but you can follow the steps below to build it yourself:

1. install [go](https://go.dev)
2. clone the repo and cd
3. `go get`
4. `go build -ldflags="-s -w"`

# usage
```bash
$ ./TeleURLUploader -help
Usage of ./TeleURLUploader:
  -env string
    	the path to the env file to use (default ".env")
```

there's a `.env.sample` file you can fill out. the only key that might require explanation is `AUTHORIZED`, it's a comma-seperated list of user ids that will be able to use the bot. you can also make `SESSION_PATH` a full path in case you'll be moving the binary around.
