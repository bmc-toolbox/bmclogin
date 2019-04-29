package main

import (
	"fmt"

	"github.com/bmc-toolbox/bmclib/devices"
	"github.com/bmc-toolbox/bmclogin"
)

func main() {

	credentials := []map[string]string{
		{"user2": "password"},
		{"Administrator": "password"},
		{"root": "calvin"},
		{"ADMIN": "ADMIN"},
	}

	c := bmclogin.Params{
		IpAddresses:     []string{"10.193.251.68", "10.193.251.161"},
		Credentials:     credentials,
		CheckCredential: true,
		Retries:         1,
	}

	connection, loginInfo, err := c.Login()
	if err == nil {

		switch (connection).(type) {
		case devices.Bmc:
			connection.(devices.Bmc).Close()
		case devices.Cmc:
			connection.(devices.Cmc).Close()
		}

		fmt.Println("Successful login")
	} else {
		fmt.Println("login failed")
	}

	fmt.Printf("%+v\n", loginInfo)
}
