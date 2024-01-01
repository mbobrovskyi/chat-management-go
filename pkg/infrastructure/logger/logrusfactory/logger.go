package logrusfactory

import (
	"github.com/mbobrovskyi/chat-management-go/pkg/baseconfig"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/logger"
	"github.com/sirupsen/logrus"
)

func NewLogger(lvl logger.Level) (*logrus.Logger, error) {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return nil, err
	}

	log := logrus.New()
	log.SetLevel(level)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: baseconfig.DateTimeFormat,
	})

	return log, nil
}
