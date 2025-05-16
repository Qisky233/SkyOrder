package api

import (
	"net/http"
	"order/internal/model"
	"order/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService *service.CustomerService
}

func NewCustomerHandler(customerService *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1001,
			"message": "参数错误",
		})
		return
	}

	if err := h.customerService.CreateCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": customer,
	})
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	customer, err := h.customerService.GetCustomer(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": customer,
	})
}

func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	customers, err := h.customerService.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": customers,
	})
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1001,
			"message": "参数错误",
		})
		return
	}

	if err := h.customerService.UpdateCustomer(uint(id), &customer); err != nil {
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

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.customerService.DeleteCustomer(uint(id)); err != nil {
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
