package service

import (
	"errors"
	"order/internal/database"
	"order/internal/model"

	"gorm.io/gorm"
)

type CustomerService struct{}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) CreateCustomer(customer *model.Customer) error {
	return database.DB.Create(customer).Error
}

func (s *CustomerService) GetCustomer(id uint) (*model.Customer, error) {
	var customer model.Customer
	if err := database.DB.First(&customer, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &customer, nil
}

func (s *CustomerService) GetAllCustomers() ([]model.Customer, error) {
	var customers []model.Customer
	result := database.DB.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (s *CustomerService) UpdateCustomer(id uint, customer *model.Customer) error {
	return database.DB.Model(&model.Customer{}).Where("id = ?", id).Updates(customer).Error
}

func (s *CustomerService) DeleteCustomer(id uint) error {
	return database.DB.Delete(&model.Customer{}, id).Error
}
