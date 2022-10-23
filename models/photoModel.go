package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	TimeModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID   uint   `json:"user_id"`
	User     *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
