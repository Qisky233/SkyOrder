package service

import (
	"errors"
	"strconv"

	"gorm.io/gorm"

	"order/internal/database"
	"order/internal/model"
)

type DemandService struct{}

func NewDemandService() *DemandService {
	return &DemandService{}
}

func (s *DemandService) CreateDemand(demand *model.Demand) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(demand).Error; err != nil {
			return err
		}

		// 增加对应的 customer 的 total
		if err := s.updateCustomerTotal(demand.UserID, 1, tx); err != nil {
			return err
		}

		return nil
	})
}

func (s *DemandService) GetAllDemands() ([]model.Demand, error) {
	var demands []model.Demand
	result := database.DB.Find(&demands)
	if result.Error != nil {
		return nil, result.Error
	}
	return demands, nil
}

func (s *DemandService) GetDemand(id uint) (*model.Demand, error) {
	var demand model.Demand
	if err := database.DB.First(&demand, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("需求不存在")
		}
		return nil, err
	}
	return &demand, nil
}

func (s *DemandService) UpdateDemand(id uint, demand *model.Demand) error {
	return database.DB.Model(&model.Demand{}).Where("id = ?", id).Updates(demand).Error
}

func (s *DemandService) DeleteDemand(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var demand model.Demand
		// 删除需求
		if err := tx.First(&demand, "id = ?", id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("需求不存在")
			}
			return err
		}

		if err := tx.Delete(&model.Demand{}, id).Error; err != nil {
			return err
		}

		// 获取对应的 customer
		var customer model.Customer
		if err := tx.Where("id = ?", demand.UserID).First(&customer).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("用户不存在")
			}
			return err
		}

		// 减少对应的 customer 的 total
		if err := s.updateCustomerTotal(customer.ID, -1, tx); err != nil {
			return err
		}

		return nil
	})
}

func (s *DemandService) updateCustomerTotal(customerID uint, delta int, tx *gorm.DB) error {
	var customer model.Customer
	if err := tx.Where("id = ?", customerID).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	// 更新 total
	parsedTotal, err := strconv.Atoi(customer.Total)
	if err != nil {
		return errors.New("invalid total value")
	}
	customer.Total = strconv.Itoa(parsedTotal + delta)
	if err := tx.Save(&customer).Error; err != nil {
		return err
	}

	return nil
}
