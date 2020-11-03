package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Tags struct {
	Id         int    `gorm:"primary_key;column:id" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	Mid        int    `gorm:"column:mid" json:"mid"`
	CategoryId int    `gorm:"column:category_id" json:"categoryId"`
}

func (Tags) TableName() string {
	return "tags"
}


func GetAll(tags *[]Tags) {
	db, err := gorm.Open("mysql", "root:123456@/dev_jingsocial?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		println(err)
	}

	defer db.Close()
	err = db.Where("mid = ? and is_default = ? and `level` = ?", 33, 6, 1).Order("id desc").Select("id,name,mid,category_id").Find(tags).Error
	if err != nil {
		fmt.Println(err)
	}
}
