package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lottery/errcode"
)

type Reward struct {
	ID        int `gorm:"primary_key,auto_increment"`
	Name      string
	Gradation int `json:"-"`
	Used      int `json:"-"`
}

func (r Reward) TableName() string {
	return "reward"
}

func GetNextRewardHandler(c *gin.Context) error {
	r := &Reward{}
	if err := r.GetNextReward(); err != nil {
		return err
	}
	return errcode.ErrorOK.WithData(r)
}

func (r *Reward) GetNextReward() (err error) {
	err = DB.Where("used = 0").Order("gradation asc", true).First(r).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errcode.ErrorServerErr
	}
	if err == gorm.ErrRecordNotFound {
		return errcode.ErrorEmptyReward
	}
	return
}
