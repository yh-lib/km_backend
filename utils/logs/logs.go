package logs

import "github.com/sirupsen/logrus"

// 封装 logrus
func Debug(fileds map[string]interface{}, msg string) {
	logrus.WithFields(fileds).Debug(msg)
}

func Info(fileds map[string]interface{}, msg string) {
	logrus.WithFields(fileds).Info(msg)
}

func Warning(fileds map[string]interface{}, msg string) {
	logrus.WithFields(fileds).Warning(msg)
}

func Error(fileds map[string]interface{}, msg string) {
	logrus.WithFields(fileds).Error(msg)
}

func Fatal(fileds map[string]interface{}, msg string) {
	logrus.WithFields(fileds).Fatal(msg)
}

func Panic(fileds map[string]interface{}, msg string) {
	logrus.WithFields(fileds).Panic(msg)
}
