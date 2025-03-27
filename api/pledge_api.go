package main

import (
	"pledge-backend/db"
)

func main() {
	//init mysql
	db.InitMysql()
	//init redis
	db.InitRedis()
}
