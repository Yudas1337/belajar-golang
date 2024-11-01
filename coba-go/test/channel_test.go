package test

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Yudas Malabi"
		fmt.Println("Berhasil kirim data ke channel")
	}()

	time.Sleep(5 * time.Second)

	data := <-channel
	fmt.Println(data)

	defer close(channel)
}
