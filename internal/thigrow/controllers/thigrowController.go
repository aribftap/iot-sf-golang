package controllers

import (
	"fmt"
	"iot-golang/internal/helpers"
	"iot-golang/internal/thigrow/models"
	"iot-golang/internal/thigrow/services"
	"net/http"
	"reflect"
	"strconv"

	v1 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ThigrowController struct {
	thigrowService services.ThigrowService
	validate       v1.Validate
}

var httpStatus int

func (controller ThigrowController) Create(c echo.Context) error {
	type payload struct {
		DeviceToken       string `json:"device_token" validate:"required"`
		KelembabanTanahTh int32  `json:"kelembaban_tanah_th" validate:"required"`
		KelembabanTanahSm int32  `json:"kelembaban_tanah_sm" validate:"required"`
		KelembabanUdara   int32  `json:"kelembaban_udara" validate:"required"`
		IntensitasCahaya  string `json:"intensitas_cahaya" validate:"required"`
		Battery           string `json:"battery" validate:"required"`
		Temperature       string `json:"temperature" validate:"required"`
		KadarGaram        string `json:"kadar_garam" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		if err := c.Bind(payloadValidator); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ValidationResponse{
				Status:   400,
				Errors:   err,
				Messages: "Invalid request body",
			})
		}
	}

	err := controller.validate.Struct(payloadValidator)
	if err != nil {
		errors := err.(v1.ValidationErrors)
		errorList := make(map[string]string)

		for _, e := range errors {
			var errMsg string
			field, _ := reflect.TypeOf(*payloadValidator).FieldByName(e.StructField())
			fieldName := field.Tag.Get("json")

			if e.Tag() == "required" {
				errMsg = fmt.Sprintf("Field %s tidak boleh kosong", fieldName)
			}
			errorList[fieldName] = errMsg
		}

		return c.JSON(http.StatusBadRequest, helpers.ValidationResponse{
			Status:   400,
			Errors:   errorList,
			Messages: "Request di tolak",
		})
	}

	result := controller.thigrowService.Create(models.Thigrow{DeviceToken: payloadValidator.DeviceToken, KelembabanTanahTh: payloadValidator.KelembabanTanahTh, KelembabanTanahSm: payloadValidator.KelembabanTanahSm, KelembabanUdara: payloadValidator.KelembabanUdara, IntensitasCahaya: payloadValidator.IntensitasCahaya, Battery: payloadValidator.Battery, Temperature: payloadValidator.Temperature, KadarGaram: payloadValidator.KadarGaram})

	if result.Status == 201 {
		httpStatus = http.StatusCreated
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller ThigrowController) Update(c echo.Context) error {
	type payload struct {
		KelembabanTanahTh int32  `json:"kelembaban_tanah_th" validate:"required"`
		KelembabanTanahSm int32  `json:"kelembaban_tanah_sm" validate:"required"`
		KelembabanUdara   int32  `json:"kelembaban_udara" validate:"required"`
		IntensitasCahaya  string `json:"intensitas_cahaya" validate:"required"`
		Battery           string `json:"battery" validate:"required"`
		Temperature       string `json:"temperature" validate:"required"`
		KadarGaram        string `json:"kadar_garam" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		if err := c.Bind(payloadValidator); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ValidationResponse{
				Status:   400,
				Errors:   err,
				Messages: "Invalid request body",
			})
		}
	}

	err := controller.validate.Struct(payloadValidator)
	if err != nil {
		errors := err.(v1.ValidationErrors)
		errorList := make(map[string]string)

		for _, e := range errors {
			var errMsg string
			field, _ := reflect.TypeOf(*payloadValidator).FieldByName(e.StructField())
			fieldName := field.Tag.Get("json")

			if e.Tag() == "required" {
				errMsg = fmt.Sprintf("Field %s tidak boleh kosong", fieldName)
			}
			errorList[fieldName] = errMsg
		}

		return c.JSON(http.StatusBadRequest, helpers.ValidationResponse{
			Status:   400,
			Errors:   errorList,
			Messages: "Request di tolak",
		})
	}

	idThigrow, _ := strconv.Atoi(c.Param("id"))
	result := controller.thigrowService.Update(int64(idThigrow), models.Thigrow{KelembabanTanahTh: payloadValidator.KelembabanTanahTh, KelembabanTanahSm: payloadValidator.KelembabanTanahSm, KelembabanUdara: payloadValidator.KelembabanUdara, IntensitasCahaya: payloadValidator.IntensitasCahaya, Battery: payloadValidator.Battery, Temperature: payloadValidator.Temperature, KadarGaram: payloadValidator.KadarGaram})

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller ThigrowController) Delete(c echo.Context) error {
	idThigrow, _ := strconv.Atoi(c.Param("id"))
	result := controller.thigrowService.Delete(int64(idThigrow))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller ThigrowController) GetAll(c echo.Context) error {
	result := controller.thigrowService.GetAll()
	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}
	return c.JSON(httpStatus, result)
}

func (controller ThigrowController) GetById(c echo.Context) error {
	idThigrow, _ := strconv.Atoi(c.QueryParam("id"))
	result := controller.thigrowService.GetById(int64(idThigrow))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller ThigrowController) GetByToken(c echo.Context) error {
	idTokenThigrow := c.QueryParam("device_token")
	result := controller.thigrowService.GetByToken(idTokenThigrow)

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func NewThigrowController(db *gorm.DB) ThigrowController {
	service := services.NewThigrowService(db)
	controller := ThigrowController{
		thigrowService: service,
		validate:       *v1.New(),
	}

	return controller
}
