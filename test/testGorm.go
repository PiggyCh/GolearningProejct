package main

import (
	"fmt"
	Models "ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123123123@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	fmt.Println("Database connection successful!")

	// // 迁移 schema
	db.AutoMigrate(&Models.UserBasic{})

	// // Create
	user := &Models.UserBasic{}
	user.Name = "jinzhu"
	user.Password = "123123"
	db.Create(user)

	// // Read
	fmt.Println("Read")
	fmt.Println(db.First(user, 1)) // find product with integer primary key

	// // Update - 将 product 的 price 更新为 200
	// db.Model(&product).Update("Price", 200)
	// // Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
}
