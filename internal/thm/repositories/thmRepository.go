package repositories

import (
	"iot-golang/internal/thm/models"

	"gorm.io/gorm"
)

type dbThm struct {
	Conn *gorm.DB
}

// Create implements ThmRepository.
func (db *dbThm) Create(thm models.Thm) error {
	return db.Conn.Create(&thm).Error
}

// Delete implements ThmRepository.
func (db *dbThm) Delete(Id int64) error {
	return db.Conn.Delete(&models.Thm{Id: Id}).Error
}

// GetAll implements ThmRepository.
func (db *dbThm) GetAll() ([]models.Thm, error) {
	var data []models.Thm
	result := db.Conn.Find(&data)
	return data, result.Error
}

// GetById implements ThmRepository.
func (db *dbThm) GetById(Id int64) (models.Thm, error) {
	var data models.Thm
	result := db.Conn.Where("id", Id).First(&data)
	return data, result.Error
}

// GetByToken implements ThmRepository.
func (db *dbThm) GetByToken(DeviceToken string) ([]models.Thm, error) {
	var data []models.Thm
	result := db.Conn.Where("device_token", DeviceToken).Find(&data)
	return data, result.Error
}

// Update implements ThmRepository.
func (db *dbThm) Update(Id int64, thm models.Thm) error {
	return db.Conn.Where("id", Id).Updates(thm).Error
}

type ThmRepository interface {
	Create(thm models.Thm) error
	Update(Id int64, thm models.Thm) error
	Delete(Id int64) error
	GetById(Id int64) (models.Thm, error)
	GetByToken(DeviceToken string) ([]models.Thm, error)
	GetAll() ([]models.Thm, error)
}

func NewThmRepository(Conn *gorm.DB) ThmRepository {
	return &dbThm{Conn: Conn}
}
