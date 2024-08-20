package Models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	fmt.Println("----------------------------------")
	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println("----------------------------------")
	return data
}

func CreateUser(user UserBasic) *gorm.DB {
	var defaultTime = time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC)
	// // Create
	user.LoginTime = defaultTime
	user.LogoutTime = defaultTime
	user.HeartbeatTime = defaultTime
	return utils.DB.Create(&user)
}

