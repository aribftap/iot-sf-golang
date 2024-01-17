package controllers

import (
	"fmt"
	"iot-golang/internal/helpers"
	"iot-golang/internal/pzem/models"
	"iot-golang/internal/pzem/services"
	"net/http"
	"reflect"
	"strconv"

	v1 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PzemController struct {
	pzemService services.PzemService
	validate    v1.Validate
}

var httpStatus int

func (controller PzemController) Create(c echo.Context) error {
	type payload struct {
		DeviceToken string `json:"device_token" validate:"required"`
		Tegangan    string `json:"tegangan" validate:"required"`
		Arus        string `json:"arus" validate:"required"`
		Daya        string `json:"daya" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
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

	result := controller.pzemService.Create(models.Pzem{DeviceToken: payloadValidator.DeviceToken, Tegangan: payloadValidator.Tegangan, Arus: payloadValidator.Arus, Daya: payloadValidator.Daya})

	if result.Status == 201 {
		httpStatus = http.StatusCreated
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller PzemController) Update(c echo.Context) error {
	type payload struct {
		Tegangan string `json:"tegangan" validate:"required"`
		Arus     string `json:"arus" validate:"required"`
		Daya     string `json:"daya" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
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

	idPzem, _ := strconv.Atoi(c.Param("id"))
	result := controller.pzemService.Update(int64(idPzem), models.Pzem{Tegangan: payloadValidator.Tegangan, Arus: payloadValidator.Arus, Daya: payloadValidator.Daya})

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller PzemController) Delete(c echo.Context) error {
	idPzem, _ := strconv.Atoi(c.Param("id"))
	result := controller.pzemService.Delete(int64(idPzem))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller PzemController) GetAll(c echo.Context) error {
	result := controller.pzemService.GetAll()
	return c.JSON(http.StatusOK, result)
}

func (controller PzemController) GetById(c echo.Context) error {
	idPzem, _ := strconv.Atoi(c.QueryParam("id"))
	result := controller.pzemService.GetById(int64(idPzem))

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func (controller PzemController) GetByToken(c echo.Context) error {
	idTokenPzem := c.QueryParam("device_token")
	result := controller.pzemService.GetByToken(idTokenPzem)

	if result.Status == 200 {
		httpStatus = http.StatusOK
	} else if result.Status == 404 {
		httpStatus = http.StatusNotFound
	} else if result.Status == 500 {
		httpStatus = http.StatusInternalServerError
	}

	return c.JSON(httpStatus, result)
}

func NewPzemController(db *gorm.DB) PzemController {
	service := services.NewPzemService(db)
	controller := PzemController{
		pzemService: service,
		validate:    *v1.New(),
	}

	return controller
}
