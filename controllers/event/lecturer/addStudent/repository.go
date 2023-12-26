package addStudentEvent

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetStudentByIDs(IDs []string) (*[]model.Student, string)
	AddStudentToEvent(event *model.Event, students *[]model.Student)
}

type repository struct {
	db *gorm.DB
}

func NewAddStudentEventRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetStudentByIDs(IDs []string) (*[]model.Student, string) {
	var student []model.Student
	err := r.db.Where("id IN ?", IDs).Find(&student)
	if err.RowsAffected == 0 {
		return nil, "STUDENT_NOTFOUND_404"
	}
	return &student, ""
}

func (r *repository) AddStudentToEvent(event *model.Event, students *[]model.Student) string {
	if err := r.db.Model(&event).Association("Students").Append(&students); err != nil {
		return err.Error()
	}

	return ""
}
