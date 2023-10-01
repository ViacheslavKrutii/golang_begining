package main

import mytransport "Proj/HomeWork/golang_begining/myInterfaces/transport"

func main() {

	mytransport.Travel("Ivan", []mytransport.City{"Kyiv", "London", "Paris"})
}
