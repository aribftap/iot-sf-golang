package models

type Thm struct {
	Id              int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	DeviceToken     string `json:"device_token" gorm:"column:device_token"`
	Temperature     string `json:"temperature" gorm:"column:temperature"`
	KelembabanUdara string `json:"kelembaban_udara" gorm:"column:kelembaban_udara"`
	Battery         string `json:"battery" gorm:"column:battery"`
}

func (Thm) TableName() string {
	return "db_sensor_thm30d"
}
