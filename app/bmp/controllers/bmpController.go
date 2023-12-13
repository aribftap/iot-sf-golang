package controllers

import (
	"fmt"
	"iot-golang/app/bmp/models"
	"iot-golang/app/bmp/services"
	"iot-golang/app/helpers"
	"net/http"
	"reflect"
	"strconv"

	v1 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BmpController struct {
	bmpService services.BmpService
	validate   v1.Validate
}

var httpStatus int

func (controller BmpController) Create(c echo.Context) error {
	type payload struct {
		DeviceToken     string `json:"device_token" validate:"required"`
		TekananUdara    string `json:"tekanan_udara" validate:"required"`
		TinggiPermukaan string `json:"tinggi_permukaan" validate:"required"`
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

	result := controller.bmpService.Create(models.Bmp{DeviceToken: payloadValidator.DeviceToken, TekananUdara: payloadValidator.TekananUdara, TinggiPermukaan: payloadValidator.TinggiPermukaan, Battery: payloadValidator.Battery})

	if result.Status == 201 {
		httpStatus = http.StatusCreated
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller BmpController) Update(c echo.Context) error {
	type payload struct {
		TekananUdara    string `json:"tekanan_udara" validate:"required"`
		TinggiPermukaan string `json:"tinggi_permukaan" validate:"required"`
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

	idIna, _ := strconv.Atoi(c.Param("id"))
	result := controller.bmpService.Update(int64(idIna), models.Bmp{TekananUdara: payloadValidator.TekananUdara, TinggiPermukaan: payloadValidator.TinggiPermukaan, Battery: payloadValidator.Battery})

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller BmpController) Delete(c echo.Context) error {
	idBmp, _ := strconv.Atoi(c.Param("id"))
	result := controller.bmpService.Delete(int64(idBmp))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller BmpController) GetAll(c echo.Context) error {
	result := controller.bmpService.GetAll()

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller BmpController) GetById(c echo.Context) error {
	idBmp, _ := strconv.Atoi(c.QueryParam("id"))
	result := controller.bmpService.GetById(int64(idBmp))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller BmpController) GetByToken(c echo.Context) error {
	idTokenBmp := c.QueryParam("device_token")
	result := controller.bmpService.GetByToken(idTokenBmp)

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func NewBmpController(db *gorm.DB) BmpController {
	service := services.NewBmpService(db)
	controller := BmpController{
		bmpService: service,
		validate:   *v1.New(),
	}

	return controller
}
