package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/bmc-toolbox/bmclib/devices"
	"github.com/bmc-toolbox/bmclogin"
)

func main() {

	log := logrus.New()
	log.Out = os.Stdout

	log.SetLevel(logrus.InfoLevel)

	credentials := []map[string]string{
		map[string]string{"foo": "bar"},
		map[string]string{"Administrator": "foobar"},
		map[string]string{"User1": "blah"},
		map[string]string{"user2": "password"},
	}

	c := bmclogin.Params{
		IpAddresses:     []string{"10.193.251.68", "10.193.251.161"},
		Credentials:     credentials,
		CheckCredential: true,
		Retries:         1,
		Logger:          log,
	}

	connection, success, loginInfo := c.Login()
	if success {

		switch (connection).(type) {
		case devices.Bmc:
			connection.(devices.Bmc).Close()
		case devices.BmcChassis:
			connection.(devices.BmcChassis).Close()
		}

		log.Info("Successful login")
	} else {
		log.Error("login failed")
	}

	fmt.Printf("%+v\n", loginInfo)
}
