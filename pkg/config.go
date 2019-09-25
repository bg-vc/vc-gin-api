package pkg

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

var (
	Log         *logrus.Logger
	LogFileName string
)

func InitLogger() {
	LogFileName = "logs/vc-gin-api.log"
	LogFileName, err := filepath.Abs(LogFileName)
	if err != nil {
		panic(err)
	}

	newLoggerFile(LogFileName)

	src, err := os.OpenFile(LogFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic("InitLogger error")
	}

	Log = logrus.New()
	Log.Out = src
}

func newLoggerFile(LoggerFile string) {
	localPath := ""
	outputPath := LoggerFile
	LogFileName = strings.Replace(filepath.Join(localPath, outputPath), "\\", "/", -1)
	fmt.Printf("1212:%v\n", LogFileName)
	_, err := os.Stat(strings.Replace(filepath.Dir(filepath.Join(localPath, outputPath)), "\\", "/", -1))
	if err != nil && os.IsNotExist(err) {
		fmt.Printf("123123:%v\n", err)
		os.MkdirAll(strings.Replace(filepath.Dir(filepath.Join(localPath, outputPath)), "\\", "/", -1), os.ModePerm)
	} else if err != nil {
		panic(err)
	}
	fmt.Printf("232323:%v\n", strings.Replace(filepath.Join(localPath, outputPath), "\\", "/", -1))
	f, err := os.OpenFile(strings.Replace(filepath.Join(localPath, outputPath), "\\", "/", -1), os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("4444:%v\n", err)
		panic(err)
	}
	defer f.Close()
}