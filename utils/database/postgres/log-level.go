package postgres

import "gorm.io/gorm/logger"

func DatabaseLogger(configLogLevel string) logger.LogLevel {
	var logLevel logger.LogLevel

	switch configLogLevel {
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	case "error":
		logLevel = logger.Error
	}

	return logLevel
}
