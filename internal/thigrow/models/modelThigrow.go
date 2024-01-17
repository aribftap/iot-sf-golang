package models

type Thigrow struct {
	Id                int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	DeviceToken       string `json:"device_token" gorm:"column:device_token"`
	KelembabanTanahTh int32  `json:"kelembaban_tanah_th" gorm:"column:kelembaban_tanah_th"`
	KelembabanTanahSm int32  `json:"kelembaban_tanah_sm" gorm:"column:kelembaban_tanah_sm"`
	KelembabanUdara   int32  `json:"kelembaban_udara" gorm:"column:kelembaban_udara"`
	IntensitasCahaya  string `json:"intensitas_cahaya" gorm:"column:i_cahaya"`
	Battery           string `json:"battery" gorm:"column:battery"`
	Temperature       string `json:"temperature" gorm:"column:temperature"`
	KadarGaram        string `json:"kadar_garam" gorm:"column:kadar_garam"`
}

func (Thigrow) TableName() string {
	return "db_sensor_thigrow"
}
