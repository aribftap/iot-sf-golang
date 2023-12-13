package models

type Pzem struct {
	Id          int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	DeviceToken string `json:"device_token" gorm:"column:device_token"`
	Tegangan    string `json:"tegangan" gorm:"column:tegangan"`
	Arus        string `json:"arus" gorm:"column:arus"`
	Daya        string `json:"daya" gorm:"column:daya"`
}

func (Pzem) TableName() string {
	return "db_sensor_pzem"
}
