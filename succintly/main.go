package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
)

const (
	myName string = `
	Giancarlos
	Salas`
	old int = 24
)

func main() {
	workServer()
}

func workServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}

func broadcast2(nsChannel chan int, cChannel chan bool) {
	numbers := []int{
		101,
		102,
		103,
		104,
		105,
		106,
		107,
		108,
		109,
		110,
	}
	i := 0

	for {
		// see which channel has items
		select {
		/* if the numbersChannel is being listened to,
		take each number sequentially from the
		slice and put it into the channel */
		case nsChannel <- numbers[i]:
			i++
			/* if we've reached the last number and
			the channel is still being listened to,
			start reading from the beginning of the
			slice again */
			if i == len(numbers) {
				i = 0
			}
		/* if we receive a message on the
		complete channel, we stop transmitting */
		case <-cChannel:
			return
		}
	}
}

func workChannels4() {
	log.Println("Init")
	numbersStation := make(chan int)
	completeChannel := make(chan bool)
	// execute broadcast in a separate thread
	go broadcast2(numbersStation, completeChannel)
	// get 100 numbers from the numbersStation channel
	for i := 0; i < 100; i++ {
		// delay for artistic effect only
		time.Sleep(100 * time.Millisecond)
		// retrieve values from the channel
		fmt.Printf("%d ", <-numbersStation)
	}
	/* once we have received 100 numbers,
	   send a message on completeChannel
	   to tell it to stop transmitting */
	completeChannel <- true
	fmt.Println("Transmission Complete.")
}

func workChannels3() {
	c := make(chan int, 1)
	c <- 2
	fmt.Println(<-c)
	c <- 2
	fmt.Println("OK.")
}

func generateAccountNumber(accountNumberChannel chan int) {
	// internal variable to store last generated account number
	var accountNumber int
	accountNumber = 30000001
	for {
		// close the channel after 5 numbers are requested
		if accountNumber > 30000005 {
			close(accountNumberChannel)
			return
		}
		accountNumberChannel <- accountNumber
		// increment the account number by 1
		accountNumber++
	}
}

func workChannels2() {
	accountNumberChannel := make(chan int)
	go ad(accountNumberChannel)
	fmt.Println(<-accountNumberChannel)
	fmt.Println(<-accountNumberChannel)
	fmt.Println(<-accountNumberChannel)
	// start the goroutine that generates account numbers
	go generateAccountNumber(accountNumberChannel)
	// slice containing new customer names
	newCustomers := []string{"SMITH", "SINGH", "JONES", "LOPEZ",
		"CLARK", "ALLEN"}
	// get a new account number for each customer
	for _, newCustomer := range newCustomers {
		// is there anything to retrieve from the channel?
		accnum, ok := <-accountNumberChannel
		if !ok {
			fmt.Printf("%s: No number available\n",
				newCustomer)
		} else {
			fmt.Printf("%s: %d\n", newCustomer, accnum)
		}
	}

}

func workChannels() {
	numbersStation := make(chan int)
	// execute broadcast in a separate thread
	go broadcast(numbersStation)
	// retrieve values from the channel
	for num := range numbersStation {
		// delay for artistic effect only
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%d ", num)
	}
}

func ad(c chan int) {
	c <- 2
	close(c)
}

func broadcast(c chan int) {
	// infinite loop to create random numbers
	for {
		/* generate a random number 0-999
		and put it into the channel */
		c <- rand.Intn(999)
	}
}

func workRoutines() {
	go message("Gian")
	message("Hola")
}

