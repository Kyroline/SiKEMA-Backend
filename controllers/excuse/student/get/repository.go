package getExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetExcuseRepository(input InputGetExcuse) (*model.Excuse, string)
	GetExcusesRepository(input InputGetExcuse) (*model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetExcuseRepository(input InputGetExcuse) (*model.Excuse, string) {
	var excuse model.Excuse
	db := r.db.Model(excuse)
	db.Preload("Absent", "student_id = ?", input.StudentID).Where("id = ?", input.ID).Find(&excuse)
	return &excuse, ""
}

func (r *repository) GetExcusesRepository(input InputGetExcuse) (*model.Excuse, string) {
	var excuse model.Excuse
	db := r.db.Model(excuse)
	db.Preload("Absent", "student_id = ?", input.StudentID).Where("id = ?", input.ID).Find(&excuse)
	return &excuse, ""
}
