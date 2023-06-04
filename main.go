package main

import (
	db "github.com/YuriBertoldi/Go-APIRest/conexaodb"
	routes "github.com/YuriBertoldi/Go-APIRest/routes"
)

func main() {
	db.Connect()
	routes.HandleRequest()
}
