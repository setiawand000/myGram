package main

import (
	"mygram/database"
	"mygram/routers"
)

func main() {
	database.DbStart()

	// fmt.Println(os.Getenv("HOST_DB"))
	// fmt.Println(os.Getenv("USER_DB"))
	// fmt.Println(os.Getenv("PASS_DB"))
	// fmt.Println(os.Getenv("PORT_DB"))

	r := routers.StartApp()
	r.Run(":8080")
}
