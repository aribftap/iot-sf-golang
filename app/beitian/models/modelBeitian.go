package models

type Beitian struct {
	Id          int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	DeviceToken string `json:"device_token" gorm:"column:device_token"`
	Latitude    string `json:"latitude" gorm:"column:latitude"`
	Longitude   string `json:"longitude" gorm:"column:longitude"`
	Battery     string `json:"battery" gorm:"column:battery"`
}

func (Beitian) TableName() string {
	return "db_sensor_beitian220"
}
