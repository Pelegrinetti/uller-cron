package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"

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

func reading(pin rpio.Pin) {
	fmt.Println("Reading...")
	count := 0
	for pin.Read() == 0 {
		count++
	}

	fmt.Printf("LDR: %d\n", count)

	time.Sleep(time.Millisecond * 300)

	reading(pin)
}

func main() {
	logrus.Info("Running!")

	err := rpio.Open()
	if err != nil {
		logrus.WithError(err).Error("Error while opening GPIO connection.")
	}

	pin := rpio.Pin(12)

	fmt.Println(pin)
	pin.Input()

	reading(pin)

	// metric := sensor.GetSensorData()
	// store.Store(metric)

	// fmt.Println(store.Read())

	// cron()
}
