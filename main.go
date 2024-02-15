package main

import (
	"munenendereba/sms-backend/db"
	"munenendereba/sms-backend/routers"
)

func main() {
	db.InitializeDatabase()
	routers.InitializeRouter().Run()
}
