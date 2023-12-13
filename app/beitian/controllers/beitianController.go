package controllers

import (
	"fmt"
	"iot-golang/app/beitian/models"
	"iot-golang/app/beitian/services"
	"iot-golang/app/helpers"
	"net/http"
	"reflect"
	"strconv"

	v1 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BeitianController struct {
	beitianService services.BeitianService
	validate       v1.Validate
}

var httpStatus int

func (controller BeitianController) Create(c echo.Context) error {
	type payload struct {
		DeviceToken string `json:"device_token" validate:"required"`
		Latitude    string `json:"latitude" validate:"required,numeric"`
		Longitude   string `json:"longitude" validate:"required,numeric"`
		Battery     string `json:"battery" validate:"required,numeric"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ValidationResponse{
			Status:   400,
			Errors:   err,
			Messages: "Request di tolak",
		})
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
			} else if e.Tag() == "numeric" {
				errMsg = fmt.Sprintf("Field %s tidak boleh huruf", fieldName)
			}
			errorList[fieldName] = errMsg

		}

		return c.JSON(http.StatusBadRequest, helpers.ValidationResponse{
			Status:   400,
			Errors:   errorList,
			Messages: "Request di tolak",
		})
	}

	result := controller.beitianService.Create(models.Beitian{DeviceToken: payloadValidator.DeviceToken, Latitude: payloadValidator.Latitude, Longitude: payloadValidator.Longitude, Battery: payloadValidator.Battery})

	if result.Status == 201 {
		httpStatus = http.StatusCreated
	} else if result.Status == 400 {
		httpStatus = http.StatusBadRequest
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}
	return c.JSON(httpStatus, result)
}

func (controller BeitianController) Update(c echo.Context) error {
	type payload struct {
		Latitude  string `json:"latitude" validate:"required"`
		Longitude string `json:"longitude" validate:"required"`
		Battery   string `json:"battery" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ValidationResponse{
			Status:   400,
			Errors:   err,
			Messages: "Invalid request body",
		})
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

	idBeitian, _ := strconv.Atoi(c.Param("id"))
	result := controller.beitianService.Update(int64(idBeitian), models.Beitian{Latitude: payloadValidator.Latitude, Longitude: payloadValidator.Longitude, Battery: payloadValidator.Battery})

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller BeitianController) Delete(c echo.Context) error {
	idBeitian, _ := strconv.Atoi(c.Param("id"))
	result := controller.beitianService.Delete(int64(idBeitian))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller BeitianController) GetAll(c echo.Context) error {
	result := controller.beitianService.GetAll()

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller BeitianController) GetById(c echo.Context) error {
	idBeitian, _ := strconv.Atoi(c.QueryParam("id"))
	result := controller.beitianService.GetById(int64(idBeitian))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller BeitianController) GetByToken(c echo.Context) error {
	idTokenBeitian := c.QueryParam("device_token")
	result := controller.beitianService.GetByToken(idTokenBeitian)

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller BeitianController) GetNewByToken(c echo.Context) error {
	idTokenBeitian := c.QueryParam("device_token")
	result := controller.beitianService.GetNewByToken(idTokenBeitian)

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func NewBeitianController(db *gorm.DB) BeitianController {
	service := services.NewBeitianService(db)
	controller := BeitianController{
		beitianService: service,
		validate:       *v1.New(),
	}

	return controller
}
