package addStudentEvent

type AddStudentToEventRequest struct {
	EventId   uint
	StudentId []string `json:"student_id"`
}
