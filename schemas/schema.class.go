package schema

type GetStudentMeta struct {
	Count       int64
	Page        int64
	ItemPerPage int64
}

type CreateClassRequest struct {
	Name     string
	Students []uint
}

type UpdateClassRequest struct {
	ID       uint
	Name     string
	Students []uint
}

type AddStudentToClassRequest struct {
	ID       uint
	Students []uint
}

type RemoveStudentFromClassRequest struct {
	ID      uint
	Student uint
}

type UpdateStudentFromClassRequest struct {
	ID       uint
	Students []uint
}
