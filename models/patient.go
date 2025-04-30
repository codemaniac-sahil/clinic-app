// models/patient.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	Gender           string    `json:"gender"`
	Address          string    `json:"address"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	MedicalHistory   string    `json:"medical_history"`
	BloodType        string    `json:"blood_type"`
	EmergencyContact string    `json:"emergency_contact"`
	CreatedBy        uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedBy        uuid.UUID `gorm:"type:uuid" json:"updated_by"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// BeforeCreate is a GORM hook that runs before creating a new record
func (p *Patient) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
