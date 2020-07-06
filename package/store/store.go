package store

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
)

// Metric represents the data obtained by the raspberry PI
type Metric struct {
	Lumity      int64
	Temperature int64
	Humidity    int64
}

// Store write metrics data
func Store(metric Metric) {
	currentTime := time.Now()
	f, osErr := os.Create(fmt.Sprintf("binaries/%d.bin", currentTime.Unix()))
	if osErr != nil {
		logrus.WithError(osErr).Error("Error while creating data file.")
	}
	defer f.Close()

	writeErr := binary.Write(f, binary.BigEndian, metric)

	if writeErr != nil {
		logrus.WithError(writeErr).Error("Error while writting data.")
	}
}

// Read get all data from data file
func Read() []Metric {
	var metrics []Metric

	var files []string

	r, rErr := regexp.Compile(".*bin$")
	if rErr != nil {
		logrus.WithError(rErr).Error("Error while compiling RegExp.")
	}

	wd, wdErr := os.Getwd()
	if wdErr != nil {
		logrus.WithError(wdErr).Error("Error while getting workdir.")
	}

	root := fmt.Sprintf("%s/binaries", wd)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if r.MatchString(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			logrus.WithError(err).Error("Error while opening data file.")
		}
		defer f.Close()

		var metric Metric
		readErr := binary.Read(f, binary.BigEndian, &metric)

		if readErr != nil {
			logrus.WithError(readErr).Error("Error while reading data file.", file)
		}

		metrics = append(metrics, metric)
	}

	return metrics
}
