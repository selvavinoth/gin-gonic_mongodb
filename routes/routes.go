package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selvavinoth/gin-gonic/controller"
	"github.com/selvavinoth/gin-gonic/db"
	"github.com/selvavinoth/gin-gonic/service"

	"github.com/selvavinoth/gin-gonic/repository"
)

func SetupRouter(conn db.Connection) {
	router := gin.Default()
	departmentRepository := repository.NewDepartmentsRepository(conn)
	departmentService := service.NewDepartmentService(departmentRepository)
	departmentController := controller.NewDepartmentController(departmentService)
	v1 := router.Group("/v1/api")
	{

		v1.GET("department", departmentController.GetDepartments)
		v1.GET("department/:id", departmentController.GetDepartment)
		v1.POST("department", departmentController.CreateDepartment)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8000")
}
