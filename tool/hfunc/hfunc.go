package main

import (
	"errors"
	"github.com/hnit-acm/hfunc/basic"
	"github.com/urfave/cli/v2"
	"os"
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
		},
	}
	app.Run(os.Args)
	//
	//flag.Parse()
	//if !flag.Parsed() {
	//	return
	//}
	//args := flag.Args()
	//fmt.Println(args)
	//argsString := utils.ArrayStringToString(args, " ")
	////expNewService, _ := regexp.Compile(`^new \S+$`)
	////expNewService, _ := regexp.Compile(`^new \S+$`)
	//switch {
	//case expNewService.MatchString(argsString):
	//	{
	//		fmt.Println("new service ", args[1])
	//		newService(basic.String(args[1]))
	//		return
	//	}
	//}
}
