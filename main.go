package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Address string `toml:"address"`
	UseTLS  bool   `toml:"use_tls"`
	Cert    struct {
		CertFile string `toml:"cert_file"`
		CertKey  string `toml:"cert_key"`
	} `toml:"cert"`
}

func main() {
	var c Config
	// configの読み込み
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.GET("/", func(e echo.Context) error {
		log.Info("あちほい")
		response := map[string]string{
			"ip":   e.RealIP(),
			"date": time.Now().Format("2006-01-02 15:04:05"),
			"utc":  time.Now().UTC().Format("2006-01-02 15:04:05"),
		}
		return e.JSON(http.StatusOK, response)
	})

	// 停止処理の検出
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// apiサーバーの起動
	go func() {
		if c.UseTLS {
			if err := e.StartTLS(c.Address, c.Cert.CertFile, c.Cert.CertKey); err != http.ErrServerClosed {
				panic(err)
			}
		} else {
			if err := e.Start(c.Address); err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	// 停止処理
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
