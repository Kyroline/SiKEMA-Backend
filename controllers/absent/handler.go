package absent

import service "attendance-is/services"

type handler struct {
	service service.AbsentService
}

func NewAbsentHandler(service service.AbsentService) handler {
	return handler{service: service}
}
