# 环境
需要安装mysql
### 数据表定义
创建models包，并定义数据库内的字段
```golang
package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LoginTime     uint64
	LogoutTime    uint64
	HeartbeatTime uint64
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
```
### 安装gin和gorm
```go get -u github.com/gin-gonic/gin```  
```go get -u gorm.io/gorm```

### 测试gorm
```golang
package main

import (
	"fmt"
	"furrychat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(`root:root@tcp(127.0.0.1:3306)/
	ginchat?charset=utf8mb4&parseTime=True&loc=Local`),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "cathan"
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1)) // 根据整型主键查找

	// Update
	db.Model(user).Update("PassWord", "1234")
}

```
### 测试gin
```go
// main.go
package main

import "furrychat/myrouter"

func main() {
	r := myrouter.Router()
	r.Run(":8081")
}

// myrouter/app.go
package myrouter

import (
	"furrychat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	return r
}

// service/index.go
package service

import "github.com/gin-gonic/gin"

func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome ! ! ",
	})
}
```