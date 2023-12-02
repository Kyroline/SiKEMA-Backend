package getAllExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllExcuseRepository() (*[]model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetAllExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAllExcuseRepository() (*[]model.Excuse, string) {
	var excuse []model.Excuse
	db := r.db.Model(excuse)

	if err := db.Find(&excuse).Error; err != nil {
		return nil, "GETALL_EXCUSE_INTERNAL_500"
	}

	return &excuse, ""
}
