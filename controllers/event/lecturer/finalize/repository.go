package finalizeEvent

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetStudentByWhere(query interface{}, args ...interface{}) (*[]model.Student, string)
	GetEventByID(ID string) (*model.Event, string)
	CreateAbsent(absent model.Absent) (*model.Absent, string)
}

type repository struct {
	db *gorm.DB
}

func NewFinalizeEventRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetStudentByWhere(query interface{}, args ...interface{}) (*[]model.Student, string) {
	var student []model.Student
	err := r.db.Where(query, args...).Find(&student)
	if err.RowsAffected == 0 {
		return nil, "STUDENT_NOTFOUND_404"
	}
	return &student, ""
}

func (r *repository) GetEventByID(ID string) (*model.Event, string) {
	var event model.Event
	if err := r.db.Model(&event).Preload("Course").Preload("Students").Where("id = ?", ID).Take(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "EVENT_NOTFOUND_404"
		}
		return nil, "EVENT_UNEXPECTED_500 : " + err.Error()
	}
	return &event, ""
}

func (r *repository) CreateAbsent(absent model.Absent) (*model.Absent, string) {
	if err := r.db.Create(&absent).Error; err != nil {
		if err == gorm.ErrDuplicatedKey {
			return nil, "ABSENT_CONFLICT_409"
		}
		return nil, err.Error()
	}

	return &absent, ""
}
