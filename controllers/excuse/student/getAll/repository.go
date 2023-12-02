package getAllExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllExcuseRepository(student model.Student) (*[]model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetAllExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAllExcuseRepository(student model.Student) (*[]model.Excuse, string) {
	var excuses []model.Excuse
	db := r.db.Model(excuses)
	db.Preload("Absent", "id = ?", student.ID).Find(&excuses)
	return &excuses, ""
}
