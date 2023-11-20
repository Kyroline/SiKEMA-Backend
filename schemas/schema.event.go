package schema

type CreateEventRequest struct {
	LecturerId uint
	CourseId   uint
	ClassId    uint
	Meet       uint
}

type FinalizeEventRequest struct {
	EventId uint
}

type AddStudentToEventRequest struct {
	EventId   uint
	StudentId []string
}
