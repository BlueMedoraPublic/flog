package main

import (
	"compress/gzip"
	"io"
	"os"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"math/rand"
)

func Worker(option *Option, writer io.WriteCloser) {
	x := os.Getenv("MAX_SLEEP")
	maxSleep := 30
	if x != "" {
		var err error
		if maxSleep, err = strconv.Atoi(x); err != nil {
			fmt.Println("MAX_SLEEP: " + x + " is not valid")
			os.Exit(1)
		}
	}
	fmt.Println("using max sleep: " + strconv.Itoa(maxSleep))


	var loc *time.Location
	var err error
	loc, err = time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for {

		created	:= time.Now().In(loc)

		log := NewLog(option.Format, created)
		_, _ = writer.Write([]byte(log + "\n"))


		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(maxSleep) // n will be between 0 and 10
		time.Sleep(time.Duration(n)*time.Second)
	}
}

// Generate generates the logs with given options
func Generate(option *Option) error {


	logFileName := option.Output
	writer, err := NewWriter(option.Type, logFileName)
	if err != nil {
		return err
	}

	go Worker(option, writer)
	go Worker(option, writer)
	go Worker(option, writer)
	go Worker(option, writer)

	for {
		time.Sleep(time.Hour * 1)
	}
	return nil
}

// NewWriter returns a closeable writer corresponding to given log type
func NewWriter(logType string, logFileName string) (io.WriteCloser, error) {
	switch logType {
	case "stdout":
		return os.Stdout, nil
	case "log":
		logFile, err := os.Create(logFileName)
		if err != nil {
			return nil, err
		}
		return logFile, nil
	case "gz":
		logFile, err := os.Create(logFileName)
		if err != nil {
			return nil, err
		}
		return gzip.NewWriter(logFile), nil
	default:
		return nil, nil
	}
}

// NewLog creates a log for given format
func NewLog(format string, t time.Time) string {
	switch format {
	case "apache_common":
		return NewApacheCommonLog(t)
	case "apache_combined":
		return NewApacheCombinedLog(t)
	case "apache_error":
		return NewApacheErrorLog(t)
	case "rfc3164":
		return NewRFC3164Log(t)
	case "rfc5424":
		return NewRFC5424Log(t)
	case "common_log":
		return NewCommonLogFormat(t)
	case "json":
		return NewJSONLogFormat(t)
	default:
		return ""
	}
}

// NewSplitFileName creates a new file path with split count
func NewSplitFileName(path string, count int) string {
	logFileNameExt := filepath.Ext(path)
	pathWithoutExt := strings.TrimSuffix(path, logFileNameExt)
	return pathWithoutExt + strconv.Itoa(count) + logFileNameExt
}
