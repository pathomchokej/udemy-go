package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////////
// #1 Basic
// func main() {
// 	//init
// 	fmt.Println("ENVIROMENT")
// 	fmt.Println("ENV: ", os.Getenv("ENV"))
// 	fmt.Println("TASK_TYPE: ", os.Getenv("TASK_TYPE"))
// 	fmt.Println("API_VERSION: ", os.Getenv("API_VERSION"))
// 	fmt.Println("CONFIG_FILE: ", os.Getenv("CONFIG_FILE"))
// 	fmt.Println("BINARY_NAME: ", os.Getenv("BINARY_NAME"))
// 	fmt.Println("PORT: ", os.Getenv("PORT"))
// 	fmt.Println("AWS_REGION: ", os.Getenv("AWS_REGION"))
// 	fmt.Println("AWS_ACCOUNT_ID: ", os.Getenv("AWS_ACCOUNT_ID"))
// 	fmt.Println("AWS_PROFILE: ", os.Getenv("AWS_PROFILE"))
// 	fmt.Println("DESIRED_COUNT: ", os.Getenv("DESIRED_COUNT"))
// 	fmt.Println("")

// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Printf("%d\n", i)
// 		}
// 	}
// }

////////////////////////////////////////////////////////////////
// #2 function parameter
// type Artist struct {
// 	Name  string
// 	Songs int
// }

// func NewRelease(a *Artist) int {
// 	a.Songs++
// 	return a.Songs
// }

// // test parameter with interface
// func timeStamp(data interface{}) {
// 	mapData, ok := data.(map[string]interface{})
// 	if ok {
// 		mapData["Matt"] = 33
// 		mapData["update_at"] = time.Now()
// 	}
// }

// func main() {
// 	me := &Artist{Name: "John", Songs: 10}
// 	fmt.Printf("%s new release = %d\n", me.Name, NewRelease(me))
// 	fmt.Printf("%s current release = %d\n", me.Name, me.Songs)

// 	data1 := map[string]interface{}{
// 		"Matt": 42,
// 	}

// 	fmt.Println(data1)
// 	timeStamp(data1)
// 	fmt.Println(data1)
// }

////////////////////////////////////////////////////////////////////////
// #3 test implement inferface
type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

// implement Stringer interface to fakeString
func (s *fakeString) String() string {
	return s.content
}

func PrintString(value interface{}) {
	fmt.Printf("\n## type is %T = %v\n", value, value)
	switch valueType := value.(type) { // compare type instead of value
	case string:
		fmt.Printf("- type %T = %v\n", valueType, valueType)
	case Stringer:
		fmt.Printf("- type %T = %v, %v\n", valueType, valueType, valueType.String())
	default:
		fmt.Printf("- Not support type %T = %v\n", value, value)
	}
}

func main() {
	s := fakeString{"string in struct fakeString"}
	PrintString(s)
	PrintString(&s)
	PrintString("string directly")
}
