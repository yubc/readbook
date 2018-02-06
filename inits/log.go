package inits

import (
	"readbook/models"
	"readbook/utils/log"
)

func logInit() {
	var err error
	conf := &log.Configs{
		LogPath:       models.Conf.LogPath,
		LogLevel:      models.Conf.LogLevel,
		LogMaxSize:    models.Conf.LogMaxSize,
		LogBuffSize:   models.Conf.LogBuffSize,
		LogHook:       models.Conf.LogHook,
		EmailHost:     models.Conf.EmailHost,
		EmailPort:     models.Conf.EmailPort,
		EmailFrom:     models.Conf.EmailFrom,
		EmailTo:       models.Conf.EmailTo,
		EmailPassword: models.Conf.EmailPassword,
	}

	models.Log, err = log.LogFiring(conf, models.Conf.LogProject)
	if err != nil {
		panic(err)
	}
}
