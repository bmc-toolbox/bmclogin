package main

import (
	"fmt"

	"github.com/bmc-toolbox/bmclib/devices"
	"github.com/bmc-toolbox/bmclogin"
)

func main() {

	credentials := []map[string]string{
		map[string]string{"user2": "password"},
		map[string]string{"Administrator": "password"},
		map[string]string{"root": "calvin"},
		map[string]string{"ADMIN": "ADMIN"},
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
		case devices.BmcChassis:
			connection.(devices.BmcChassis).Close()
		}

		fmt.Println("Successful login")
	} else {
		fmt.Println("login failed")
	}

	fmt.Printf("%+v\n", loginInfo)
}
