package model

type Class struct {
	ID       uint      `json:"id,omitempty" gorm:"primaryKey"`
	Name     string    `json:"name,omitempty" gorm:"type:varchar(25);unique"`
	Students []Student `json:"students,omitempty" gorm:"foreignKey:ClassID"`
	Courses  []Course  `json:"courses,omitempty" gorm:"many2many:enrollments"`
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
