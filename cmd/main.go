package main

import (
	"context"
	"db-exec-layer/config"
	"db-exec-layer/pkg/log"
	"db-exec-layer/protocol"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"os"

	"github.com/panjf2000/gnet/v2"
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

	_, cancel := context.WithCancel(context.Background())
	// wait for signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {

		<-sigs
		cancel()
	}()

	options := []gnet.Option{
		gnet.WithReusePort(cfg.ServerConfig.ReusePort),
		gnet.WithMulticore(cfg.ServerConfig.Multicore),
		gnet.WithTCPKeepAlive(time.Second * time.Duration(cfg.ServerConfig.TCPKeepAlive)),
	}

	tcpServer := protocol.NewTCPServer(cfg.ServerConfig.Port)

	err = protocol.Run(tcpServer, fmt.Sprintf("tcp://:%d", cfg.ServerConfig.Port), options...)
	if err != nil {
		logrus.Errorf("启动失败:%+v", err)
		return err
	}

	return err
}
