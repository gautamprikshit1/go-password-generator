package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Let's generate a password for you!")
	var chars string = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	fmt.Print("Please enter the desired length of password: ")
	var length int
	var pass string
	_, err := fmt.Scanf("%d", &length)
	rand.Seed(time.Now().UnixNano())
	if err != nil {
		log.Fatal(fmt.Println(err))
	} else {
		if length < 8 {
			log.Fatal(fmt.Println("Password length should be greater than 8."))
		} else {
			i := 0
			for i < length {
				var rand_num = rand.Intn(len(chars))
				pass += string(chars[rand_num])
				i++
			}
		}
	}
	fmt.Println("Your required password is " + pass)
}
