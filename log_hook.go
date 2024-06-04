package securefilechanger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type FileLogHook struct{}

func (h *FileLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *FileLogHook) Fire(entry *logrus.Entry) error {
	funcName := strings.ReplaceAll(
		entry.Caller.Func.Name(),
		"github.com/KodokuOdius/SecureFileChanger/", "",
	)

	logFileDate := entry.Time.Format("20060102")
	logTime := entry.Time.Format("2006-01-02 15:04:05.27")
	fileMsg := fmt.Sprintf("[%s] %s (%s) = %s\n", entry.Level, logTime, funcName, entry.Message)
	folderPath := filepath.Join(os.Getenv("CLOUD_HOME"), "logs/")
	filePath := filepath.Join(folderPath, fmt.Sprintf("%s_companycloud.log", logFileDate))

	_ = os.MkdirAll(folderPath, os.ModePerm)

	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer logFile.Close()

	_, err = logFile.WriteString(fileMsg)
	if err != nil {
		return err
	}
	return nil
}
