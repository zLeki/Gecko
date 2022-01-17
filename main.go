package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"math/rand"
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
                        `+words[random]+`
	`)

}
func clear() {
	fmt.Print("\033[H\033[2J")
}
func login() {
	color.Yellow("[i] Email address: >>>") // would use a sql database but i dont want a monkey with my oracle online login
	scan, err := fmt.Scan()
	if err != nil {
		return
	}
	log.Println(scan)
	color.Yellow("[i] Password: >>>")
	
}
func main() {
	title()
}
