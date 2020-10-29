package model

import (
	"fmt"
	"myblog/utils"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}

	defer db.Close()
	//db.Close()
}
