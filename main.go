package main

import (
	"fmt"
	"os"
)

func main() {
	//init
	fmt.Println("ENVIROMENT")
	fmt.Println("ENV: ", os.Getenv("ENV"))
	fmt.Println("TASK_TYPE: ", os.Getenv("TASK_TYPE"))
	fmt.Println("API_VERSION: ", os.Getenv("API_VERSION"))
	fmt.Println("CONFIG_FILE: ", os.Getenv("CONFIG_FILE"))
	fmt.Println("BINARY_NAME: ", os.Getenv("BINARY_NAME"))
	fmt.Println("PORT: ", os.Getenv("PORT"))
	fmt.Println("AWS_REGION: ", os.Getenv("AWS_REGION"))
	fmt.Println("AWS_ACCOUNT_ID: ", os.Getenv("AWS_ACCOUNT_ID"))
	fmt.Println("AWS_PROFILE: ", os.Getenv("AWS_PROFILE"))
	fmt.Println("DESIRED_COUNT: ", os.Getenv("DESIRED_COUNT"))
	fmt.Println("")
}
