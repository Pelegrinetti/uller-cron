package db

import (
    "github.com/jinzhu/gorm"
    "github.com/sirupsen/logrus"
)

// Metric represents all collected data
type Metric struct {
	gorm.Model
	Lumity      float64
	Temperature float64
	Humidity    float64
}

func (m Metric) Create(db *gorm.DB) {
    logrus.Info("Creating metric...")
    db.Create(&Metric{
        Lumity:      m.Lumity, 
        Temperature: m.Temperature, 
        Humidity:    m.Humidity,
    })
    logrus.Info("Metric created!")
}