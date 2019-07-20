package rndimage

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/svlapin/rndimage/pkg/rndimage"
	"github.com/urfave/cli"
)

// RunCli runs rndimage CLI tool
func RunCli() {
	app := cli.NewApp()
	app.Name = rndimage.Name
	app.Usage = "generate jpeg images from random pixels"
	app.Version = rndimage.Version
	app.Author = rndimage.Author

	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"gen", "g"},
			Usage:   "generate an image of a given size",
			Action: func(c *cli.Context) error {
				if !c.Args().Present() {
					return fmt.Errorf("not enough arguments (\"rndimage help\" for help)")
				}

				widthStr := c.Args().Get(0)
				heightStr := c.Args().Get(1)

				if widthStr == "" {
					return fmt.Errorf("missing width")
				}
				if heightStr == "" {
					return fmt.Errorf("missing height")
				}
				width, err := strconv.Atoi(widthStr)
				if err != nil {
					return fmt.Errorf("could not use %v as width", widthStr)
				}

				height, err := strconv.Atoi(c.Args().Get(1))
				if err != nil {
					return fmt.Errorf("could not use %v as height", heightStr)
				}
				return rndimage.Encode(os.Stdout, rndimage.Generate(width, height))
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
