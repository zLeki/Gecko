package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/zLeki/Gecko/structs"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"time"
)

var client http.Client
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

type Settings struct {
	Token       string `json:"token"`
	Guild       string `json:"guild"`
	Threads     int    `json:"threads"`
	Whitelisted string `json:"whitelisted"`
	Message     string `json:"message"`
}

func MassDM(ids []string) {
	type ChannelInfo struct {
		ID            string      `json:"id"`
		Type          int         `json:"type"`
		LastMessageID interface{} `json:"last_message_id"`
		Recipients    []struct {
			ID            string `json:"id"`
			Username      string `json:"username"`
			Avatar        string `json:"avatar"`
			Discriminator string `json:"discriminator"`
			PublicFlags   int    `json:"public_flags"`
			Bot           bool   `json:"bot"`
		} `json:"recipients"`
	}
	f, _ := ioutil.ReadFile("settings.json")
	var settings Settings
	err := json.Unmarshal(f, &settings)
	if err != nil {
		log.Fatalf("Error opening settings", err)
	}
	token := settings.Token
	for _, v := range ids {
		log.Println(v)
		var body = []byte(`
{"recipients":["` + v + `"]}
`)
		req, _ := http.NewRequest("POST", "https://discord.com/api/v9/users/@me/channels", bytes.NewBuffer(body))
		req.Header.Set("content-type", "application/json")
		req.Header.Set("Authorization", token)
		resp, _ := client.Do(req)
		var data ChannelInfo
		err := json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			log.Fatalf("Error decoding channel info: %v", err)
		}
		rand.Seed(time.Now().UnixNano())
		min := 30
		max := 90
		var body2 = []byte(`
		{
    "content": "` + settings.Message + `",
    "nonce": "23282321949104537` + strconv.Itoa(rand.Intn(max-min+1)+min) + `",
    "tts": false
}
`)
		req1, _ := http.NewRequest("POST", "https://discord.com/api/v9/channels/"+data.ID+"/messages", bytes.NewBuffer(body2))
		req1.Header.Set("content-type", "application/json")
		req1.Header.Set("Authorization", token)
		resp1, _ := client.Do(req1)
		dataByts, _ := ioutil.ReadAll(resp1.Body)
		log.Println(string(dataByts))

	}
}
func Pullids(guildid string) {
	var b = 0
	f, _ := ioutil.ReadFile("settings.json")
	var settings Settings
	err := json.Unmarshal(f, &settings)
	if err != nil {
		log.Fatalf("Error opening settings", err)
	}
	token := settings.Token
	var ids = []string{}
	req1, _ := http.NewRequest("GET", "https://discord.com/api/guilds/"+guildid+"/channels", nil)
	req1.Header.Set("content-type", "application/json")
	req1.Header.Set("Authorization", token)
	resp1, _ := client.Do(req1)
	var data1 structs.GuidldChannelData
	err = json.NewDecoder(resp1.Body).Decode(&data1)
	if err != nil {
		return
	}
	for _, v := range data1 {

		req, _ := http.NewRequest("GET", "https://discord.com/api/v9/channels/"+v.ID+"/messages?limit=100", nil)
		req.Header.Set("content-type", "application/json")
		req.Header.Set("Authorization", token)
		resp, _ := client.Do(req)
		if resp.StatusCode == 200 {

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(resp.Body)
			var data structs.ChannelData
			err := json.NewDecoder(resp.Body).Decode(&data)
			if err != nil {
				log.Fatalf("Error decoding JSON: %v", err)
			}

			for _, v := range data {
				if !contains(ids, v.Author.ID) {
					color.Green("[V] Pulled ID successfully! %v", b)
					ids = append(ids, v.Author.ID)
					b += 1
				}
			}
		}
	}
	color.Yellow("[-] Dming everyone..")
	MassDM(ids)
}
func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
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
		color.Yellow("[i] Guild ID: >>> ") // would use a sql database but i dont want a monkey with my oracle online login
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		guildid := input.Text()
		Pullids(guildid)
		time.Sleep(900 * time.Second)
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
	title()
	color.Yellow("[i] Email address: >>> ") // would use a sql database but i dont want a monkey with my oracle online login
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	email := input.Text()
	color.Yellow("[i] Password: >>> ")
	input2 := bufio.NewScanner(os.Stdin)
	input2.Scan()
	password := input2.Text()
	req, _ := http.Get("https://www.toptal.com/developers/hastebin/raw/uyojuzulir")
	var data jsonData
	err := json.NewDecoder(req.Body).Decode(&data)
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
	login()

}
