package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"order/internal/database"
	"order/internal/model"
)

type AdminService struct{}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (s *AdminService) Login(username, password string) (string, *model.Sys_account, error) {
	var admin model.Sys_account
	if err := database.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, errors.New("用户名或密码错误")
		}
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return "", nil, errors.New("用户名或密码错误")
	}

	// 生成JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": admin.ID,
		"role":     admin.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("@Inc123456."))

	println(tokenString)
	if err != nil {
		return "", nil, err
	}

	return tokenString, &admin, nil
}

func (s *AdminService) CreateAdmin(username, password, role string, profile *model.Sys_profile) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := &model.Sys_account{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(admin).Error; err != nil {
			return err
		}

		if profile != nil {
			profile.AdminID = admin.ID
			if err := tx.Create(profile).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *AdminService) GetAdminList(page, size int, role string) ([]model.Sys_account, int64, error) {
	var admins []model.Sys_account
	var total int64

	query := database.DB.Model(&model.Sys_account{})
	if role != "" {
		query = query.Where("role = ?", role)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset((page - 1) * size).Limit(size).Find(&admins).Error; err != nil {
		return nil, 0, err
	}

	return admins, total, nil
}

func (s *AdminService) GetAdminProfileByID(adminID uint) (*model.Sys_profile, error) {
	var profile model.Sys_profile
	if err := database.DB.Where("admin_id = ?", adminID).First(&profile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户信息不存在")
		}
		return nil, err
	}
	return &profile, nil
}

func (s *AdminService) UpdateAdmin(id uint, role string, profile *model.Sys_profile) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 更新角色
		if role != "" {
			if err := tx.Model(&model.Sys_account{}).Where("id = ?", id).Update("role", role).Error; err != nil {
				return err
			}
		}

		// 更新或创建管理员个人信息
		if profile != nil {
			profile.AdminID = id

			// 检查是否存在对应的 profile 记录
			var existingProfile model.Sys_profile
			if err := tx.Where("admin_id = ?", id).First(&existingProfile).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// 如果记录不存在，返回错误
					return errors.New("账号不存在")
				} else {
					// 如果发生其他错误，则返回错误
					return err
				}
			}

			// 如果记录存在，则更新现有的 profile
			profile.ID = existingProfile.ID
			if err := tx.Save(profile).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *AdminService) DeleteAdmin(id uint) error {
	// 删除管理员信息
	if err := database.DB.Where("id = ?", id).Delete(&model.Sys_account{}).Error; err != nil {
		return err
	}

	// 删除管理员个人信息
	if err := database.DB.Where("admin_id = ?", id).Delete(&model.Sys_profile{}).Error; err != nil {
		return err
	}

	return nil
}
