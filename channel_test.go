package golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// membuat channel
	channel := make(chan string)

	// mengirim data ke channel
	//channel <- "hello"

	// menerima data dari channel
	//data := <-channel

	//fmt.Println(<-channel)
	//fmt.Println(data)

	// menutup channel
	//close(channel)
	//defer close(channel)

	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		channel <- "world"
		fmt.Println("selesai kirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(time.Duration(5) * time.Second)
	defer close(channel)
}

// belajar menggunakan channel sebagai parameter

func GiveMeResponse(channel chan string) {
	time.Sleep(time.Duration(2) * time.Second)
	channel <- "hello"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(time.Duration(5) * time.Second)
}

// belajar channel in dan out

func OnlyIn(channel chan<- string) {
	time.Sleep(time.Duration(2) * time.Second)
	channel <- "hello"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(time.Duration(5) * time.Second)
}

// belajar buffered channel

func TestBufferedChannel(t *testing.T) {
	// salah satu solusi untuk error deadlock
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "hello"
	channel <- "world"
	channel <- "bye"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("selesai")
}

// belajar range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke-" + strconv.Itoa(i)
		}
		// setelah mengirim data close channel nya disini
		close(channel)
		// jika tidak diclose perulangan akan terus berlanjut tapi data sudah tidak ada
		// muncul error deadlock
	}()

	for data := range channel {
		fmt.Println("Menerima data: ", data)
	}

	fmt.Println("selesai kirim data ke channel")
}

// belajar select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel1 : ", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel2 : ", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// belajar default channel
func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel1 : ", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel2 : ", data)
			counter++
		default:
			fmt.Println(" no data")
		}
		if counter == 2 {
			break
		}
	}
}

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x++
			}
		}()
	}

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println(x)
}
