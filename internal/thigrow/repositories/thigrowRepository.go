package repositories

import (
	"iot-golang/internal/thigrow/models"

	"gorm.io/gorm"
)

type dbThigrow struct {
	Conn *gorm.DB
}

// GetByToken implements ThigrowRepository.
func (db *dbThigrow) GetByToken(DeviceToken string) ([]models.Thigrow, error) {
	var data []models.Thigrow
	result := db.Conn.Where("device_token", DeviceToken).Find(&data)
	return data, result.Error
}

// Create implements ThigrowRepository.
func (db *dbThigrow) Create(thigrow models.Thigrow) error {
	return db.Conn.Create(&thigrow).Error
}

// Delete implements ThigrowRepository.
func (db *dbThigrow) Delete(Id int64) error {
	return db.Conn.Delete(&models.Thigrow{Id: Id}).Error
}

// GetAll implements ThigrowRepository.
func (db *dbThigrow) GetAll() ([]models.Thigrow, error) {
	var data []models.Thigrow
	result := db.Conn.Find(&data)
	return data, result.Error
}

// GetById implements ThigrowRepository.
func (db *dbThigrow) GetById(Id int64) (models.Thigrow, error) {
	var data models.Thigrow
	result := db.Conn.Where("id", Id).First(&data)
	return data, result.Error
}

// Update implements ThigrowRepository.
func (db *dbThigrow) Update(Id int64, thigrow models.Thigrow) error {
	return db.Conn.Where("id", Id).Updates(thigrow).Error
}

type ThigrowRepository interface {
	Create(thigrow models.Thigrow) error
	Update(Id int64, thigrow models.Thigrow) error
	Delete(Id int64) error
	GetById(Id int64) (models.Thigrow, error)
	GetByToken(DeviceToken string) ([]models.Thigrow, error)
	GetAll() ([]models.Thigrow, error)
}

func NewThigrowRepository(Conn *gorm.DB) ThigrowRepository {
	return &dbThigrow{Conn: Conn}
}
