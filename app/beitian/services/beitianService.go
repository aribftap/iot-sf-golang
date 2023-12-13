package services

import (
	"fmt"
	"iot-golang/app/beitian/models"
	"iot-golang/app/beitian/repositories"
	"iot-golang/app/helpers"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type beitianService struct {
	beitianRepo repositories.BeitianRepository
}

// GetNewByToken implements BeitianService.
func (service *beitianService) GetNewByToken(DeviceToken string) helpers.Response {
	var response helpers.Response
	data, err := service.beitianRepo.GetNewByToken(DeviceToken)

	if err != nil {
		log.Printf("ERROR: Gagal mengambil data sensor dengan token : %s, error: %v", DeviceToken, err)
		response.Status = 500
		response.Messages = fmt.Sprintf("Gagal mengambil data sensor dengan token : %s", DeviceToken)
	} else {
		log.Printf("INFO: Tidak menemukan data sensor dengan token : %s", DeviceToken)
		response.Status = 404
		response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan token : %s", DeviceToken)
	}

	log.Printf("SUCCESS: Berhasil mengambil data sensor dengan token : %s", DeviceToken)
	response.Status = 200
	response.Messages = fmt.Sprintf("Berhasil mengambil data sensor dengan token : %s", DeviceToken)
	response.Data = data

	return response
}

// Create implements BeitianService.
func (service *beitianService) Create(beitian models.Beitian) helpers.Response {
	var response helpers.Response
	if err := service.beitianRepo.Create(beitian); err != nil {
		log.Printf("ERROR: Gagal membuat data sensor baru, error: %v", err)
		response.Status = 500
		response.Messages = "Gagal membuat data sensor baru"
	} else {
		log.Println("SUCCESS: Berhasil membuat data sensor baru")
		response.Status = 201
		response.Messages = "Berhasil membuat data sensor baru"
	}

	return response
}

// Delete implements BeitianService.
func (service *beitianService) Delete(Id int64) helpers.Response {
	var response helpers.Response

	// Cek apakah data ada sebelum dihapus
	_, findErr := service.beitianRepo.GetById(Id)
	if findErr != nil {
		if findErr == gorm.ErrRecordNotFound {
			log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d", Id)
			response.Status = 404
			response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		} else {
			log.Printf("ERROR: Gagal mengambil data sensor dengan id : %d, error: %v", Id, findErr)
			response.Status = 500
			response.Messages = fmt.Sprintf("Gagal mengambil data sensor dengan id : %d", Id)
		}
		return response
	}

	// Data ditemukan, lakukan penghapusan
	err := service.beitianRepo.Delete(Id)
	if err != nil {
		log.Printf("ERROR: Gagal menghapus data sensor dengan id : %d, error: %v", Id, err)
		response.Status = 500
		response.Messages = fmt.Sprintf("Gagal menghapus data sensor dengan id : %d", Id)
	} else {
		// Jika penghapusan berhasil
		log.Println("SUCCESS: Data sensor berhasil dihapus")
		response.Status = 200
		response.Messages = "Data sensor berhasil dihapus"
	}

	return response
}

// GetAll implements BeitianService.
func (service *beitianService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.beitianRepo.GetAll()
	if err != nil {
		log.Printf("ERROR: Gagal mengambil seluruh data sensor : error %v", err)
		response.Status = 500
		response.Messages = "Gagal mengambil seluruh data sensor"
	} else {
		log.Println("SUCCESS: Berhasil mengambil semua data sensor")
		response.Status = 200
		response.Messages = "Berhasil mengambil semua data sensor"
		response.Data = data
	}
	return response
}

// GetById implements BeitianService.
func (service *beitianService) GetById(Id int64) helpers.Response {
	var response helpers.Response
	data, err := service.beitianRepo.GetById(Id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d", Id)
			response.Status = 404
			response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		} else {
			log.Printf("ERROR: Gagal mengambil data sensor dengan id : %d, error: %v", Id, err)
			response.Status = 500
			response.Messages = fmt.Sprintf("Gagal mengambil data sensor dengan id : %d", Id)
		}
	} else {
		battery, _ := strconv.ParseFloat(data.Battery, 64)
		data.Battery = fmt.Sprintf("%.2f", battery)

		log.Printf("SUCCESS: Berhasil mengambil data sensor dengan id : %d", Id)
		response.Status = 200
		response.Messages = fmt.Sprintf("Berhasil mengambil data sensor dengan id : %d", Id)
		response.Data = data
	}

	return response
}

func (service *beitianService) GetByToken(DeviceToken string) helpers.Response {
	var response helpers.Response
	data, err := service.beitianRepo.GetByToken(DeviceToken)

	if err != nil {
		log.Printf("ERROR: Gagal mengambil data sensor dengan token : %s, error: %v", DeviceToken, err)
		response.Status = 500
		response.Messages = fmt.Sprintf("Gagal mengambil data sensor dengan token : %s", DeviceToken)
		return response
	}

	var foundData []models.Beitian
	for _, deviceData := range data {
		if deviceData.DeviceToken == DeviceToken {
			// Konversi string ke tipe data float64
			battery, _ := strconv.ParseFloat(deviceData.Battery, 64)
			deviceData.Battery = fmt.Sprintf("%.2f", battery)
			foundData = append(foundData, deviceData)
		}
	}

	if len(foundData) == 0 {
		log.Printf("INFO: Tidak menemukan data sensor dengan token : %s", DeviceToken)
		response.Status = 404
		response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan token : %s", DeviceToken)
	} else {
		log.Printf("SUCCESS: Berhasil mengambil data sensor dengan token : %s", DeviceToken)
		response.Status = 200
		response.Messages = fmt.Sprintf("Berhasil mengambil data sensor dengan token : %s", DeviceToken)
		response.Data = foundData
	}

	return response
}

// Update implements BeitianService.
func (service *beitianService) Update(Id int64, beitian models.Beitian) helpers.Response {
	var response helpers.Response

	// Cek apakah data ada sebelum di update
	_, findErr := service.beitianRepo.GetById(Id)
	if findErr != nil {
		if findErr == gorm.ErrRecordNotFound {
			log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d", Id)
			response.Status = 404
			response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		} else {
			log.Printf("ERROR: Gagal mengambil data sensor dengan id : %d, error: %v", Id, findErr)
			response.Status = 500
			response.Messages = fmt.Sprintf("Gagal mengambil data sensor dengan id : %d", Id)
		}
	}

	// Data ditemukan, lakukan update
	err := service.beitianRepo.Update(Id, beitian)
	if err != nil {
		log.Printf("ERROR: Gagal mengubah data sensor dengan id : %d, error: %v", Id, err)
		response.Status = 500
		response.Messages = fmt.Sprintf("Gagal mengubah data sensor dengan id : %d", Id)
	} else {
		// Jika update data berhasil
		log.Println("SUCCESS: Berhasil mengubah data sensor")
		response.Status = 200
		response.Messages = "Berhasil mengubah data sensor"
	}

	return response
}

type BeitianService interface {
	Create(beitian models.Beitian) helpers.Response
	Update(Id int64, beitian models.Beitian) helpers.Response
	Delete(Id int64) helpers.Response
	GetById(Id int64) helpers.Response
	GetByToken(DeviceToken string) helpers.Response
	GetAll() helpers.Response
	GetNewByToken(DeviceToken string) helpers.Response
}

func NewBeitianService(db *gorm.DB) BeitianService {
	return &beitianService{beitianRepo: repositories.NewBeitianRepository(db)}
}
