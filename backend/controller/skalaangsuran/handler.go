package skalaangsuran

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GenerateSkalaAngsuran(c *gin.Context) {
	dataCustomer, err := h.Service.GenerateSkalaAngsuran()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Mengambil Data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sukses Generate Data :D",
		"data":    dataCustomer,
	})
}
