package services

import (
	"fmt"
	"iot-golang/internal/helpers"
	"iot-golang/internal/ina/models"
	"iot-golang/internal/ina/repositories"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type inaService struct {
	inaRepo repositories.InaRepository
}

// Create implements InaService.
func (service *inaService) Create(ina models.Ina) helpers.Response {
	var response helpers.Response
	if err := service.inaRepo.Create(ina); err != nil {
		log.Println("ERROR: Gagal membuat data sensor baru, error :" + err.Error())
		response.Status = 500
		response.Messages = "Gagal membuat data sensor baru"
	} else {
		log.Println("SUCCESS: Berhasil membuat data sensor baru")
		response.Status = 201
		response.Messages = "Berhasil membuat data sensor baru"
	}
	return response
}

// Delete implements InaService.
func (service *inaService) Delete(Id int64) helpers.Response {
	var response helpers.Response

	// Cek apakah data ada sebelum dihapus
	_, findErr := service.inaRepo.GetById(Id)
	if findErr != nil {
		log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d, error: %v", Id, findErr)
		response.Status = 404
		response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		return response
	}

	// Data ditemukan, lakukan penghapusan
	err := service.inaRepo.Delete(Id)
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

// GetAll implements InaService.
func (service *inaService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.inaRepo.GetAll()
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

// GetById implements InaService.
func (service *inaService) GetById(Id int64) helpers.Response {
	var response helpers.Response
	data, err := service.inaRepo.GetById(Id)

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
		Tegangan, _ := strconv.ParseFloat(data.Tegangan, 64)
		Arus, _ := strconv.ParseFloat(data.Arus, 64)
		Daya, _ := strconv.ParseFloat(data.Daya, 64)

		data.Tegangan = fmt.Sprintf("%.2f", Tegangan)
		data.Arus = fmt.Sprintf("%.2f", Arus)
		data.Daya = fmt.Sprintf("%.2f", Daya)

		log.Printf("SUCCESS: Berhasil mengambil data sensor dengan id : %d", Id)
		response.Status = 200
		response.Messages = fmt.Sprintf("Berhasil mengambil data sensor dengan id : %d", Id)
		response.Data = data
	}

	return response
}

// GetByToken implements InaService.
func (service *inaService) GetByToken(DeviceToken string) helpers.Response {
	var response helpers.Response
	data, err := service.inaRepo.GetByToken(DeviceToken)

	if err != nil {
		log.Printf("ERROR: Gagal mengambil data sensor dengan token : %s, error: %v", DeviceToken, err)
		response.Status = 500
		response.Messages = fmt.Sprintf("Gagal mengambil data sensor dengan token : %s", DeviceToken)
		return response
	}

	var foundData []models.Ina
	for _, deviceData := range data {
		if deviceData.DeviceToken == DeviceToken {
			// Konversi string ke tipe data float64
			Tegangan, _ := strconv.ParseFloat(deviceData.Tegangan, 64)
			Arus, _ := strconv.ParseFloat(deviceData.Arus, 64)
			Daya, _ := strconv.ParseFloat(deviceData.Daya, 64)

			deviceData.Tegangan = fmt.Sprintf("%.2f", Tegangan)
			deviceData.Arus = fmt.Sprintf("%.2f", Arus)
			deviceData.Daya = fmt.Sprintf("%.2f", Daya)

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

// Update implements InaService.
func (service *inaService) Update(Id int64, ina models.Ina) helpers.Response {
	var response helpers.Response

	// Cek apakah data ada sebelum di update
	_, findErr := service.inaRepo.GetById(Id)
	if findErr != nil {
		log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d, error: %v", Id, findErr)
		response.Status = 404
		response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		return response
	}

	// Data ditemukan, lakukan update
	err := service.inaRepo.Update(Id, ina)
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

type InaService interface {
	Create(ina models.Ina) helpers.Response
	Update(Id int64, ina models.Ina) helpers.Response
	Delete(Id int64) helpers.Response
	GetById(Id int64) helpers.Response
	GetByToken(DeviceToken string) helpers.Response
	GetAll() helpers.Response
}

func NewInaService(db *gorm.DB) InaService {
	return &inaService{inaRepo: repositories.NewInaRepository(db)}
}
