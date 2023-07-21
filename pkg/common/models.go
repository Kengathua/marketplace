package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        *string    `gorm:"primary_key;unique;type:uuid;column:id;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `gorm:"autoUpdateTime:false;column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoCreateTime:true;column:updated_at;not null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy uuid.UUID  `gorm:"type:uuid;column:created_by;not null" json:"created_by"`
	UpdatedBy uuid.UUID  `gorm:"type:uuid;column:updated_by;not null" json:"updated_by"`
	IsActive  bool       `json:"is_active"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	if base.ID == nil {
		id := uuid.New().String()
		base.ID = &id
	}

	return nil
}

type BaseModel struct {
	Base
	BusinessPartner string `gorm:"column:business_partner" json:"business_partner"`
}

type BioData struct {
	BaseModel
	Title       string     `json:"title"`
	FirstName   string     `json:"first_name"`
	OtherNames  string     `json:"other_names"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	Gender      string     `json:"gender"`
	DateOfBirth *time.Time `json:"date_of_birth"`
}
