package getEvent

import (
	model "attendance-is/models"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func NewGetStudentEventRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetEventByID(ID string) (*model.Event, string) {
	var event model.Event
	if err := r.db.Model(&event).Preload("Class").Preload("Course").Preload("Students").Preload("Lecturer").Where("id = ?", ID).Take(&event).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "STUDENT_NOTFOUND_404"
		}
		return nil, err.Error()
	}

	return &event, ""
}
