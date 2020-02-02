package utils

import "github.com/sirupsen/logrus"

// Log loger to be shared across
var Log = logrus.New()

func init() {
	Log.SetLevel(logrus.InfoLevel)
}
