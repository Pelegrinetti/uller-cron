package sensor

import (
	"math/rand"

	"github.com/Pelegrinetti/uller/package/store"
)

// GetSensorData get all sensor's data
func GetSensorData() store.Metric {
	return store.Metric{
		Lumity:      rand.Int63(),
		Temperature: rand.Int63(),
		Humidity:    rand.Int63(),
	}
}
