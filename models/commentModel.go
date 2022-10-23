package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	TimeModel
	UserID  uint
	User    *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PhotoID uint   `json:"photo_id" form:"photo_id"`
	Photo   *Photo `json:"photo" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Message string `json:"message" form:"message" valid:"required~Comment is required"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
