package main

import (
	"munenendereba/sms-backend/db"
	"munenendereba/sms-backend/router"
)

func main() {
	db.InitializeDatabase()
	router.InitializeRouter().Run()
}
