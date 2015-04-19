package teeloader

/*
#cgo CFLAGS: -DUSE_LIBUSB  -Wall -O2
#cgo LDFLAGS:  -lusb
#include "load_linux_only.h"
*/
import "C"
import "fmt"
//import "unsafe"
//import "strings"

func Teensy_load(mmcu string, vendor_id string, device_id string, hex_path string, wait int, hard_reboot int, no_reboot int, verbose int) int {
	fmt.Println("Mcu: ", mmcu)
	fmt.Println("Hex file: ", hex_path)
	fmt.Println("Vendor Id: ", vendor_id)
	fmt.Println("Device Id: ", device_id)
	fmt.Println("Wait: ", wait)
	fmt.Println("Hard reboot: ", hard_reboot)
	fmt.Println("No reboot: ", no_reboot)
	fmt.Println("Verbose: ", verbose)
	C.load_teensy_with_options(C.CString(mmcu), C.CString(vendor_id), C.CString(device_id), C.CString(hex_path), C.int(wait), C.int(hard_reboot), C.int(no_reboot), C.int(verbose))
	return 0
}
