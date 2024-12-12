package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255); not null"` //用户名
	Password  string `gorm:"type:varchar(255); not null"` //密码
	IsAdmin   int    `gorm:"type:tinyint(1); not null"`   //权限 0:普通用户 1:管理员 2:超级管理员
	Email     string `gorm:"type:varchar(255)"`           //邮箱
	AcceptNum int64  `gorm:"type:int(11); default:0"`     //通过次数
	Submit    int64  `gorm:"type:int(11); default:0"`     //提交次数
}

func (User) TableName() string {
	return "user"
}
