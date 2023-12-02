package getExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetExcuseRepository(student model.Student) (*model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetExcuseRepository(student model.Student) (*model.Excuse, string) {
	var excuse model.Excuse
	db := r.db.Model(excuse)
	db.Preload("Absent").Where("absents.student_id = ?", student.ID).Find(&excuse)
	return &excuse, ""
}
