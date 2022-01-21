//go:build !windows

package main

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/xmapst/alertingwebhook"
    "github.com/xmapst/alertingwebhook/cmd"
    "github.com/xmapst/alertingwebhook/routers"
    "gopkg.in/alecthomas/kingpin.v2"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

// @title Alerting Webhook
// @version v1.0.0
// @description This is a os remote executor orchestration script interface.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	kingpin.Version(alertingwebhook.VersionIfo())
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	alertingwebhook.PrintHeadInfo()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	gin.SetMode(gin.ReleaseMode)
	if cmd.DebugEnabled {
		gin.SetMode(gin.DebugMode)
	}
	gin.DisableConsoleColor()
	srv := &http.Server{
		Addr:         cmd.ListenAddress,
		WriteTimeout: 600 * time.Second,
		ReadTimeout:  600 * time.Second,
		IdleTimeout:  600 * time.Second,
		Handler:      routers.Router(),
	}
	go func() {
		logrus.Infof("listen address [%s]", cmd.ListenAddress)
		if err := srv.ListenAndServe(); err != nil {
			logrus.Error(err)
		}
	}()

	<-signals
	logrus.Info("shutdown server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
