package controllers

import (
	"fmt"
	"iot-golang/app/helpers"
	"iot-golang/app/ina/models"
	"iot-golang/app/ina/services"
	"net/http"
	"reflect"
	"strconv"

	v1 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type InaController struct {
	inaService services.InaService
	validate   v1.Validate
}

var httpStatus int

func (controller InaController) Create(c echo.Context) error {
	type payload struct {
		DeviceToken string `json:"device_token" validate:"required"`
		Tegangan    string `json:"tegangan" validate:"required"`
		Arus        string `json:"arus" validate:"required"`
		Daya        string `json:"daya" validate:"required"`
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

	result := controller.inaService.Create(models.Ina{DeviceToken: payloadValidator.DeviceToken, Tegangan: payloadValidator.Tegangan, Arus: payloadValidator.Arus, Daya: payloadValidator.Daya})

	if result.Status == 201 {
		httpStatus = http.StatusCreated
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller InaController) Update(c echo.Context) error {
	type payload struct {
		Tegangan string `json:"tegangan" validate:"required"`
		Arus     string `json:"arus" validate:"required"`
		Daya     string `json:"daya" validate:"required"`
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

	idIna, _ := strconv.Atoi(c.Param("id"))
	result := controller.inaService.Update(int64(idIna), models.Ina{Tegangan: payloadValidator.Tegangan, Arus: payloadValidator.Arus, Daya: payloadValidator.Daya})

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller InaController) Delete(c echo.Context) error {
	idIna, _ := strconv.Atoi(c.Param("id"))
	result := controller.inaService.Delete(int64(idIna))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller InaController) GetAll(c echo.Context) error {
	result := controller.inaService.GetAll()
	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller InaController) GetById(c echo.Context) error {
	idIna, _ := strconv.Atoi(c.QueryParam("id"))
	result := controller.inaService.GetById(int64(idIna))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller InaController) GetByToken(c echo.Context) error {
	idTokenIna := c.QueryParam("device_token")
	result := controller.inaService.GetByToken(idTokenIna)

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func NewInaController(db *gorm.DB) InaController {
	service := services.NewInaService(db)
	controller := InaController{
		inaService: service,
		validate:   *v1.New(),
	}

	return controller
}
