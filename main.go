package main

import (
	"fmt"
	"log"
	"os"
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
// type Stringer interface {
// 	String() string
// }

// type fakeString struct {
// 	content string
// }

// // implement Stringer interface to fakeString
// func (s *fakeString) String() string {
// 	return s.content
// }

// func PrintString(value interface{}) {
// 	fmt.Printf("\n## type is %T = %v\n", value, value)
// 	//switch value.(type) { // compare type instead of value
// 	switch valueType := value.(type) { // compare type instead of value
// 	case string:
// 		fmt.Printf("- type %T = %v\n", valueType, valueType)
// 		//fmt.Printf("- type %T = %v\n", value, value.(string))
// 	case Stringer:
// 		fmt.Printf("- type %T = %v, %v\n", valueType, valueType, valueType.String())
// 		//fmt.Printf("- type %T = %v, %v\n", value, value, value.(Stringer).String())
// 	default:
// 		fmt.Printf("- Not support type %T = %v\n", value, value)
// 	}
// }

// func main() {
// 	s := fakeString{"string in struct fakeString"}
// 	PrintString(s)
// 	PrintString(&s)
// 	PrintString("string directly")
// }

////////////////////////////////////////////////////////////////////////
// #4 struct fields
// type Point struct {
// 	X, Y int
// }

// var (
// 	p = Point{1, 2}
// 	q = &Point{1, 2}
// 	r = Point{Y: 1}
// 	s = Point{}
// )

// type BootCamp struct {
// 	Lat, Lon float64
// 	Date     time.Time
// }

// func main() {
// 	fmt.Println(p, q, r, s)

// 	event1 := &BootCamp{}
// 	event2 := new(BootCamp)
// 	fmt.Println(event1, " and ", event2, " : pointer check[event1 == event2 is ", event1 == event2, "] value check[*event1 == *event2 is ", *event1 == *event2, "]")
// }

////////////////////////////////////////////////////////////////////////
// #5 composition struct
type User struct {
	ID             int
	Name, Location string
}

type Player struct {
	User
	GameID int
}

type Player2 struct {
	*User
	GameID int
}

func (u *User) Greeting() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Job struct {
	Command string
	Logger  *log.Logger
}

type Job2 struct {
	Command string
	*log.Logger
}

func main() {
	player := Player{
		User:   User{ID: 0, Name: "Matt", Location: "LA"},
		GameID: 10,
	}
	fmt.Printf("Player : %+v\n", player)
	fmt.Printf("ID %d, Name: %s, Location: %s, GameID: %d\n", player.ID, player.Name, player.Location, player.GameID)
	fmt.Println(player.Greeting())

	player2 := Player2{
		&User{ID: 0, Name: "John", Location: "USA"},
		219,
	}
	fmt.Printf("Player2 : %+v\n", player2)
	fmt.Printf("ID %d, Name: %s, Location: %s, GameID: %d\n", player2.ID, player2.Name, player2.Location, player2.GameID)
	fmt.Println(player2.Greeting())

	job := Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	fmt.Printf("Job : %+v\n", job)
	job.Logger.Printf("test")

	job2 := Job2{}
	fmt.Printf("Job2 : %+v\n", job2)
	job2.Logger = log.New(os.Stderr, "Job: ", log.Ldate)
	job2.Command = "demo2"
	fmt.Printf("Job2 : %+v\n", job2)
	job2.Printf("test2")
}
