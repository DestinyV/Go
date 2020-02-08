package models

import "time"

type LotteryRecords struct {
	ID         int        `gorm:"primary_key,auto_increment" json:"id"`
	OpenID     string     `gorm:"column:openid" json:"-"`
	RewardID   int        `json:"-"`
	LuckyNum   string     `json:"luckyNum"`
	CreateTime *time.Time `grom:"current_timestamp()" json:"createTime"`
	User       User       `gorm:"foreignKey:OpenID" json:"user"`
	Reward     Reward     `gorm:"foreignKey:RewardID" json:"reward"`
}
