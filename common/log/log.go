package log

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"taibai/common/conf"
	"taibai/consts"
	"github.com/t-tomalak/logrus-easy-formatter"
)

var Trace *logrus.Logger

func Init() error {
	config := conf.Get()

	file, err := os.OpenFile(config.Log.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return errors.New("failed to log to file")
	}

	Trace = &logrus.Logger{
		Out:   file,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: consts.TimeFormat,
			LogFormat:       "[%lvl%]: %time% - %msg%\r",
		},
	}


	return nil
}
