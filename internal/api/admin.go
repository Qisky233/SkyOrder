package api

import (
	"net/http"
	"order/internal/model"
	"order/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService *service.AdminService
}

func NewAdminHandler(adminService *service.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Captcha  string `json:"captcha" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1001,
			"message": "参数错误",
		})
		return
	}

	// TODO: 验证码验证

	token, admin, err := h.adminService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    1007,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
			"admin": admin,
		},
	})
}

func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var req struct {
		Username string             `json:"username" binding:"required"`
		Password string             `json:"password" binding:"required"`
		Role     string             `json:"role" binding:"required"`
		Profile  *model.Sys_profile `json:"profile"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1001,
			"message": "参数错误",
		})
		return
	}

	if err := h.adminService.CreateAdmin(req.Username, req.Password, req.Role, req.Profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
	})
}

func (h *AdminHandler) GetAdminList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	role := c.Query("role")

	admins, total, err := h.adminService.GetAdminList(page, size, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"list":  admins,
			"total": total,
		},
	})
}

func (h *AdminHandler) GetAdminProfile(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	profile, err := h.adminService.GetAdminProfileByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": profile,
	})
}

func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var req struct {
		Role    string             `json:"role"`
		Profile *model.Sys_profile `json:"profile"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1001,
			"message": "参数错误",
		})
		return
	}

	if err := h.adminService.UpdateAdmin(uint(id), req.Role, req.Profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1005,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

func (h *AdminHandler) DeleteAdmin(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.adminService.DeleteAdmin(uint(id)); err != nil {
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
