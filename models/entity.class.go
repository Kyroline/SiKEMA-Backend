package model

type Class struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"type:varchar(25);unique"`
	Students []Student `gorm:"foreignKey:ClassID" json:"students,omitempty"`
	Courses  []Course  `gorm:"many2many:enrollments" json:"courses,omitempty"`
}

// func (c *Class) ToGetClassResponse() schema.GetClassResponse {
// 	return schema.GetClassResponse{
// 		ID:       c.ID,
// 		Name:     c.Name,
// 		Students: c.StudentsToResponse(),
// 	}
// }

// func (s *Student) ToStudentResponse() schema.StudentResponse {
// 	return schema.StudentResponse{
// 		ID:   s.ID,
// 		Nim:  s.Nim,
// 		Name: s.Name,
// 	}
// }

// func (c *Class) StudentsToResponse() []schema.StudentResponse {
// 	var studentsResponse []schema.StudentResponse
// 	for _, student := range c.Students {
// 		studentsResponse = append(studentsResponse, student.ToStudentResponse())
// 	}
// 	return studentsResponse
// }
