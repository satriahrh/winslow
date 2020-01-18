package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satriahrh/winslow/commands"
	"github.com/satriahrh/winslow/ent"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func init() {
	if err := os.Chdir(os.Getenv("HOME") + "/.winslow"); err != nil {
		if err := os.Chdir(os.Getenv("HOME")); err != nil {
			log.Fatalln(err)
		}
		if err := os.Mkdir(".winslow", 755); err != nil {
			log.Fatalln(err)
		}
		if err := os.Chdir(os.Getenv("HOME") + "/.winslow"); err != nil {
			log.Fatalln(err)
		}
	}

	f, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(f)
}

func main() {
	client, err := ent.Open("mysql", os.Getenv("WINSLOW_DSN"))
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed opening connection to db\n%v", err))
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalln(fmt.Sprintf("failed creating schema resources: %v", err))
	}

	app := &cli.App{
		Name:  "winslow",
		Usage: "Better project management",
		Action: func(c *cli.Context) error {
			fmt.Println("The name is Frederick Winslow Taylor.")
			return nil
		},
		Commands: commands.NewCommands(client),
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
