package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lottery/errcode"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type LuckyNumber struct {
	ID       int    `gorm:"primary_key;auto_increment" json:"id"`
	OpenID   string `gorm:"column:openid" json:"-"`
	LuckyNum string `gorm:"default:''" json:"luckyNum"`
}

func (LuckyNumber) TableName() string {
	return "lucky_num"
}

// 获取用户的抽奖号码
func GetLuckyNumByOpenID(c *gin.Context) error {
	var err error
	openid, err := GetUserFromSession(c)
	if err != nil {
		return err
	}
	l := &LuckyNumber{}
	if err = DB.Where(&LuckyNumber{OpenID: openid}).First(&l).Error; gorm.IsRecordNotFoundError(err) {
		l = nil
		return errcode.ErrorOK
	} else if err != nil {
		return errcode.ErrorServerErr
	}
	return errcode.ErrorOK.WithData(l)
}

var lock sync.Mutex // 生成抽奖号码时需要加锁

// 生成抽奖号码
func GenerateLuckyNum(openid string, db *gorm.DB) (l LuckyNumber, err error) {
	lock.Lock()
	defer lock.Unlock()
	var nums []LuckyNumber
	if err = DB.Find(&nums).Error; err != nil {
		return
	}
	if err = checkRepeat(openid, nums); err != nil {
		return
	}
	luckyNum, err := getRandomLuckyNum(nums)
	if err != nil {
		return
	}
	l.LuckyNum = fmt.Sprintf("%03d", luckyNum)
	l.OpenID = openid
	err = db.Create(&l).Error
	return
}

// 检查重复签到
func checkRepeat(openid string, nums []LuckyNumber) error {
	for _, ln := range nums {
		if ln.OpenID == openid {
			return errcode.ErrorAlreadyCheckin
		}
	}
	return nil
}

// 获取一个不和已有记录重复的抽奖号码
func getRandomLuckyNum(nums []LuckyNumber) (num int, err error) {
	num = 800 + rand.Intn(200) // 基数800+200以内的随机数
	for _, ln := range nums {
		n, err := strconv.Atoi(ln.LuckyNum)
		if err != nil {
			return -1, err
		}
		if n == num {
			return getRandomLuckyNum(nums)
		}
	}
	return
}

// 获取幸运号码
func GetLuckyNumber() (l LuckyNumber, err error) {
	var records []LotteryRecords
	err = DB.Find(&records).Error
	if err != nil {
		return
	}
	var luckyNumPool []LuckyNumber
	if len(records) == 0 {
		err = DB.Find(&luckyNumPool).Error
	} else {
		nums := make([]string, len(records))
		for k, v := range records {
			nums[k] = v.LuckyNum
		}
		err = DB.Where("lucky_num NOT IN(?)", nums).Find(&luckyNumPool).Error
	}
	if err != nil {
		return
	}
	if len(luckyNumPool) == 0 {
		return l, errcode.ErrorEmptyReward
	}
	idx := rand.Intn(len(luckyNumPool))
	l = luckyNumPool[idx]
	return
}

// 抽奖
func LotteryDraw(c *gin.Context) error {
	openid, err := GetUserFromSession(c)
	if err != nil {
		return err
	}
	user := &User{OpenID: openid}
	err = DB.Model(user).First(user).Error
	if err != nil {
		return err
	}
	if user.Master != 1 {
		return errcode.ErrorAuthority
	}
	// 获取下次开奖类型
	nextReward := &Reward{}
	if err = nextReward.GetNextReward(); err != nil {
		return err
	}
	// 抽取号码
	luckyNum, err := GetLuckyNumber()
	if err != nil {
		return err
	}
	tx := DB.Begin()
	err = tx.Model(&nextReward).Update("used", 1).Error
	if err != nil {
		tx.Rollback()
		return errcode.ErrorServerErr
	}
	createTime := time.Now()
	record := &LotteryRecords{
		OpenID:     luckyNum.OpenID,
		RewardID:   nextReward.ID,
		LuckyNum:   luckyNum.LuckyNum,
		CreateTime: &createTime,
	}
	err = tx.Create(record).Error
	if err != nil {
		tx.Rollback()
		return errcode.ErrorServerErr
	}
	tx.Commit()
	DB.Model(record).Related(&record.Reward)
	DB.Model(record).Related(&record.User, "OpenID")
	return errcode.ErrorOK.WithData(record)
}
