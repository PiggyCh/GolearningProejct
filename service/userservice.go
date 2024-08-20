package service

import (
	Models "ginchat/models"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags         首页
// @Success      200  {string}  json{"code", "message"}
// @Router       /user/GetUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*Models.UserBasic, 10)
	data = Models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})

}

// CreateUser
// @Tags         用户模块
// @Summary      创建用户
// @param        name query string true "用户名"
// @param        password query string true "密码"
// @param        repassword query string true "确认密码"
// @Success      200  {string}  json{"code", "message"}
// @Router       /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := Models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	// 参数验证
	if name == "" || password == "" || repassword == "" {
		c.JSON(400, gin.H{
			"message": "参数不能为空",
		})
		return
	}

	if password != repassword {
		c.JSON(400, gin.H{
			"message": "两次密码不一致",
		})
		return
	}

	user.Name = name
	user.Password = password
	// 这里修正了对错误的处理方式
	if db := Models.CreateUser(user); db.Error != nil {
		c.JSON(500, gin.H{
			"message": "创建用户失败",
			"error":   db.Error.Error(), // 提供更详细的错误信息
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "创建成功",
	})
}
