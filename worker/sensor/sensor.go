package sensor

import (
	"github.com/Pelegrinetti/uller-cron/package/store"
	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio"
)

// GetSensorData get all sensor's data
func GetSensorData() store.Metric {
	err := rpio.Open()
	if err != nil {
		logrus.WithError(err).Error("Error while opening GPIO connection.")
	}

	pin := rpio.Pin(21)

	pin.Input()

	return store.Metric{
		Lumity:      int64(pin.Read()),
		Temperature: 0,
		Humidity:    0,
	}
}
