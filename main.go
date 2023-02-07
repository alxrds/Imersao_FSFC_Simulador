package main

import (
	route2 "github.com/alxrds/imersaofsfc-simulador/application/route"
	"fmt"
)

func main(){
	route := route2.Route{
		ID: "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}