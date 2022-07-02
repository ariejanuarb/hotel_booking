package app

import (
	"github.com/julienschmidt/httprouter"
	"hotel_booking/controller"
	"hotel_booking/exception"
)

func NewRouter(floorController controller.FloorController, meetingRoomController controller.MeetingRoomController, rolesController controller.RolesController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/floors", floorController.FindAll)
	router.GET("/api/floors/:floorId", floorController.FindById)
	router.POST("/api/floors", floorController.Create)
	router.PUT("/api/floors/:floorId", floorController.Update)
	router.DELETE("/api/floors/:floorId", floorController.Delete)

	router.GET("/api/meetingrooms", meetingRoomController.FindAll)
	router.GET("/api/meetingrooms/:meetingRoomId", meetingRoomController.FindById)
	router.POST("/api/meetingrooms", meetingRoomController.Create)
	router.PUT("/api/meetingrooms/:meetingRoomId", meetingRoomController.Update)
	router.DELETE("/api/meetingrooms/:meetingRoomId", meetingRoomController.Delete)

	router.GET("/api/roles", rolesController.FindAll)
	router.GET("/api/roles/:rolesId", rolesController.FindById)
	router.POST("/api/roles", rolesController.Create)
	router.PUT("/api/roles/:rolesId", rolesController.Update)
	router.DELETE("/api/roles/:rolesId", rolesController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
