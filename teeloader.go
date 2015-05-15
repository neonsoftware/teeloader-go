package teeloader

/*
#cgo darwin CFLAGS: -DUSE_LIBUSB  -I/usr/local/include -Wall -O2
#cgo darwin  LDFLAGS:  -framework IOKit  -lobjc -framework CoreFoundation /usr/local/lib/libusb.a /usr/local/lib/libusb-1.0.a
#cgo linux CFLAGS: -DUSE_LIBUSB  -Wall -O2
#cgo linux,386 LDFLAGS:  ./libusb-dep/libusb-0.1-linux-x86.a ./libusb-dep/libusb-1.0-linux-x86.a -ludev
#cgo linux,amd64 LDFLAGS:  ./libusb-dep/libusb-0.1-linux-x64.a ./libusb-dep/libusb-1.0-linux-x64.a -ludev
#cgo windows CFLAGS: -O2 -Wall -DUSE_WIN32 
#cgo windows LDFLAGS: -lhid -lsetupapi
#include "load_linux_only.h"
#include "get_usb_devices.h"
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

func GetConnectedUSBTeensy() []string {

	buffer := make([]byte, 256)
	productIds := []string {}

	teensys := C.getConnectedTeensys( (*C.char)(unsafe.Pointer(&buffer[0])) )

	if ( teensys > 0 ){
		for _, productId := range strings.Split( string(buffer) ,"#") {
			if ( productId[0] == '0' ){
				//fmt.Println("Adding : ", productId, " which is not null")
				productIds = append ( productIds , productId )
			}
		}
	}

	return productIds
}