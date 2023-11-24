package schema

type CreateEventRequest struct {
	LecturerId uint `json:"lecturer_id"`
	CourseId   uint `json:"course_id"`
	ClassId    uint `json:"class_id"`
	Meet       uint `json:"meet"`
}

type FinalizeEventRequest struct {
	EventId uint
}

type AddStudentToEventRequest struct {
	EventId   uint
	StudentId []string `json:"student_id"`
}
