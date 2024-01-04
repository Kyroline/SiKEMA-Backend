package getEvent

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetEvent() (*[]model.Event, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetEventRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetEvent() (*[]model.Event, string) {
	var event []model.Event
	if err := r.db.Model(&event).Preload("Class").Preload("Course").Preload("Students").Preload("Lecturer").Find(&event).Error; err != nil {
		return nil, "EVENT_UNEXPECTED_500 : " + err.Error()
	}

	if len(event) < 1 {
		return nil, "EVENT_NOTFOUND_404"
	}
	return &event, ""
}
