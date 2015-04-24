package teeloader

/*
#cgo CFLAGS: -O2 -Wall 
#cgo LDFLAGS: -lhid -lsetupapi
#include "load_linux_only.h"
*/
import "C"
import "fmt"
import "unsafe"
import "strings"

func Teensy_load(mmcu string, vendor_id string, device_id string, hex_path string, wait int, hard_reboot int, no_reboot int, verbose int) int {
	fmt.Println("\n_____ Starting GO Teensy loader wrapper \n")
	fmt.Println("Hex file: \t", hex_path)
	fmt.Println("Mcu: \t\t", mmcu)
	fmt.Println("Vendor Id: \t", vendor_id)
	fmt.Println("Device Id: \t", device_id)
	fmt.Println("Wait: \t\t", wait)
	fmt.Println("Hard reboot: \t", hard_reboot)
	fmt.Println("No reboot: \t", no_reboot)
	fmt.Println("Verbose: \t", verbose)
	fmt.Println("\n_____ calling base routine ")
	result := C.load_teensy_with_options(C.CString(mmcu), C.CString(vendor_id), C.CString(device_id), C.CString(hex_path), C.int(wait), C.int(hard_reboot), C.int(no_reboot), C.int(verbose))
	if result == 0 {
		fmt.Println("\n---> Firmware loaded. \n")
	}else{
		fmt.Println("\n---> Firmware not loaded. \n")
	}
	return 0
}