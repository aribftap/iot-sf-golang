package repositories

import (
	"iot-golang/app/bmp/models"

	"gorm.io/gorm"
)

type dbBmp struct {
	Conn *gorm.DB
}

// Create implements BmpRepository.
func (db *dbBmp) Create(bmp models.Bmp) error {
	return db.Conn.Create(&bmp).Error
}

// Delete implements BmpRepository.
func (db *dbBmp) Delete(Id int64) error {
	return db.Conn.Delete(&models.Bmp{Id: Id}).Error
}

// GetAll implements BmpRepository.
func (db *dbBmp) GetAll() ([]models.Bmp, error) {
	var data []models.Bmp
	result := db.Conn.Find(&data)
	return data, result.Error
}

// GetById implements BmpRepository.
func (db *dbBmp) GetById(Id int64) (models.Bmp, error) {
	var data models.Bmp
	result := db.Conn.Where("id", Id).First(&data)
	return data, result.Error
}

// GetByToken implements BmpRepository.
func (db *dbBmp) GetByToken(DeviceToken string) ([]models.Bmp, error) {
	var data []models.Bmp
	result := db.Conn.Where("device_token", DeviceToken).Find(&data)
	return data, result.Error
}

// Update implements BmpRepository.
func (db *dbBmp) Update(Id int64, bmp models.Bmp) error {
	return db.Conn.Where("id", Id).Updates(bmp).Error
}

type BmpRepository interface {
	Create(bmp models.Bmp) error
	Update(Id int64, bmp models.Bmp) error
	Delete(Id int64) error
	GetById(Id int64) (models.Bmp, error)
	GetByToken(DeviceToken string) ([]models.Bmp, error)
	GetAll() ([]models.Bmp, error)
}

func NewBmpRepository(Conn *gorm.DB) BmpRepository {
	return &dbBmp{Conn: Conn}
}
