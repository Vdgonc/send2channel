package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vdgonc/send2channel/pkg/config"
	"github.com/vdgonc/send2channel/pkg/slack"
)

func main() {
	cf := config.New()

	appCmd := flag.String("appname", "", "The application name")
	statusCmd := flag.String("status", "init", "The deploy status")

	flag.Parse()

	if *appCmd == "" {
		os.Exit(1)
	}

	ch := make(chan string)

	for _, cn := range cf.Channels {
		go slack.SendMessage(cn, *appCmd, *statusCmd, ch)
	}

	fmt.Println(<-ch)

}
