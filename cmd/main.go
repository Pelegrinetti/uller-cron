package main

import (
	"time"

	"github.com/Pelegrinetti/uller/package/db"
	"github.com/Pelegrinetti/uller/worker/sensor"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func cron(database *gorm.DB) {
	logrus.Info("Running cron...")

	metric := sensor.GetSensorData()
	metric.Create(database)

	time.Sleep(time.Second * 3)

	logrus.Info("Cron finished!")

	cron(database)
}

func main() {
	logrus.Info("Running!")

	logrus.Info("Creating database instance...")
	database := db.New()
	db.Migrate(database)
	defer logrus.Info("Closing database connection!")
	defer database.Close()

	cron(database)
}
