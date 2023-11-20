package course

import service "attendance-is/services"

type handler struct {
	service service.CourseService
}

func NewCourseHandler(service service.CourseService) handler {
	return handler{service: service}
}
