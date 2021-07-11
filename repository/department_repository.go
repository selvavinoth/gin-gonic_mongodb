package repository

import (
	"time"

	"github.com/selvavinoth/gin-gonic/db"
	"github.com/selvavinoth/gin-gonic/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DepartmentsCollection = "departments"

type DepartmentRepository interface {
	GetAll() ([]*models.Department, error)
	Create(department *models.Department) (*models.Department, error)
	GetById(id string) (*models.Department, error)
}
type departmentRepository struct {
	c *mgo.Collection
}

func NewDepartmentsRepository(conn db.Connection) DepartmentRepository {
	return &departmentRepository{c: conn.DB().C(DepartmentsCollection)}
}

func (r *departmentRepository) GetAll() (departments []*models.Department, err error) {
	err = r.c.Find(bson.M{}).All(&departments)
	return departments, err
}

func (r *departmentRepository) GetById(id string) (department *models.Department, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&department)
	return department, err
}

func (r *departmentRepository) Create(department *models.Department) (*models.Department, error) {
	department.Id = bson.NewObjectId()
	department.CreatedOn = time.Now()
	err := r.c.Insert(&department)
	return department, err
}
