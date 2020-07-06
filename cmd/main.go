package main

import (
	"fmt"
	"time"

	"github.com/Pelegrinetti/uller/package/store"

	"github.com/Pelegrinetti/uller/worker/sensor"
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

	fmt.Println(store.Read())

	cron()
}
