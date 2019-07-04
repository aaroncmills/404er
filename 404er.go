package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var alphanum = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@$^&*()_+-=,.")

func randomizer(r int) string {
	b := make([]rune, r)
	for i := range b {
		b[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("What domain do you wish to 404?")
	var url string
	fmt.Scanln(&url)
	fmt.Println("How many requests do you wish to make? Be reasonable")
	var num int
	fmt.Scanln(&num)

	start := time.Now()

	sum := 0
	failed := 0
	success := 0
	forbidden := 0
	percent := 0.0

	for i := 0; i < num; i++ {
		sum += i
		var randurl string = url + "/" + randomizer(rand.Intn(128))
		resp, err := http.Get("http://" + randurl)
		if err != nil {
			fmt.Println(err)
		}
		if resp.StatusCode == 404 {
			failed++
		}
		if resp.StatusCode == 200 {
			success++
		}
		if resp.StatusCode == 403 {
			forbidden++
		}
		fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode), "  ", randurl)
		defer resp.Body.Close()
	}
	nf := float64(num)
	ff := float64(failed)
	percent = ff / nf
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("You had", num, "requests which completed in", elapsed, "Of those,", failed, "were 404's. You had", success, "valid URL's hit. You had a", percent*100, "%", "404 rate.")
	if forbidden > 0 {
		fmt.Println("You had", forbidden, "403 URL hits. These are worth a closer look...")
	}
}
