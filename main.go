package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	zun  = "ズン"
	doko = "ドコ"
	kiyoshi = "キ・ヨ・シ！"
)

var zundoko = []string{zun, doko}

func main() {
	rand.Seed(time.Now().UnixNano())
	zndk := make(chan string)
	quit := make(chan bool)

	go func() {
		for {
			z1, z2, z3, z4, d := <-zndk, <-zndk, <-zndk, <-zndk, <-zndk
			fmt.Println(z1, z2, z3, z4, d)

			if isZundoko(z1, z2, z3, z4, d) {
				break
			}
		}
		quit <- true
	}()

	for {
		select {
		case <- quit:
			fmt.Println(kiyoshi)
			return
		default:
			go func() {
				zndk <- zundoko[rand.Int()%2]
			}()
		}
		time.Sleep(time.Nanosecond)
	}
}

func isZundoko(z1, z2, z3, z4, d string) bool {
	return z1 == zun && z2 == zun && z3 == zun && z4 == zun && d == doko
}
