package services

import (
	"fmt"
	"iot-golang/app/bmp/models"
	"iot-golang/app/bmp/repositories"
	"iot-golang/app/helpers"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type bmpService struct {
	bmpRepo repositories.BmpRepository
}

// Create implements BmpService.
func (service *bmpService) Create(bmp models.Bmp) helpers.Response {
	var response helpers.Response
	if err := service.bmpRepo.Create(bmp); err != nil {
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

// Delete implements BmpService.
func (service *bmpService) Delete(Id int64) helpers.Response {
	var response helpers.Response

	// Cek apakah data ada sebelum dihapus
	_, findErr := service.bmpRepo.GetById(Id)
	if findErr != nil {
		log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d, error: %v", Id, findErr)
		response.Status = 404
		response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		return response
	}

	// Data ditemukan, lakukan penghapusan
	err := service.bmpRepo.Delete(Id)
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

// GetAll implements BmpService.
func (service *bmpService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.bmpRepo.GetAll()
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

// GetById implements BmpService.
func (service *bmpService) GetById(Id int64) helpers.Response {
	var response helpers.Response
	data, err := service.bmpRepo.GetById(Id)

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
		TekananUdara, _ := strconv.ParseFloat(data.TekananUdara, 64)
		TinggiPermukaan, _ := strconv.ParseFloat(data.TinggiPermukaan, 64)
		battery, _ := strconv.ParseFloat(data.Battery, 64)

		data.TekananUdara = fmt.Sprintf("%.2f", TekananUdara)
		data.TinggiPermukaan = fmt.Sprintf("%.2f", TinggiPermukaan)
		data.Battery = fmt.Sprintf("%.2f", battery)

		log.Printf("SUCCESS: Berhasil mengambil data sensor dengan id : %d", Id)
		response.Status = 200
		response.Messages = fmt.Sprintf("Berhasil mengambil data sensor dengan id : %d", Id)
		response.Data = data
	}

	return response
}

// GetByToken implements BmpService.
func (service *bmpService) GetByToken(DeviceToken string) helpers.Response {
	var response helpers.Response
	data, err := service.bmpRepo.GetByToken(DeviceToken)

	if err != nil {
		log.Printf("ERROR: Gagal mengambil data sensor dengan token : %s, error: %v", DeviceToken, err)
		response.Status = 500
		response.Messages = fmt.Sprintf("Gagal mengambil data sensor dengan token : %s", DeviceToken)
		return response
	}

	var foundData []models.Bmp
	for _, deviceData := range data {
		if deviceData.DeviceToken == DeviceToken {
			// Konversi string ke tipe data float64
			TekananUdara, _ := strconv.ParseFloat(deviceData.TekananUdara, 64)
			TinggiPermukaan, _ := strconv.ParseFloat(deviceData.TinggiPermukaan, 64)
			battery, _ := strconv.ParseFloat(deviceData.Battery, 64)

			deviceData.TekananUdara = fmt.Sprintf("%.2f", TekananUdara)
			deviceData.TinggiPermukaan = fmt.Sprintf("%.2f", TinggiPermukaan)
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

// Update implements BmpService.
func (service *bmpService) Update(Id int64, bmp models.Bmp) helpers.Response {
	var response helpers.Response

	// Cek apakah data ada sebelum di update
	_, findErr := service.bmpRepo.GetById(Id)
	if findErr != nil {
		log.Printf("ERROR: Tidak menemukan data sensor dengan id : %d, error: %v", Id, findErr)
		response.Status = 404
		response.Messages = fmt.Sprintf("Tidak menemukan data sensor dengan id : %d", Id)
		return response
	}

	// Data ditemukan, lakukan update
	err := service.bmpRepo.Update(Id, bmp)
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

type BmpService interface {
	Create(bmp models.Bmp) helpers.Response
	Update(Id int64, bmp models.Bmp) helpers.Response
	Delete(Id int64) helpers.Response
	GetById(Id int64) helpers.Response
	GetByToken(DeviceToken string) helpers.Response
	GetAll() helpers.Response
}

func NewBmpService(db *gorm.DB) BmpService {
	return &bmpService{bmpRepo: repositories.NewBmpRepository(db)}
}
