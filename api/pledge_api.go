package main

import (
	"fmt"
	"pledge-backend/db"
)

func main() {
	fmt.Println("init")
	db.InitMysql()
}
