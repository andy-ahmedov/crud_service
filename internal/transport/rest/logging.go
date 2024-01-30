package rest

import "github.com/sirupsen/logrus"

func logField(handler string, problem string) logrus.Fields {
	return logrus.Fields{
		"handler": handler,
		"problem": problem,
	}
}

func logError(handler string, problem string, err error) {
	logrus.WithFields(logField(handler, problem)).Error(err)
}