func message(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

type rect struct {
	height, width int
}

type square struct {
	rect
}

func (r rect) area() (area int) {
	area = r.width * r.height
	return
}

func (s square) area() int {
	return s.height * s.height
}

type logger interface {
	write(value string)
}

type consoleLogger struct{}

func (s consoleLogger) write(value string) {
	fmt.Printf("Text: %s\n", value)
}

func workInterfaces() {
	console := &consoleLogger{}
	items := []string{
		"Item1",
		"Item2",
		"Item3",
	}

	writeLog(console, "Hola", "Logger")
	writeLog(console, items...)

	var anything interface{} = "something"
	aInt, ok := anything.(int)
	if !ok {
		fmt.Println("Cannot turn input into an integer")
	} else {
		fmt.Println(aInt)
	}

}

func writeLog(log logger, a ...string) {
	for _, value := range a {
		log.write(value)
	}
}

func workStructs() {
	r := square{}
	r.height = 5
	myRectangle := &rect{
		height: 12,
		width:  20,
	}
	myRectangle.width = 5
	r2 := new(rect)
	// var pp *int = new(int)
	pp := new(int)
	*pp = 4

	fmt.Println(r)
	fmt.Println(r2)
	fmt.Println(*pp)
	fmt.Println(*myRectangle)
	fmt.Println("Area", r.area())
}

func workPointers() {
	var p *int
	value := 4
	p = &value
	*p = 6

	fmt.Println(*p)
}

func workMaps() {
	actor := map[string]int{
		"Paltrow": 43,
		"Cruise":  53,
		"Redford": 79,
		"Diaz":    43,
		"Kilmer":  56,
		"Pacino":  75,
	}
	props := make(map[string]int, 2)
	props["a1"] = 3
	props["a2"] = 4
	actor["add"] = 34
	fmt.Println(actor)
	fmt.Println(props)

	age2 := actor["abc"]
	fmt.Println("Default if not found: ", age2)
	if age, ok := actor["Pacino"]; ok {
		fmt.Println("Age: ", age)
	}

	for i := 1; i < 4; i++ {
		fmt.Printf("\nRUN NUMBER %d\n", i)
		for key, value := range actor {
			fmt.Printf("%s : %d years old\n", key, value)
		}
	}
}

func workSlices() {
	mySlice2 := make([]int, 4)
	fmt.Printf("Capacity 1: %d\n", cap(mySlice2))
	fmt.Printf("Length 1: %d\n", len(mySlice2))
	mySlice2 = append(mySlice2, 19)
	fmt.Printf("Length 1: %d\n", len(mySlice2))
	fmt.Printf("Capacity 1: %d\n", cap(mySlice2))

	mySlice := make([]int, 0, 8)
	fmt.Printf("Contents: %v\n", mySlice)
	fmt.Printf("Number of Items: %d\n", len(mySlice))
	fmt.Printf("Capacity: %d\n", cap(mySlice))

	mySlice = append(mySlice, 1, 3, 5, 7, 9, 11, 13, 17)
	fmt.Printf("Contents: %v\n", mySlice)
	fmt.Printf("Number of Items: %d\n", len(mySlice))
	fmt.Printf("Capacity: %d\n", cap(mySlice))
	mySlice = append(mySlice, 19)

	fmt.Printf("Contents: %v\n", mySlice)
	fmt.Printf("Number of Items: %d\n", len(mySlice))
	fmt.Printf("Capacity: %d\n", cap(mySlice))

	mySlice3 := make([]int, 3)
	copy(mySlice3, mySlice)
	fmt.Printf("Contents: %v\n", mySlice3)
	fmt.Printf("Number of Items: %d\n", len(mySlice3))
	fmt.Printf("Capacity: %d\n", cap(mySlice3))
}

func applyArr(items [5]int) {
	items[0] = 2
}

func printArr() {
	totals := [...]int{1, 2, 3, 4, 5}
	applyArr(totals)
	var arr = [5][2]int{{0, 0}, {2, 4}, {1, 3}, {5, 7}, {6, 8}}
	fmt.Println(totals[:3])
	fmt.Println(arr)
}

func printArgs() {
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	for index := 0; index < 2; index++ {
		fmt.Println(index)
	}
}

func showOldStatus(edad string) {
	old, _ := strconv.Atoi(edad)
	switch {
	case old > 0 && old < 10:
		fmt.Println("Child")
	case old < 20:
		fmt.Println("Adolescente")
	case old < 30:
		fmt.Println("Young")
	case old < 60:
		fmt.Println("Señor")
	default:
		fmt.Println("Anciano")
	}
}

func showType(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Println("Boolean")
	case int:
		fmt.Println("Entero")
	case string:
		fmt.Println("Cadena")
	}
}

func hello() {
	fmt.Println("Hello", myName)
	fmt.Println("Type", reflect.TypeOf(old))
}

func sum() {
	a, b := os.Args[1], os.Args[2]
	n1, _ := strconv.Atoi(a)
	n2, _ := strconv.Atoi(b)

	c := n1 + n2

	fmt.Println("Result", a, "+", b, "=", c)
}
