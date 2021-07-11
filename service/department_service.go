package service

import (
	"log"

	"github.com/selvavinoth/gin-gonic/models"
	"github.com/selvavinoth/gin-gonic/repository"
)

type DepartmentService interface {
	GetAll() ([]*models.Department, error)
	Create(*models.Department) (*models.Department, error)
	GetById(id string) (*models.Department, error)
}
type departmentService struct {
	departmentRepository repository.DepartmentRepository
}

func NewDepartmentService(r repository.DepartmentRepository) DepartmentService {
	return &departmentService{
		departmentRepository: r,
	}
}

func (u departmentService) GetAll() ([]*models.Department, error) {
	return u.departmentRepository.GetAll()
}
func (u departmentService) GetById(id string) (*models.Department, error) {
	return u.departmentRepository.GetById(id)
}
func (u departmentService) Create(department *models.Department) (*models.Department, error) {
	log.Printf("starting 2")
	return u.departmentRepository.Create(department)
}
