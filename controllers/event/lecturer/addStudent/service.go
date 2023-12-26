package addStudentEvent

type service struct {
	repository Repository
}

func NewAddStudentToEventService(repository Repository) *service {
	return &service{repository: repository}
}

func AddStudentToEventService() {

}
