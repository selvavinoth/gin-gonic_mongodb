package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selvavinoth/gin-gonic/models"
	"github.com/selvavinoth/gin-gonic/service"
)

type DepartmentController interface {
	GetDepartments(c *gin.Context)
	GetDepartment(c *gin.Context)
	CreateDepartment(c *gin.Context)
}
type departmentController struct {
	departmentService service.DepartmentService
}

func NewDepartmentController(s service.DepartmentService) DepartmentController {
	return departmentController{
		departmentService: s,
	}
}

func (d departmentController) GetDepartments(c *gin.Context) {
	departments, err := d.departmentService.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": departments})
	}
}

func (d departmentController) CreateDepartment(c *gin.Context) {
	log.Printf("starting")
	var model models.Department
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department, err := d.departmentService.Create(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": department})
	}
}

func (d departmentController) GetDepartment(c *gin.Context) {
	id := c.Param("id")
	department, err := d.departmentService.GetById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": department})
	}
}
