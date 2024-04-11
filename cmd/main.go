package main

import (
	"db-exec-layer/config"
	"db-exec-layer/pkg/log"
	"fmt"

	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	RunCli()
}

const (
	fileFlag      = "configfile"
	fileShortFlag = "c"
)

func RunCli() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     fileFlag,
				Aliases:  []string{fileShortFlag},
				Usage:    "file path",
				Required: true,
			},
		},
		Action: Start,
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Errorf("run error %v", err)
		return
	}
}

func Start(c *cli.Context) error {

	if c.String(fileFlag) == "" {
		fmt.Println("param error")
		os.Exit(1)
	}

	cfg := &config.DbLayerConfig{}
	err := config.ResolveConfig(c.String(fileFlag), cfg)
	if err != nil {
		fmt.Println("config init error")
		return err
	}

	log.InitLog(cfg.Log.LogFilePath, cfg.Log.LogName, cfg.Log.LogLevel)

	// dbPool, err := db.NewDBEngine(&cfg.DatabaseConfig)
	// if err != nil {
	// 	return err
	// }

	// ctx, cancel := context.WithCancel(context.Background())
	// // wait for signals
	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// go func() {
	// 	<-sigs
	// 	cancel()
	// }()

	return err
}
