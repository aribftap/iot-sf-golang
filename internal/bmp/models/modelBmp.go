package models

type Bmp struct {
	Id              int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	DeviceToken     string `json:"device_token" gorm:"column:device_token"`
	TekananUdara    string `json:"tekanan_udara" gorm:"column:tekanan_udara"`
	TinggiPermukaan string `json:"tinggi_permukaan" gorm:"column:tinggi_permukaan"`
	Battery         string `json:"battery" gorm:"column:battery"`
}

func (Bmp) TableName() string {
	return "db_sensor_bmp180"
}
