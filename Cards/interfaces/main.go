package main

import "fmt"

// type user struct{
// 	name string
// }
type bot interface{
	getGreeting() string
	//getGreeting(string, int) (string, error)
	//getBotVersion() float64
	//respondtoUser(user) string
}
type spanishBot struct{}
type englishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
