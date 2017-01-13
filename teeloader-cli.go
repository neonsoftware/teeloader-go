package teeloader

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {

	// This software structure is described in the README of the command line interface library github.com/codegangsta/cli
	app := cli.NewApp()
	app.Name = "teeloader-osx-cli"
	app.Author = "Francesco Cervigni from Paul J Stoffregen's Teensy Loader 2.0"
	app.Usage = "Teensy firmware loader"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "mcu",
			Value: "mk20dx256",
			Usage: "Teensy chip",
		},
		cli.StringFlag{
			Name:  "file, f",
			Usage: "Input hefx file",
		},
		cli.StringFlag{
			Name:  "vendorId",
			Value: "0x16c0",
			Usage: "USB Vendor ID",
		},
		cli.StringFlag{
			Name:  "productId",
			Value: "0x0486",
			Usage: "USB Product ID",
		},
		cli.BoolFlag{
			Name:  "verbose, V",
			Usage: "Verbose",
		},
		cli.BoolFlag{
			Name:  "wait, w",
			Usage: "Wait",
		},
		cli.BoolFlag{
			Name:  "hardreboot, r",
			Usage: "Hard Reboot",
		},
		cli.BoolFlag{
			Name:  "noreboot, n",
			Usage: "No Reboot",
		},
		cli.BoolFlag{
			Name:  "l",
			Usage: "List all Teensy USB devices",
		},
	}
	app.Action = func(c *cli.Context) {

		verbose := 0
		wait := 1
		hardreboot := 0
		noreboot := 0
		listActive := false

		if c.Bool("l") {
			fmt.Println("\n\n\tList mode active.\n")
			//teensys := teeloader.GetConnectedUSBTeensy()
			//fmt.Println("\tTeensy Devices List : Size : ", len(teensys), " - Strings : ", teensys, "\n" )
			listActive = true
		}

		if c.String("file") == "" && !listActive {
			fmt.Println("\n\tPlease specify input hex file with --file or -f\n")
			return
		}

		if c.Bool("V") {
			fmt.Println("\n\tVerbose mode active.\n")
			verbose = 1
		}
		if c.Bool("w") {
			fmt.Println("\n\tVerbose mode active.\n")
			wait = 1
		}
		if c.Bool("r") {
			fmt.Println("\n\tVerbose mode active.\n")
			hardreboot = 1
		}
		if c.Bool("n") {
			fmt.Println("\n\tVerbose mode active.\n")
			noreboot = 1
		}

		Teensy_load(c.String("mcu"), c.String("vendorId"), c.String("productId"), c.String("file"), wait, hardreboot, noreboot, verbose)
	}

	app.Run(os.Args)
}
