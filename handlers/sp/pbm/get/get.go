package getSPHandler

import (
	getSP "attendance-is/controllers/sp/pbm/get"
	util "attendance-is/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getSP.Service
}

func NewGetSPHandler(service getSP.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetSPHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	input := getSP.InputGetSP{
		ID: uint(id),
	}

	res, err := h.service.GetSPService(input)
	switch err {
	case "ERR_CONFLICT_409":
		util.ErrorRespose(c, http.StatusConflict, "SP ID already exist")
		return
	default:
		util.APIResponse(c, http.StatusOK, res, nil)
		return
	}
}
