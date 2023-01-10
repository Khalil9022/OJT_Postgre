package checklistpencairan

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

func (h *Handler) GetDataBranch(c *gin.Context) {
	databranch, err := h.Service.GetDataBranch()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Mengambil Data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sukses Mengambil data branch :D",
		"data":    databranch,
	})
}

func (h *Handler) GetDataCompany(c *gin.Context) {
	datacompany, err := h.Service.GetDataCompany()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Mengambil Data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sukses Mengambil data company :D",
		"data":    datacompany,
	})
}
