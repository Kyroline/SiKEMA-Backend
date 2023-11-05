package createEvent

type Input struct {
	// LecturerId uint will be automatically inserted using jwt token
	Id        uint
	ClassId   uint
	Meet      uint
	StudentId []uint
	Status    bool
}
