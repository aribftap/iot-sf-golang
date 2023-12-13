package main

import (
	beitianController "iot-golang/app/beitian/controllers"
	bmpController "iot-golang/app/bmp/controllers"
	inaController "iot-golang/app/ina/controllers"
	pzemController "iot-golang/app/pzem/controllers"
	thigrowController "iot-golang/app/thigrow/controllers"
	thmController "iot-golang/app/thm/controllers"
	"iot-golang/config"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	config.LoadEnv()
}

func main() {
	db := config.InitDB()

	route := echo.New()
	apiIoTSf := route.Group("api/iot-sf/")

	// route for sensor beitian220
	beitianController := beitianController.NewBeitianController(db)
	apiIoTSf.POST("beitian/create", beitianController.Create)
	apiIoTSf.GET("beitian/get_all", beitianController.GetAll)
	apiIoTSf.GET("beitian/detail", beitianController.GetById)
	apiIoTSf.GET("beitian/detail/token", beitianController.GetByToken)
	apiIoTSf.GET("beitian/detail/new/token", beitianController.GetNewByToken)
	apiIoTSf.PUT("beitian/update/:id", beitianController.Update)
	apiIoTSf.DELETE("beitian/delete/:id", beitianController.Delete)

	// route for sensor bmp180
	bmpController := bmpController.NewBmpController(db)
	apiIoTSf.POST("bmp/create", bmpController.Create)
	apiIoTSf.GET("bmp/get_all", bmpController.GetAll)
	apiIoTSf.GET("bmp/detail", bmpController.GetById)
	apiIoTSf.GET("bmp/detail/token", bmpController.GetByToken)
	apiIoTSf.PUT("bmp/update/:id", bmpController.Update)
	apiIoTSf.DELETE("bmp/delete/:id", bmpController.Delete)

	// route for sensor ina219
	inaController := inaController.NewInaController(db)
	apiIoTSf.POST("ina/create", inaController.Create)
	apiIoTSf.GET("ina/get_all", inaController.GetAll)
	apiIoTSf.GET("ina/detail", inaController.GetById)
	apiIoTSf.GET("ina/detail/token", inaController.GetByToken)
	apiIoTSf.PUT("ina/update/:id", inaController.Update)
	apiIoTSf.DELETE("ina/delete/:id", inaController.Delete)

	// route for sensor pzem
	pzemController := pzemController.NewPzemController(db)
	apiIoTSf.POST("pzem/create", pzemController.Create)
	apiIoTSf.GET("pzem/get_all", pzemController.GetAll)
	apiIoTSf.GET("pzem/detail", pzemController.GetById)
	apiIoTSf.GET("pzem/detail/token", pzemController.GetByToken)
	apiIoTSf.PUT("pzem/update/:id", pzemController.Update)
	apiIoTSf.DELETE("pzem/delete/:id", pzemController.Delete)

	// route for sensor thigrow
	thigrowController := thigrowController.NewThigrowController(db)
	apiIoTSf.POST("thigrow/create", thigrowController.Create)
	apiIoTSf.GET("thigrow/get_all", thigrowController.GetAll)
	apiIoTSf.GET("thigrow/detail", thigrowController.GetById)
	apiIoTSf.GET("thigrow/detail/token", thigrowController.GetByToken)
	apiIoTSf.PUT("thigrow/update/:id", thigrowController.Update)
	apiIoTSf.DELETE("thigrow/delete/:id", thigrowController.Delete)

	// route for sensor thm30d
	thmController := thmController.NewThmController(db)
	apiIoTSf.POST("thm/create", thmController.Create)
	apiIoTSf.GET("thm/get_all", thmController.GetAll)
	apiIoTSf.GET("thm/detail", thmController.GetById)
	apiIoTSf.GET("thm/detail/token", thmController.GetByToken)
	apiIoTSf.PUT("thm/update/:id", thmController.Update)
	apiIoTSf.DELETE("thm/delete/:id", thmController.Delete)

	route.Start(":" + os.Getenv("PORT"))
}
