package event

import service "attendance-is/services"

type handler struct {
	service service.EventService
}

func NewEventHandler(service service.EventService) handler {
	return handler{service: service}
}
