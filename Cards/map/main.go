package main

import "fmt"

func main(){

	// colors := map[string]string{
	// 	"Red":"#FF0",
	// 	"Black": "#ccc",
	// 	"Green": "#ddd",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	colors["White"] = "#FFF"
	colors["Black"] = "#FFF"
	colors["Blue"] = "#FFF"
	colors["Red"] = "#FFB"
	delete(colors, "Red")


	printMap(colors)
	fmt.Println(colors)
}

func printMap(m map[string]string){
	for key, value := range m{
		fmt.Println(key,value)
	}
}