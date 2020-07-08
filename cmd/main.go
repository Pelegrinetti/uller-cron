package main

import (
	"time"

	"github.com/Pelegrinetti/uller-cron/package/store"
	"github.com/Pelegrinetti/uller-cron/worker/sensor"
	"github.com/sirupsen/logrus"
)

func cron() {
	logrus.Info("Running cron...")

	metric := sensor.GetSensorData()
	store.Store(metric)

	logrus.Info("Cron finished!")

	time.Sleep(time.Second * 3)

	cron()
}

func main() {
	logrus.Info("Running!")

	metric := sensor.GetSensorData()
	store.Store(metric)

	cron()
}
