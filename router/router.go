package router

import (
	"curd-api/controller"

	"github.com/gin-gonic/gin"
)

// InitRouter() : is used to initialize the routes for the application
// func InitRouter() {
// 	router := gin.Default()
// 	//add new routes here
// 	router.GET("/user", controller.GetAllUser())
// 	router.POST("/user", controller.InsertUser())
// 	router.GET("/user/:id", controller.GetDataByID())
// 	router.POST("/user/update", controller.UpdateById())
// 	router.DELETE("/:id", controller.DeleteByID())
// 	http.ListenAndServe(":5000", router)

// }

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/user", controller.GetAllUser())

	r.GET("/user/:id", controller.GetDataByID())
	r.POST("/user", controller.InsertUser())
	r.POST("user/update", controller.UpdateById())
	r.DELETE("/user/:id", controller.DeleteByID())
	return r
}
