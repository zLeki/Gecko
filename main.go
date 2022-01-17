package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

var words = []string{"virus!!!", "asm monkey", "leki shit coder fr fr", "github.com/zLeki", "idk", "buy my merch"}

func title() {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(words))
	clear()
	color.Green(`
		
                                 _                 
                                | |               
		  __ _  ___  ___| | _____      
		 / _ |/ _ \/ __| |/ / _ \               
		| (_| |  __/ (__|   < (_)| 	
		\__, |\___|\___|_|\_\___/   
		__/ |					       
		|___/					      

		Gecko • Created by Leki#6796
                        ` + words[random] + `
	`)

}
func clear() {
	for i := 0; i < 100; i++ {
		fmt.Print("\n")
	}
}
func dmSpam() {
	type settings struct {
		Token   string `json:"token"`
		Guild   string `json:"guild"`
		Threads int    `json:"threads"`
	}
	var data settings
	dataFile, _ := ioutil.ReadFile("settings.json")
	err := json.Unmarshal(dataFile, &data)
	if err != nil {
		log.Fatalf("error unmarshaling settings.json: %v", err)
	}

}
func settings() {

}
func menu() {
	title()
	color.Yellow("[1] Dm spam            [2] Settings")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	if input.Text() == "1" {
		title()
		for i := 0; i < 10; i++ {
			go dmSpam()
		}
		for {
			time.Sleep(900 * time.Second)
		}
	} else if input.Text() == "2" {
		settings()
	}
}
func login() {
	type jsonData struct {
		Users []struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"users"`
	}

	color.Yellow("[i] Email address: >>> ") // would use a sql database but i dont want a monkey with my oracle online login
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	email := input.Text()
	color.Yellow("[i] Password: >>> ")
	input2 := bufio.NewScanner(os.Stdin)
	input2.Scan()
	password := input2.Text()
	jsonFile, _ := ioutil.ReadFile("users.json")
	var data jsonData
	err := json.Unmarshal(jsonFile, &data)
	if err != nil {
		log.Fatalf("Error occured while decoding json: %v", err)
	}
	for _, v := range data.Users {
		if v.Email == email {
			if v.Password == password {
				title()
				color.Green("[V] Welcome back, %v", v.Email)
				time.Sleep(2 * time.Second)
				menu()
				return
			}
		}
	}
	log.Println("Unknown email or password")
	time.Sleep(3 * time.Second)
	os.Exit(0)
}
func main() {
	title()
	login()
}
