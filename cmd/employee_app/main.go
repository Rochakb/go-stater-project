package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	err := (&cli.App{
		Name:     "go_project",
		Usage:    "Go starter project crud apis",
		Version:  "1.0",
		Before:   badge,
		Flags:    Flags(),
		Commands: commands,
	}).Run(os.Args)
	if err != nil {
		fmt.Println("Something Went Wrong. Failed to start go_project.: " + err.Error())
		log.Fatal(
			fmt.Sprintf(
				"-- \nfailed to start go_project. \n--\n Caused By:\n%s\n--",
				errorstack(err.Error()),
			),
		)
	}
}
