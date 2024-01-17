package services

import (
	"fmt"
	"iot-golang/internal/helpers"
	"iot-golang/internal/thm/models"
	"iot-golang/internal/thm/repositories"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type thmService struct {
	thmRepo repositories.ThmRepository
}

// Create implements ThmService.
func (service *thmService) Create(thm models.Thm) helpers.Response {
	var response helpers.Response
	if err := service.thmRepo.Create(thm); err != nil {
		log.Println("ERROR: Gagal membuat data sensor baru" + err.Error())
		response.Status = 500
		response.Messages = "Gagal membuat data sensor baru"
	} else {
		log.Println("SUCCESS: Berhasil membuat data sensor baru")
		response.Status = 201
		response.Messages = "Berhasil membuat data sensor baru"
	}
	return response
}

// Delete implements ThmService.
func (service *thmService) Delete(Id int64) helpers.Response {
	var response helpers.Response

	// Cek apakah data ada sebelum dihapus
	_, findErr := service.thmRepo.GetById(Id)
	if findErr != nil {
		log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d, error: %v", Id, findErr)
		response.Status = 404
		response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		return response
	}

	// Data ditemukan, lakukan penghapusan
	err := service.thmRepo.Delete(Id)
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

// GetAll implements ThmService.
func (service *thmService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.thmRepo.GetAll()
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

// GetById implements ThmService
func (service *thmService) GetById(Id int64) helpers.Response {
	var response helpers.Response
	data, err := service.thmRepo.GetById(Id)

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
		// Konversi string ke tipe data float64
		temperature, _ := strconv.ParseFloat(data.Temperature, 64)
		kelembabanUdara, _ := strconv.ParseFloat(data.KelembabanUdara, 64)
		battery, _ := strconv.ParseFloat(data.Battery, 64)

		data.Temperature = fmt.Sprintf("%.2f", temperature)
		data.KelembabanUdara = fmt.Sprintf("%.2f", kelembabanUdara)
		data.Battery = fmt.Sprintf("%.2f", battery)

		log.Printf("SUCCESS: Berhasil mengambil data sensor dengan id : %d", Id)
		response.Status = 200
		response.Messages = fmt.Sprintf("Berhasil mengambil data sensor dengan id : %d", Id)
		response.Data = data
	}

	return response
}

// GetByToken implements ThmService.
func (service *thmService) GetByToken(DeviceToken string) helpers.Response {
	var response helpers.Response
	data, err := service.thmRepo.GetByToken(DeviceToken)

	if err != nil {
		log.Printf("ERROR: Gagal mengambil data sensor dengan token : %s, error: %v", DeviceToken, err)
		response.Status = 500
		response.Messages = fmt.Sprintf("Gagal mengambil data sensor dengan token : %s", DeviceToken)
		return response
	}

	var foundData []models.Thm
	for _, deviceData := range data {
		if deviceData.DeviceToken == DeviceToken {
			// Konversi string ke tipe data float64
			temperature, _ := strconv.ParseFloat(deviceData.Temperature, 64)
			kelembabanUdara, _ := strconv.ParseFloat(deviceData.KelembabanUdara, 64)
			battery, _ := strconv.ParseFloat(deviceData.Battery, 64)

			deviceData.Temperature = fmt.Sprintf("%.2f", temperature)
			deviceData.KelembabanUdara = fmt.Sprintf("%.2f", kelembabanUdara)
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

// Update implements ThmService.
func (service *thmService) Update(Id int64, thm models.Thm) helpers.Response {
	var response helpers.Response

	// Cek apakah data ada sebelum di update
	_, findErr := service.thmRepo.GetById(Id)
	if findErr != nil {
		log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d, error: %v", Id, findErr)
		response.Status = 404
		response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		return response
	}

	// Data ditemukan, lakukan update
	err := service.thmRepo.Update(Id, thm)
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

type ThmService interface {
	Create(thm models.Thm) helpers.Response
	Update(Id int64, thm models.Thm) helpers.Response
	Delete(Id int64) helpers.Response
	GetById(Id int64) helpers.Response
	GetByToken(DeviceToken string) helpers.Response
	GetAll() helpers.Response
}

func NewThmService(db *gorm.DB) ThmService {
	return &thmService{thmRepo: repositories.NewThmRepository(db)}
}
