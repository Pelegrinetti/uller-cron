package sensor

import "github.com/Pelegrinetti/uller/package/db"

// GetSensorData get all sensor's data
func GetSensorData() db.Metric {
    return db.Metric{
        Lumity: 14.5214,
        Temperature: 18.531,
        Humidity: 35.6732,
    }
}
