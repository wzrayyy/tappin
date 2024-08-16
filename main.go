package main

import (
	"encoding/json"
	"fmt"
	"os"
	_ "sync"
	"time"
	_ "time"

	"github.com/wzrayyy/tappin/internal/clicker"
)

func die_if(err error) {
	if err != nil {
		panic(err)
	}
}

type account struct {
	Name    string
	Phone   string
	AuthKey string
	UserID  int
}

func main() {
	data, err := os.ReadFile("./accounts.json")
	die_if(err)

	var accounts []account
	json.Unmarshal(data, &accounts)
	account := accounts[0]

	c, err := clicker.NewClicker(account.AuthKey, account.UserID, clicker.Config{
		TapsPerSecond:   0,
		TapInterval:     2,
		UpdateFrequency: 3,
	})
	die_if(err)
	fmt.Println("a")

	go c.Start()

	time.Sleep(3 * time.Second)

	c.Config.TapInterval = 1
	time.Sleep(10 * time.Second)

	err = c.Stop()

	die_if(err)
}
