package updateExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateExcuseRepository(input *model.Excuse) (*model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewUpdateExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateExcuseRepository(input *model.Excuse) (*model.Excuse, string) {
	return nil, ""
}
