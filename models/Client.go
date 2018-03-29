package models

import (
	"fmt"
)

type Client struct {
	name     string
	account  string
	password string
	balance  string
	active   string
}

func (c *Client) CheckClient(account string, password string) bool {
	if account == c.account {
		if password == c.password {
			fmt.Println("Welcome " + c.name)
		}
		return true
	} else {
		return false
	}

}

func (c *Client) AddClient(l [5]string) {
	c.name = l[0]
	c.account = l[1]
	c.password = l[2]
	c.balance = l[3]
	c.active = l[4]
}
