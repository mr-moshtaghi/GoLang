package main

import "fmt"

type Vertex struct {
	Latitude, Longitude float64
}

func main() {
	cities := map[string]Vertex{}
	cities["tehran"] = Vertex{53.80, 47.7}
	cities["isfahan"] = Vertex{40.60, 30.5}
	fmt.Println(cities)
}
