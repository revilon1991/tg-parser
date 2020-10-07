package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/revilon1991/tg-parser/internal/client"
	"github.com/revilon1991/tg-parser/internal/consumer/update"
	"github.com/revilon1991/tg-parser/internal/controller"
	"github.com/revilon1991/tg-parser/internal/useCase/fetchChannelMembersInfo"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
)

var (
	Version  = "0"
	CommitID = "0"
	commands = []cli.Command{
		{
			Name:   "run-server",
			Usage:  "Run server with api end-points",
			Action: api,
		},
		{
			Name:   "fetch-members",
			Usage:  "Fetch and store channel members info",
			Action: fetchChannelMembersInfo.Handle,
		},
	}
)

func main() {
	_ = godotenv.Load()
	app := cli.NewApp()
	app.Name = "Telegram Parser"
	app.Usage = "Parsing telegram channels and users"
	app.Version = fmt.Sprintf("%s - %s", Version, CommitID)

	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error %v", err)
	}
}

func api(c *cli.Context) {
	log.Println("starting server...")

	clientStorage := client.NewClient()

	update.Handle(clientStorage)

	log.SetFlags(log.LstdFlags | log.Llongfile)

	controller.GetMeAction(clientStorage)
	controller.GetMembersAction(clientStorage)
	controller.GetUserAction(clientStorage)
	controller.GetPhotoAction(clientStorage)
	controller.GetChannelInfoAction(clientStorage)
	controller.GetChannelAction(clientStorage)
	controller.Proxy(clientStorage)

	controller.GetStorageChannelList()
	controller.GetStorageMemberList()

	err := http.ListenAndServe(":"+client.Port, nil)

	if err != nil {
		log.Fatal(fmt.Sprintf("Run server error: %s", err.Error()))
	}
}
