package models

import (
	"fmt"

	"github.com/RafaMarquesLearn/go-bank/helpers"
)

type Client struct {
	name     string
	account  int
	password string
	balance  float64
	active   bool
}

func (c Client) AddClient(l []string) Client {
	c.name = l[0]
	c.account = helpers.ConvToInt(l[1])
	c.password = l[2]
	c.balance = helpers.ConvToFloat(l[3])
	c.active = helpers.ConvToBool(l[4])
	return c
}

func (c *Client) CheckClient(account int, password string) bool {
	if account == c.account {
		if password == c.password {
			fmt.Println("Welcome " + c.name)
		}
		return true
	} else {
		return false
	}

}

func (c *Client) GetClientInfo(account int) string {
	if c.account == account {
		fmt.Println("Client name: " + c.name)
		if c.active == false {
			message := "This client account is not active!"
			return message
		} else {
			message := "This client account is active!"
			return message
		}
	} else {
		message := "No client found for this account number!"
		return message
	}
}

func (c *Client) SetClientStatus(status bool) {
	var change string
	switch status {
	case true:
		fmt.Println("This client account is active. Do you wish do deactivate it?(y or n)")
		fmt.Scanln(&change)
		if change == "y" {
			if c.balance > 0 {
				fmt.Println("This account has a positive balance. Can't deactivate it!")
			} else {
				c.active = false
				fmt.Println("Account number " + string(c.account) + " deactivated!")
			}
		}
	case false:
		fmt.Println("This client account is deactive. Do you wish do activate it?(y or n)")
		fmt.Scanln(&change)
		if change == "y" {
			c.active = true
			fmt.Println("Account number " + string(c.account) + " activated!")
		}
	}
}
