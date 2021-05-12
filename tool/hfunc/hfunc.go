package main

import (
	"errors"
	"fmt"
	"github.com/hnit-acm/hfunc/basic"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "new",
				Aliases: nil,
				Usage:   "new a service or request",
				Subcommands: []*cli.Command{
					{
						Name:    "service",
						Aliases: []string{"s"},
						Usage:   "",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "name",
								Aliases: []string{
									"n",
								},
								Usage: "service name",
							},
						},
						Action: func(c *cli.Context) error {
							serviceName := c.Args().Get(0)
							switch serviceName {
							case "":
								{
									return errors.New("service name is not empty")
								}
							default:
								if newService(basic.String(serviceName)) {
									return nil
								} else {
									return errors.New("new service failed")
								}
							}
						},
					},
					{
						Name:    "req",
						Aliases: []string{"r"},
						Usage:   "",
						Action: func(c *cli.Context) error {
							service_name := c.Args().Get(0)
							switch service_name {
							case "":
								{
									return errors.New("service name is not empty")
								}
							default:
								if newService(basic.String(service_name)) {
									return nil
								} else {
									return errors.New("new service failed")
								}
							}
						},
					},
				},
			},
			{
				Name:    "sync",
				Aliases: nil,
				Usage:   "sync config directory between service and service",
				Action:  nil,
			},
			{
				Name:    "swag",
				Aliases: nil,
				Usage:   "start an api doc server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "port",
						Aliases: []string{
							"p",
						}},
				},
				Action: func(c *cli.Context) error {
					path := c.Args().Get(0)
					if path == "" {
						path = "./docs"
					}
					port := c.String("port")
					if port == "" {
						port = "4000"
					}
					return InitSwag(path+"/swagger.json", port)
				},
			},
			{
				Name: "version",
				Aliases: []string{
					"v",
				},
				Usage: "sync config directory between service and service",
				Action: func(ctx *cli.Context) error {
					fmt.Printf("hfunc version %v %v/%v\n", Version, runtime.GOOS, runtime.GOARCH)
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		logh.Error(err)
		return
	}
}
