package pencairanreport

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

func (h *Handler) GetAllCustomerAs9(c *gin.Context) {
	datacustomer, err := h.Service.GetAllCustomerAs9()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Mengambil Data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sukses Mengambil data Customer :D",
		"data":    datacustomer,
	})
}

func (h *Handler) GetSpesifikCustomerAs9(c *gin.Context) {
	var req DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}

	datacustomer, err := h.Service.GetSpesifikCustomerAs9(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Mengambil Data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sukses Mengupdate data Customer :D",
		"data":    datacustomer,
	})

}

func (h *Handler) GetAllCustomerAs01(c *gin.Context) {
	datacustomer, err := h.Service.GetAllCustomerAs01()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Mengambil Data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sukses Mengambil data Customer :D",
		"data":    datacustomer,
	})
}

func (h *Handler) GetSpesifikCustomerAs01(c *gin.Context) {
	var req DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}

	datacustomer, err := h.Service.GetSpesifikCustomerAs01(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Mengambil Data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sukses Mengupdate data Customer :D",
		"data":    datacustomer,
	})

}

func (h *Handler) UpdateApprovalStatus(c *gin.Context) {

	var req ReqPpk

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}

	err := h.Service.UpdateApprovalStatus(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Mengambil Data",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sukses mengupdate approval status :D",
	})
}
