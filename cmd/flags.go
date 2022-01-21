package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/xmapst/alertingwebhook/utils"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	DebugEnabled  bool
	ListenAddress string
	RoboterMap    = map[string]string{}
)

func init() {
	// flags
	kingpin.Flag(
		"addr",
		`host:port for execution.`,
	).Default(":8888").StringVar(&ListenAddress)
	kingpin.Flag(
		"debug",
		`Enable debug messages`,
	).Default("false").BoolVar(&DebugEnabled)
	kingpin.Flag(
		"robot",
		"robot token sec key. Example: --robot xxx=xxx --robot aaa=bbbb",
	).Required().StringMapVar(&RoboterMap)
	// log format init
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&utils.FileFormatter{})
}
