package getAllSPHandler

import (
	getAllSP "attendance-is/controllers/sp/pbm/getAll"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getAllSP.Service
}

func NewGetAllSPHandler(service getAllSP.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAllSPHandler(c *gin.Context) {
	res, err := h.service.GetAllSPService()
	switch err {
	case "ERR_CONFLICT_409":
		util.ErrorRespose(c, http.StatusConflict, "SP ID already exist")
		return
	default:
		util.APIResponse(c, http.StatusOK, res, nil)
		return
	}
}
