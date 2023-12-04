package service

import (
	model "attendance-is/models"
	schema "attendance-is/schemas"
	"errors"

	"gorm.io/gorm"
)

type EventService struct {
	DB *gorm.DB
}

func NewEventService(db *gorm.DB) EventService {
	return EventService{DB: db}
}

func (s *EventService) GetEvent() (*[]model.Event, error) {
	var event []model.Event
	if err := s.DB.Model(&event).Preload("Class").Preload("Course").Preload("Students").Preload("Lecturer").Find(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (s *EventService) FindEvent(id uint) (*model.Event, error) {
	var event model.Event
	if err := s.DB.Model(&event).Preload("Class").Preload("Course").Preload("Students").Preload("Lecturer").Where("id = ?", id).Find(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func (s *EventService) CreateEvent(data schema.CreateEventRequest) (*model.Event, error) {
	newEvent := model.Event{LecturerID: data.LecturerId, CourseID: data.CourseId, ClassID: data.ClassId, Meet: data.Meet, Status: 0}
	if err := s.DB.Create(&newEvent).Error; err != nil {
		return nil, err
	}

	return &newEvent, nil
}

func (s *EventService) AddStudentToEvent(data schema.AddStudentToEventRequest) (*model.Event, error) {
	var students []model.Student
	if err := s.DB.Where("nim IN ?", data.StudentId).Find(&students).Error; err != nil {
		return nil, err
	}

	var event model.Event
	if err := s.DB.Where("id = ?", data.EventId).Find(&event).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Model(&event).Association("Students").Append(&students); err != nil {
		return nil, err
	}

	return &event, nil
}

func (s *EventService) FinalizeEvent(data schema.FinalizeEventRequest) (*model.Event, error) {
	var event model.Event

	var absentStudent []model.Student

	s.DB.Model(model.Event{}).Preload("Class.Students").Preload("Students").Where("id = ?", data.EventId).Find(&event)

	if event.Status == 2 {
		return nil, errors.New("Conflict")
	}

	var IDs []uint
	for _, element := range event.Students {
		IDs = append(IDs, element.ID)
	}

	s.DB.Where("class_id = ?", event.ClassID).Where("id NOT IN ?", IDs).Find(&absentStudent)

	if err := s.DB.Transaction(func(tx *gorm.DB) error {
		for _, element := range absentStudent {
			tx.Create(&model.Absent{Student: element, Event: event})
		}
		return nil
	}); err != nil {
		return nil, err
	}

	event.Status = 2
	if err := s.DB.Save(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}
