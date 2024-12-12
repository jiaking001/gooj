package dal

import (
	"gooj/dal/mysql"
	"gooj/dal/redis"
)

func init() {
	mysql.Init()
	redis.Init()
}
