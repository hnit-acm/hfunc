package main

import (
	"errors"
	"fmt"
	"github.com/hnit-acm/hfunc/hbasic"
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
				Usage:   "new a services or request",
				Subcommands: []*cli.Command{
					{
						Name:    "services",
						Aliases: []string{"s"},
						Usage:   "",
						Flags: []cli.Flag{
							//&cli.StringFlag{
							//	Name: "name",
							//	Aliases: []string{
							//		"n",
							//	},
							//	Usage: "services name",
							//},
							&cli.StringFlag{
								Name: "port",
								Aliases: []string{
									"p",
								},
								Value: "8000",
								Usage: "port number",
							},
						},
						Action: func(c *cli.Context) error {
							//fmt.Println(c.String("name"))
							//fmt.Println(c.String("port"))

							//fmt.Println(c.NArg())
							//fmt.Println(c.Args().Tail())
							//fmt.Println(c.Args().First())
							serviceName := c.Args().Get(c.NArg() - 1)
							port := c.String("port")
							switch serviceName {
							case "":
								{
									return errors.New("services name is not empty")
								}
							default:
								if newService(hbasic.String(serviceName), hbasic.String(port)) {
									return nil
								} else {
									return errors.New("new services failed")
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
									return errors.New("services name is not empty")
								}
							default:
								// todo
								if newService(hbasic.String(service_name), hbasic.String(service_name)) {
									return nil
								} else {
									return errors.New("new services failed")
								}
							}
						},
					},
				},
			},
			{
				Name:    "sync",
				Aliases: nil,
				Usage:   "sync config directory between services and services",
				Action:  nil,
			},
			{
				Name:    "swag",
				Aliases: nil,
				Usage:   "start an apih doc serverh",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "port",
						Aliases: []string{
							"p",
						}},
					&cli.StringFlag{
						Name: "rewrite",
						Aliases: []string{
							"r",
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
					rewrite := c.String("rewrite")

					return InitSwag(path+"/swagger.json", port, rewrite)
				},
			},
			//{
			//	Name: "redis",
			//	Aliases: []string{
			//		"r",
			//	},
			//	Usage: "start redis instance or cluster",
			//	Flags: []cli.Flag{
			//		&cli.StringFlag{
			//			Name:        "instance",
			//			Aliases: []string{
			//				"i",
			//			},
			//		},
			//		&cli.StringFlag{
			//			Name:        "cluster",
			//			Aliases: []string{
			//				"c",
			//			},
			//		},
			//	},
			//	Action: func(ctx *cli.Context) error {
			//		fmt.Printf("hfunc version %v %v/%v\n", Version, runtime.GOOS, runtime.GOARCH)
			//		return nil
			//	},
			//},
			//{
			//	Name: "doctor",
			//	Aliases: []string{
			//		"d",
			//	},
			//	Usage: "check environment",
			//	Action: func(ctx *cli.Context) error {
			//		fmt.Printf("hfunc version %v %v/%v\n", Version, runtime.GOOS, runtime.GOARCH)
			//		return nil
			//	},
			//},
			{
				Name: "version",
				Aliases: []string{
					"v",
				},
				Usage: "current version",
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
