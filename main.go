/*
Copyright © 2023 Ashwin Singh <ashwinsinghsingh666672@gmail.com>
*/
package main

import (
	"github.com/90r1ll4/gotldwizard/cmd"

	"github.com/fatih/color"
)

func logo() {
	color.Blue(`
	#   _______        ______       _  _  _ _____ ______ _______  ______ ______   #
	#      |    |      |     \      |  |  |   |    ____/ |_____| |_____/ |     \  #
	#      |    |_____ |_____/      |__|__| __|__ /_____ |     | |    \_ |_____/  #
	#                                                                             #
	`)
	color.Red("								By:- Ashwin (@90r1ll4)")

}

func main() {
	logo()
	cmd.Execute()
}
