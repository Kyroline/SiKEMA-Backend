package updateExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateExcuseRepository(*model.Excuse) (*model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewUpdateExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateExcuseRepository(input *model.Excuse) (*model.Excuse, string) {
	db := r.db.Model(input)

	var count int64
	if err := db.Where("id = ?", input.AbsentID).Count(&count).Error; err != nil {
		return nil, "UPDATE_EXCUSE_NOTFOUND_404"
	}

	if err := db.Save(&input).Error; err != nil {
		return nil, "UPDATE_EXCUSE_INTERNAL_500"
	}
	return input, ""
}
