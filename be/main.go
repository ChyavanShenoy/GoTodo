package main

import "be/functions"

func main() {
	PORT := ":8080"
	functions.CheckCreateDB()
	StartServer(PORT)
}
