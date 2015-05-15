# teeloader

## Dependencies 

- OSX 
  - brew install libusb-compat

- Linux 
  - You need both the static library of libusb-0.1 and libusb-1.0.19
  - You can find them in the folder libusb-dep/linux
  - Copy *.h files to /usb/local/include 
  - On x64 : 
    - Copy libusb-dep/libusb-0.1-linux-x64.a to /usb/local/lib
    - create folder /usb/local/lib/libusb-1.0
    - Copy libusb-dep/libusb-1.0-linux-x64.a to /usb/local/lib/libusb-1.0 
  - On x86 : 
    - Copy libusb-dep/libusb-0.1-linux-x86.a to /usb/local/lib
    - create folder /usb/local/lib/libusb-1.0
    - Copy libusb-dep/libusb-1.0-linux-x86.a to /usb/local/lib/libusb-1.0 

- Windows
    - No dependencies known

## Install 

```
go get github.com/neonsoftware/teeloader
```
