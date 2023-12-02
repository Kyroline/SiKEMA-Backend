package service

import (
	model "attendance-is/models"
	schema "attendance-is/schemas"

	"gorm.io/gorm"
)

type AbsentService struct {
	DB *gorm.DB
}

func NewAbsentService(db *gorm.DB) AbsentService {
	return AbsentService{DB: db}
}

func (s *AbsentService) GetAbsentsByStudent(meta *schema.Metadata, studentID int) (*[]model.Absent, error) {
	var data []model.Absent
	if err := s.DB.Model(&[]model.Absent{}).Where("student_id = ?", studentID).Preload("Event").Preload("Event.Course").Preload("Event.Lecturer").Preload("Event.Class").Count(&meta.Count).Limit(meta.ItemPerPage).Offset(meta.ItemPerPage * (meta.Page - 1)).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *AbsentService) GetAbsentByStudent(meta *schema.Metadata, studentID int, ID int) (*model.Absent, error) {
	var data model.Absent
	if err := s.DB.Model(&model.Absent{}).Where("student_id = ?", studentID).Preload("Event").Preload("Event.Course").Preload("Event.Lecturer").Preload("Event.Class").Count(&meta.Count).Limit(meta.ItemPerPage).Offset(meta.ItemPerPage*(meta.Page-1)).Find(&data, ID).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *AbsentService) GetAbsentByLecturer(data schema.GetAbsentByLecturerRequest) (*[]model.Absent, error) {
	var absent []model.Absent
	err := s.DB.Where("event_id = ?", data.EventId).Find(&absent).Error
	if err != nil {
		return nil, err
	}
	return &absent, nil
}

func (s *AbsentService) GetAbsentsByPBM(meta *schema.Metadata) (*[]model.Absent, error) {
	var data []model.Absent
	if err := s.DB.Model(&model.Absent{}).Preload("Event").Preload("Event.Course").Preload("Event.Lecturer").Preload("Event.Class").Count(&meta.Count).Limit(meta.ItemPerPage).Offset(meta.ItemPerPage * (meta.Page - 1)).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *AbsentService) GetAbsentByPBM(meta *schema.Metadata, ID int) (*model.Absent, error) {
	var data model.Absent
	if err := s.DB.Model(&model.Absent{}).Preload("Event").Preload("Event.Course").Preload("Event.Lecturer").Preload("Event.Class").Count(&meta.Count).Limit(meta.ItemPerPage).Offset(meta.ItemPerPage*(meta.Page-1)).Find(&data, ID).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
