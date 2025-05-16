package model

import "time"

// Sys_account 管理员表
type Sys_account struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Username  string    `gorm:"size:50;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null" json:"-"`
	Role      string    `gorm:"size:20;not null" json:"role"` // system_admin, normal_admin
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Sys_profile 管理员个人信息表
type Sys_profile struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	AdminID   uint      `gorm:"not null" json:"admin_id"`
	Name      string    `gorm:"size:50" json:"name"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Email     string    `gorm:"size:100" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Demand 需求表
type Demand struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"type:text;not null" json:"title"`
	Descs       string    `gorm:"type:text" json:"descs"`
	Images      string    `gorm:"type:text" json:"images"`
	Attachments string    `gorm:"type:text" json:"attachments"`
	Price       string    `gorm:"type:text" json:"price"`
	Status      string    `gorm:"type:text" json:"status"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Customer 用户表
type Customer struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	Phone     string    `gorm:"type:text" json:"phone"`
	QQ        string    `gorm:"type:text" json:"qq"`
	WX        string    `gorm:"type:text" json:"wx"`
	Total     string    `gorm:"type:text" json:"total"`
	CreatedAt time.Time `json:"created_at"`
}
