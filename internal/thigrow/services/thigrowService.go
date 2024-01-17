package services

import (
	"fmt"
	"iot-golang/internal/helpers"
	"iot-golang/internal/thigrow/models"
	"iot-golang/internal/thigrow/repositories"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type thigrowService struct {
	thigrowRepo repositories.ThigrowRepository
}

// GetByToken implements ThigrowService.
func (service *thigrowService) GetByToken(DeviceToken string) helpers.Response {
	var response helpers.Response
	data, err := service.thigrowRepo.GetByToken(DeviceToken)

	if err != nil {
		log.Printf("ERROR: Gagal mengambil data sensor dengan token : %s, error: %v", DeviceToken, err)
		response.Status = 500
		response.Messages = fmt.Sprintf("Gagal mengambil data sensor dengan token : %s", DeviceToken)
		return response
	}

	var foundData []models.Thigrow
	for _, deviceData := range data {
		if deviceData.DeviceToken == DeviceToken {
			// Konversi string ke tipe data float64
			intenCahaya, _ := strconv.ParseFloat(deviceData.IntensitasCahaya, 64)
			battery, _ := strconv.ParseFloat(deviceData.Battery, 64)
			temperature, _ := strconv.ParseFloat(deviceData.Temperature, 64)

			deviceData.IntensitasCahaya = fmt.Sprintf("%.2f", intenCahaya)
			deviceData.Battery = fmt.Sprintf("%.2f", battery)
			deviceData.Temperature = fmt.Sprintf("%.2f", temperature)

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

// Create implements ThigrowService.
func (service *thigrowService) Create(thigrow models.Thigrow) helpers.Response {
	var response helpers.Response
	if err := service.thigrowRepo.Create(thigrow); err != nil {
		log.Println("ERROR: Gagal membuat data sensor baru" + err.Error())
		response.Status = 500
		response.Messages = "Gagal membuat data sensor baru"
	} else {
		response.Status = 201
		response.Messages = "Berhasil membuat data sensor baru"
	}
	return response
}

// Delete implements ThigrowService.
func (service *thigrowService) Delete(Id int64) helpers.Response {
	var response helpers.Response

	// Cek apakah data ada sebelum dihapus
	_, findErr := service.thigrowRepo.GetById(Id)
	if findErr != nil {
		log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d, error: %v", Id, findErr)
		response.Status = 404
		response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		return response
	}

	// Data ditemukan, lakukan penghapusan
	err := service.thigrowRepo.Delete(Id)
	if err != nil {
		log.Printf("ERROR: Gagal menghapus data sensor dengan id : %d, error: %v", Id, err)
		response.Status = 500
		response.Messages = fmt.Sprintf("Gagal menghapus data sensor dengan id : %d", Id)
	} else {
		// Jika penghapusan berhasil
		response.Status = 200
		response.Messages = "Data sensor berhasil dihapus"
	}

	return response
}

// GetAll implements ThigrowService.
func (service *thigrowService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.thigrowRepo.GetAll()
	if err != nil {
		log.Printf("ERROR: Gagal mengambil seluruh data sensor : error %v", err)
		response.Status = 500
		response.Messages = "Gagal mengambil seluruh data sensor"
	} else {
		response.Status = 200
		response.Messages = "Berhasil mengambil semua data sensor"
		response.Data = data
	}
	return response
}

// GetById implements ThigrowService.
func (service *thigrowService) GetById(Id int64) helpers.Response {
	var response helpers.Response
	data, err := service.thigrowRepo.GetById(Id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("INFO: Tidak menemukan data sensor dengan id : %d", Id)
			response.Status = 404
			response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		} else {
			log.Printf("ERROR: Gagal mengambil data sensor dengan id : %d, error: %v", Id, err)
			response.Status = 500
			response.Messages = fmt.Sprintf("Gagal mengambil data sensor dengan id : %d", Id)
		}
	} else {
		// Konversi string ke tipe data float64
		intenCahaya, _ := strconv.ParseFloat(data.IntensitasCahaya, 64)
		battery, _ := strconv.ParseFloat(data.Battery, 64)
		temperature, _ := strconv.ParseFloat(data.Temperature, 64)

		data.IntensitasCahaya = fmt.Sprintf("%.2f", intenCahaya)
		data.Battery = fmt.Sprintf("%.2f", battery)
		data.Temperature = fmt.Sprintf("%.2f", temperature)

		response.Status = 200
		response.Messages = fmt.Sprintf("Berhasil mengambil data sensor dengan id : %d", Id)
		response.Data = data
	}

	return response
}

// Update implements ThigrowService.
func (service *thigrowService) Update(Id int64, thigrow models.Thigrow) helpers.Response {
	var response helpers.Response

	// Cek apakah data ada sebelum di update
	_, findErr := service.thigrowRepo.GetById(Id)
	if findErr != nil {
		log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d, error: %v", Id, findErr)
		response.Status = 404
		response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		return response
	}

	// Data ditemukan, lakukan update
	err := service.thigrowRepo.Update(Id, thigrow)
	if err != nil {
		log.Printf("ERROR: Gagal mengubah data sensor dengan id : %d, error: %v", Id, err)
		response.Status = 500
		response.Messages = fmt.Sprintf("Gagal mengubah data sensor dengan id : %d", Id)
	} else {
		// Jika update data berhasil
		response.Status = 200
		response.Messages = "Berhasil mengubah data sensor"
	}

	return response
}

type ThigrowService interface {
	Create(thigrow models.Thigrow) helpers.Response
	Update(Id int64, thigrow models.Thigrow) helpers.Response
	Delete(Id int64) helpers.Response
	GetById(Id int64) helpers.Response
	GetByToken(DeviceToken string) helpers.Response
	GetAll() helpers.Response
}

func NewThigrowService(db *gorm.DB) ThigrowService {
	return &thigrowService{thigrowRepo: repositories.NewThigrowRepository(db)}
}
