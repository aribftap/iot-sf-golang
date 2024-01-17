package controllers

import (
	"fmt"
	"iot-golang/internal/helpers"
	"iot-golang/internal/thm/models"
	"iot-golang/internal/thm/services"
	"net/http"
	"reflect"
	"strconv"

	v1 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ThmController struct {
	thmService services.ThmService
	validate   v1.Validate
}

var httpStatus int

func (controller ThmController) Create(c echo.Context) error {
	type payload struct {
		DeviceToken     string `json:"device_token" validate:"required"`
		Temperature     string `json:"temperature" validate:"required"`
		KelembabanUdara string `json:"kelembaban_udara" validate:"required"`
		Battery         string `json:"battery" validate:"required"`
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

	result := controller.thmService.Create(models.Thm{DeviceToken: payloadValidator.DeviceToken, Temperature: payloadValidator.Temperature, KelembabanUdara: payloadValidator.KelembabanUdara, Battery: payloadValidator.Battery})

	if result.Status == 201 {
		httpStatus = http.StatusCreated
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller ThmController) Update(c echo.Context) error {
	type payload struct {
		Temperature     string `json:"temperature" validate:"required"`
		KelembabanUdara string `json:"kelembaban_udara" validate:"required"`
		Battery         string `json:"battery" validate:"required"`
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

	idThm, _ := strconv.Atoi(c.Param("id"))
	result := controller.thmService.Update(int64(idThm), models.Thm{Temperature: payloadValidator.Temperature, KelembabanUdara: payloadValidator.KelembabanUdara, Battery: payloadValidator.Battery})

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	}

	return c.JSON(httpStatus, result)
}

func (controller ThmController) Delete(c echo.Context) error {
	idThm, _ := strconv.Atoi(c.Param("id"))
	result := controller.thmService.Delete(int64(idThm))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	}

	return c.JSON(httpStatus, result)
}

func (controller ThmController) GetAll(c echo.Context) error {
	result := controller.thmService.GetAll()

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller ThmController) GetById(c echo.Context) error {
	idThm, _ := strconv.Atoi(c.QueryParam("id"))
	result := controller.thmService.GetById(int64(idThm))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}
	return c.JSON(httpStatus, result)
}

func (controller ThmController) GetByToken(c echo.Context) error {
	idTokenBmp := c.QueryParam("device_token")
	result := controller.thmService.GetByToken(idTokenBmp)

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	}

	return c.JSON(httpStatus, result)
}

func NewThmController(db *gorm.DB) ThmController {
	service := services.NewThmService(db)
	controller := ThmController{
		thmService: service,
		validate:   *v1.New(),
	}

	return controller
}
