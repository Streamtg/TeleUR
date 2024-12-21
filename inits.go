package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// get env file name from args and parse
func init() {
	var envFile string
	flag.StringVar(&envFile, "env", ".env", "the path to the env file to use")
	flag.Parse()

	if err := godotenv.Load(envFile); err != nil {
		log.Fatalln(err)
	}
}

// load the env into our globals
func init() {
	apiID, err := strconv.Atoi(os.Getenv("API_ID"))
	if err != nil {
		log.Fatalln("couldn't load api id:", err)
	}

	ApiID = apiID
	ApiHash = os.Getenv("API_HASH")
	BotToken = os.Getenv("BOT_TOKEN")
	SessionPath = os.Getenv("SESSION_PATH")

	authSpl := strings.Split(os.Getenv("AUTHORIZED"), ",")
	authIds := make([]int64, 0, len(authSpl))

	for _, s := range authSpl {
		p, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalln("couldn't parse user id:", s)
		}
		authIds = append(authIds, p)
	}

	Authorized = authIds
}
