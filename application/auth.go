package application

import (
	"fmt"

	Common "../common"
	DB "../common/db"
	Logger "../common/logger"
	"github.com/gofiber/fiber"
)

//注册
func RegisterHandle() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		mobile := c.Body("mobile")
		password := c.Body("password")
		Logger.Infof(mobile)
		if mobile == "" || password == "" {
			c.JSON("params is not null")
			return
		}
		//查询是否注册过
		userLogin := new(Common.UserLogin)
		row := DB.MysqlDb.QueryRow("select id, mobile from user_login where mobile=?", mobile)
		err := row.Scan(&userLogin.Id, &userLogin.Mobile)
		if err == nil {
			c.JSON("mobile is exist")
			return
		}
		//添加记录
		passwordMd5Str := Common.Md5V(password)
		ret, _ := DB.MysqlDb.Prepare("insert into user_login(mobile, password) value (?, ?)")
		addRet, _ := ret.Exec(mobile, passwordMd5Str)
		ins_nums, _ := addRet.RowsAffected()
		if ins_nums > 0 {
			c.JSON(200)
			return
		} else {
			c.JSON(201)
			return
		}
	}
}

//登陆
func LoginHandle() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		result := 0
		for i := 1; i <= 10; i++ {
			result = fibonacci(i)
			fmt.Printf("fibonacci(%d) is: %d\n", i, result)
		}
	}
}

func fibonacci(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
