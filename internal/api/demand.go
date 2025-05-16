package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"order/internal/model"
	"order/internal/service"
)

type DemandHandler struct {
	demandService *service.DemandService
}

func NewDemandHandler(demandService *service.DemandService) *DemandHandler {
	return &DemandHandler{
		demandService: demandService,
	}
}

func (h *DemandHandler) CreateDemand(c *gin.Context) {
	var demand model.Demand
	if err := c.ShouldBindJSON(&demand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1001,
			"message": "参数错误",
		})
		return
	}

	if err := h.demandService.CreateDemand(&demand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": demand,
	})
}

func (h *DemandHandler) GetDemand(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	demand, err := h.demandService.GetDemand(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": demand,
	})
}

func (h *DemandHandler) GetAllDemands(c *gin.Context) {
	demands, err := h.demandService.GetAllDemands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": demands,
	})
}

func (h *DemandHandler) UpdateDemand(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var demand model.Demand
	if err := c.ShouldBindJSON(&demand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1001,
			"message": "参数错误",
		})
		return
	}

	if err := h.demandService.UpdateDemand(uint(id), &demand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func (h *DemandHandler) DeleteDemand(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.demandService.DeleteDemand(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
