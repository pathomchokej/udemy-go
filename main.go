package main

import (
	"fmt"
)

type Artist struct {
	Name  string
	Songs int
}

func NewRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	//init
	// fmt.Println("ENVIROMENT")
	// fmt.Println("ENV: ", os.Getenv("ENV"))
	// fmt.Println("TASK_TYPE: ", os.Getenv("TASK_TYPE"))
	// fmt.Println("API_VERSION: ", os.Getenv("API_VERSION"))
	// fmt.Println("CONFIG_FILE: ", os.Getenv("CONFIG_FILE"))
	// fmt.Println("BINARY_NAME: ", os.Getenv("BINARY_NAME"))
	// fmt.Println("PORT: ", os.Getenv("PORT"))
	// fmt.Println("AWS_REGION: ", os.Getenv("AWS_REGION"))
	// fmt.Println("AWS_ACCOUNT_ID: ", os.Getenv("AWS_ACCOUNT_ID"))
	// fmt.Println("AWS_PROFILE: ", os.Getenv("AWS_PROFILE"))
	// fmt.Println("DESIRED_COUNT: ", os.Getenv("DESIRED_COUNT"))
	// fmt.Println("")

	// //fmt.Println(hello.Hello())

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%d\n", i)
	// 	}
	// }

	me := &Artist{Name: "John", Songs: 10}
	fmt.Printf("%s new release = %d\n", me.Name, NewRelease(me))
	fmt.Printf("%s current release = %d\n", me.Name, me.Songs)

}
