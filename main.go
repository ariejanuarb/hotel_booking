package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"hotel_booking/app"
	"hotel_booking/controller"
	"hotel_booking/helper"
	"hotel_booking/middleware"
	"hotel_booking/repository"
	"hotel_booking/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	floorRepository := repository.NewFloorRepository()
	floorService := service.NewFloorService(floorRepository, db, validate)
	floorController := controller.NewFloorController(floorService)

	meetingRoomRepository := repository.NewMeetingRoomRepository()
	meetingRoomService := service.NewMeetingRoomService(meetingRoomRepository, db, validate)
	meetingRoomController := controller.NewMeetingRoomController(meetingRoomService)

	rolesRepository := repository.NewRolesRepository()
	rolesService := service.NewRolesService(rolesRepository, db, validate)
	rolesController := controller.NewRolesController(rolesService)

	router := app.NewRouter(floorController, meetingRoomController, rolesController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
