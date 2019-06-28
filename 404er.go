package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var alphanum = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func urlinput() {
	fmt.Println("What domain do you wish to 404?")
	var url string
	fmt.Scanln(&url)
	fmt.Println(url)
}

func randomizer(r int) string {
	b := make([]rune, r)
	for i := range b {
		b[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return string(b)
}

func main() {
	//create the seed
	rand.Seed(time.Now().UnixNano())

	fmt.Println("What domain do you wish to 404?")
	var url string
	fmt.Scanln(&url)
	fmt.Println("How many requests do you wish to make? Be reasonable.")
	var num int
	fmt.Scanln(&num)

	sum := 0
	for i := 0; i < num; i++ {
		sum += i
		//fmt.Println(url + "/" + randomizer(64))
		var randurl string = url + "/" + randomizer(rand.Intn(128))
		fmt.Println(randurl)

		resp, err := http.Get("http://" + randurl)
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		//body, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	log.Fatalln(err)

	}
}

//log.Println(string(body))
//}
