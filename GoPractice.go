/* Dependency Management 

go mod init -  creates mod file

go.mod:

module cfg
require "package location - github repo url" v1.2.0

go.sum - packages you have aquired

*/

/* GO CLI

go build - builds into exe
./goFileExe

time go run - times the execution of a go program

go get - package manager
fmt - format code

*/

/* Types
Go has very strict type system - little implicit type conversion

Types:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

/*
Type Conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
*/

/*
Programs start with "package main" - this will turn into .exe file and not a package 
*/
package main

import (
	"fmt"
)

// 'Var' Declares a list of variables - Package level variables. Type is the last thing in the list
/* var c, python, java bool
* var b, boolOne, strOne = true, false, "no!"
*/

func add(x int, y int) int {
	return x + y
}

// Multiple results
func swap(x, y string) (string, string) {
	
	// Short declarations - not used outside of a function.
	// := vs = - := declares a NEW variable
	/*
	k := 3
	c, python, java := true, false, "no!"
	*/

	return y, x
}



func loops(x int) {
	sum := 0
	for i := 0; i < x; i++ {
		sum += i
	}
	fmt.Println(sum)

	if sum > 100 {
		fmt.Printf("Beeeg value %v\n", sum)
	}

	a := 3
	for a < 3{
		fmt.Printf("runnin\n")
	}

	for {
		// like while true loop - must break yourself
		fmt.Printf("runnin\n")
		break;
	}

	loons := []string {"birdOne", "birdTwo", "birdThree"}
	for i, name := range loons {
		fmt.Printf("Index: %v, Bird: %v\n", i, name)
	} 

	// _ is a wildcard for "un used variable"
	for _, name := range loons {
		fmt.Printf("Index: %v, Bird: %v\n", name)
	} 
}

func ifs() {
	a := 1
	b := 2
	if frac := a/b; frac > 1 {
		fmt.Printf("true\n")
	}
}

func strings() {
	book := "test strings"
	fmt.Printf("first byte of string: %v\n", book[0]) // - access first byte of string (byte = unit8)

	// access characters by slicing
	// book[:4] // 0-4 charactrs

	// STRINGS ARE UNICODE - can take any character
}

func maps() {
	stocks := map[string]float64 {
		"amzn":1699.9,
		"goog": 1124.3, // Must have trailing comma
	}

	fmt.Printf("value in map: %v", stocks["amzn"]) // prints the zero value if not there
	//	value, ok := stocks["amzn"] // prints the true or false

	// delete(stocks, "amzn") // delete key value pair

	// for key := range stocks {

	// }

	// for key, value := range stocks {

	// }

}

/* GO - pass by value if primitive
  maps, slices, objects are pass by reference
*/
func pointers(n *int) {

	// Use a pointer to pass by reference 
	*n *= 2 // dereference pointer to keep the value when returned
}

func errors()(int, error) {
	s1, err := sqrt(-2.0)
	if err != nil {
		// ERROR HANDLING FOR ALL FUNCTIONS
		return 0.0, fmt.Errorf("string here") // - creats error type 
	} else {
		fmt.Printf(s1)
	}
	return s1
}

func defer_ex() {
	defer fmt.Print("defer") // always called after code is run - used as a stack with multiple defers 

	fmt.Printf("Do work")
}

func networking() {
	resp, err := http.Get("https://api.guthub.com")
	if err != nil {
		// error happened
	}
	resp.Header.Get("Content-Length")
	resp.body.Close()
}

type Trade struct {
	Symbol string // Uppercase variables are accessiabe to other packages
	Volume int
	Price float64
	Buy bool
}

func ObjectOriented() {
	t1 := Trade{"MSFT", 1000, 1393.33, false}
	fmt.Printf(t1.Symbol)
	fmt.Printf("Object: %+v\n", t1)
}

/* first part of method signature is a RECIEVER
Reciever adds method to the struct ex) t1.GetValue()
*/
func (t1 *Trade) GetValue() float64 {
	return float64(t1.Price) * t1.Volume
}

/* Constructors - should start with "New" and returnd a pointer trade*/
func NewTrade(sym string) (*Trade, error) {
	trade := &Trade{
		Symbol: sym,
		Volume: 100
		Price: 100.00
		Buy: false
	}

	return trade, nil
}

/* Interfaces */
type Shape interface {
	Area(num int) (n float64, err error)
}


/* Error Handeling - 
   go get github/pkg/erros - error handleing with exception 

   ex)
   if err != nill {
		return nil errors.Wrap(err, "can't do task") - Wrap() will return a nice stack trace
	}
*/


/* Concurrency 
	go routines
	WaitGroup() - from the sync package - waits for all go routines to finsihed
*/
func ConcurrencyExamples() {

	urls := []string { "urlOne", "urlTwo", "urlThree"}

	var wg sync.WaitGroup
	for index, url := range urls {
		wg.Add(1)
		go func(url string) {
			http.Get(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

/* Channels - Used to manage go routines
   set value aue at one end, recieve at the other end. Blocked when there is nothing on recieving end */
func channels() {
	ch := make(chan int) // make a channel of ints

	go func() {
		// Send number of the channel
		ch <- 353
	}()

	val := <-ch
	fmt.Printf("Got: %v", val) 

	/* Channels will store values until their are retrieved */
}




func main() {
	fmt.Println(add(42, 13))
	fmt.Println("Hello, World!")
	a, b:= swap("yoyo", "ma")
	fmt.Println(a, b)
	
	loops(50)

	var x int
	x = 1
	y := 2
	fmt.Printf("x=%v, type of %T\n", x, y)

}

