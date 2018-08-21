package logging

import (
	"Gin-Todo/pkg/setting"
	"fmt"
	"time"
)

func getLogFilePath() string {

	return fmt.Sprintf("%s%s", setting.RuntimeRootPath, setting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.LogSaveName,
		time.Now().Format(setting.TimeFormat),
		setting.LogFileExt,
	)
}
