package database

import (
	"github.com/yamakiller/magicWeb/library/db/mysql"
	"github.com/yamakiller/magicWeb/library/models"
)

//AdminUserQuery Return page admin user and total
func AdminUserQuery(sqlHandle, account, order string, page, pageSize int) (users []models.AdminUser, total int, err error) {
	fileds := "id,account,password,nick,admin_profile.name,source,state,fail,fail_last_time,logined,Logined_last_time,logined_ip"
	if account != "" {
		err = mysql.Instance().
			DB(sqlHandle).
			Select(fileds).
			Where("account LIKE %?%", account).
			Limit(pageSize).
			Offset((page - 1) * pageSize).
			Order("CreatedAt " + order).Find(users).Error
	} else {
		err = mysql.Instance().
			DB(sqlHandle).
			Select(fileds).
			Limit(pageSize).
			Offset((page - 1) * pageSize).
			Order("CreatedAt " + order).Find(users).Error
	}
	if err != nil {
		return
	}

	mysql.Instance().DB(sqlHandle).Model(&models.AdminUser{}).Count(&total)
	return
}
