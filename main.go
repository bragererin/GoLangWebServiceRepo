package main

import (
	controllers "Itential/AverageComputation/GoLangWebService/controller"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
