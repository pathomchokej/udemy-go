package main

import (
	"fmt"
	"strings"
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
// type User struct {
// 	ID             int
// 	Name, Location string
// }

// type Player struct {
// 	User
// 	GameID int
// }

// type Player2 struct {
// 	*User
// 	GameID int
// }

// func (u *User) Greeting() string {
// 	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
// }

// type Job struct {
// 	Command string
// 	Logger  *log.Logger
// }

// type Job2 struct {
// 	Command string
// 	*log.Logger
// }

// func main() {
// 	player := Player{
// 		User:   User{ID: 0, Name: "Matt", Location: "LA"},
// 		GameID: 10,
// 	}
// 	fmt.Printf("Player : %+v\n", player)
// 	fmt.Printf("ID %d, Name: %s, Location: %s, GameID: %d\n", player.ID, player.Name, player.Location, player.GameID)
// 	fmt.Println(player.Greeting())

// 	player2 := Player2{
// 		&User{ID: 0, Name: "John", Location: "USA"},
// 		219,
// 	}
// 	fmt.Printf("Player2 : %+v\n", player2)
// 	fmt.Printf("ID %d, Name: %s, Location: %s, GameID: %d\n", player2.ID, player2.Name, player2.Location, player2.GameID)
// 	fmt.Println(player2.Greeting())

// 	job := Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
// 	fmt.Printf("Job : %+v\n", job)
// 	job.Logger.Printf("test")

// 	job2 := Job2{}
// 	fmt.Printf("Job2 : %+v\n", job2)
// 	job2.Logger = log.New(os.Stderr, "Job: ", log.Ldate)
// 	job2.Command = "demo2"
// 	fmt.Printf("Job2 : %+v\n", job2)
// 	job2.Printf("test2")
// }

//////////////////////////////////
// #6 Collection Types

// func main() {
// 	var a [2]string
// 	a[0] = "hello"
// 	a[1] = "world"

// 	fmt.Println(a[0], a[1])
// 	fmt.Println(a)

// 	b := [...]string{"hello", "world2", "3", "4", "5", "6", "7"}
// 	fmt.Println(b)
// 	fmt.Printf("%q \n", b)
// 	fmt.Println(b[1:4])

// 	var c [2][3]string
// 	for col := 0; col < 2; col++ {
// 		for row := 0; row < 3; row++ {
// 			c[col][row] = fmt.Sprintf("C%d,R%d", col, row)
// 		}
// 	}
// 	fmt.Printf("%q \n", c)

// 	d := [][]string{{"1", "2", "3"}, {"4", "5", "6", "7"}}
// 	fmt.Printf("%q \n", d)

// 	arr1 := []string{"1", "2", "3", "4", "5"}
// 	arr2 := []string{"4", "5", "6", "7"}
// 	fmt.Printf("%q \n", arr1)
// 	arr3 := append(arr1, arr2...)
// 	fmt.Printf("%q \n", arr3)
// }

////////////////////////////////
// #7 range and map
// func main() {
// 	arr1 := []string{"1", "2", "3", "4", "5"}
// 	for index, value := range arr1 {
// 		fmt.Printf("[%d] value is %s \n", index, value)
// 	}
// 	fmt.Println("================================")

// 	for index := range arr1 {
// 		fmt.Printf("[%d] value is %s \n", index, arr1[index])
// 	}
// 	fmt.Println("================================")

// 	for _, value := range arr1 {
// 		fmt.Printf("value is %s \n", value)
// 	}
// 	fmt.Println("================================")

// 	phoneNumber := map[string]int{
// 		"KFC":       1112,
// 		"PIZZA HUT": 1150,
// 		"HOME":      0,
// 	}

// 	for key, value := range phoneNumber {
// 		fmt.Printf("[%s]value is %d \n", key, value)
// 	}
// 	fmt.Println("================================")

// 	for key := range phoneNumber {
// 		fmt.Printf("[%s]value is %d \n", key, phoneNumber[key])
// 	}
// 	fmt.Println("================================")

// 	for _, value := range phoneNumber {
// 		fmt.Printf("value is %d \n", value)
// 	}
// 	fmt.Println("================================")

// }

// var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
// 	"Emma", "Isabella", "Emily", "Madison",
// 	"Ava", "Olivia", "Sophia", "Abigail",
// 	"Elizabeth", "Chloe", "Samantha",
// 	"Addison", "Natalie", "Mia", "Alexis"}

// var test = map[string]int{
// 	"a": 1,
// 	//"a": 2,
// 	"b": 3,
// }

// func main() {

// 	var result = map[int][]string{}

// 	var maxLength int = 0
// 	for _, name := range names {
// 		length := len(name)
// 		result[length] = append(result[length], name)
// 		if maxLength < length {
// 			maxLength = length
// 		}
// 	}
// 	fmt.Printf("%#v\n\n", result)
// 	// %v is map[3:[Ava Mia] 4:[Evan Neil Adam Matt Emma] 5:[Emily Chloe] 6:[Martin Olivia Sophia Alexis] 7:[Katrina Madison Abigail Addison Natalie] 8:[Isabella Samantha] 9:[Elizabeth]]
// 	// %#v is map[int][]string{3:[]string{"Ava", "Mia"}, 4:[]string{"Evan", "Neil", "Adam", "Matt", "Emma"}, 5:[]string{"Emily", "Chloe"}, 6:[]string{"Martin", "Olivia", "Sophia", "Alexis"}, 7:[]string{"Katrina", "Madison", "Abigail", "Addison", "Natalie"}, 8:[]string{"Isabella", "Samantha"}, 9:[]string{"Elizabeth"}}

// 	var result2 = make([][]string, maxLength)
// 	for nameLength, categories := range result {
// 		result2[nameLength-1] = append(result2[nameLength-1], categories...)
// 	}
// 	fmt.Printf("%#v\n\n", result2)
// }

// var lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

// func WordCount(s string) map[string]int {
// 	words := strings.Fields(s)
// 	result := make(map[string]int)
// 	for _, word := range words {
// 		result[word]++
// 	}

// 	return result
// }

// func main() {
// 	result := WordCount(lorem)
// 	fmt.Printf("%#v", result)
// }

//////////////////////////////////
// #8 switch cases

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {

	for _, user := range users {

		length := len(user)
		lowercaseUser := strings.ToLower(user)
		for i := 0; i < length; i++ {
			// number of coin can eurn
			var coin = 0
			switch lowercaseUser[i] {
			case 'a', 'e':
				coin = 1
			case 'i':
				coin = 2
			case 'o':
				coin = 3
			case 'u':
				coin = 4
			}

			// validate max coin can earn
			maxEarn := 10 - distribution[user]
			if coin > maxEarn {
				coin = maxEarn
			}

			// validate exist coins
			if coins < coin {
				coin = coins
			}

			distribution[user] += coin
			coins -= coin

			if distribution[user] >= 10 || coins <= 0 {
				break
			}
		}

		if coins <= 0 {
			break
		}
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
