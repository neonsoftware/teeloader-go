#define TEENSY_VENDOR_ID 0x16c0


#ifdef USE_LIBUSB

#include "libusb.h"

int getConnectedTeensys (char * out)
{
	int teensy_found = 0;

	// discover devices
	libusb_device **list;
	libusb_context *context = NULL;

	int result = libusb_init(&context); 

	if (0 > result)
 	{
   		fprintf(stderr, "libusb_init() failed with %d.\n", result);
   		return -1;
 	}

	ssize_t cnt = libusb_get_device_list(context, &list);

	//printf("%lu devices found\n", cnt);

	ssize_t i = 0;
	
	for (i = 0; i < cnt; i++) {
	    libusb_device *device = list[i];
	    struct libusb_device_descriptor desc;
	    if ( ( libusb_get_device_descriptor(device, &desc) == 0 ) && ( desc.idVendor == TEENSY_VENDOR_ID ) )
    	{
    		printf("[%zd] Teensy : VendorId : 0x%.4x - ProductId : 0x%.4x\n", i, desc.idVendor, desc.idProduct);
    		sprintf(out, "0x%.4x#", desc.idProduct);
    		teensy_found ++;
    		out += 7;
    	}
	}
	// if (found) {
	//     libusb_device_handle *handle;
	//     err = libusb_open(found, &handle);
	//     if (err)
	//         perror("Danger !");
	//     // etc
	// }
	libusb_free_device_list(list, 1);
	return teensy_found;
}

#endif






#ifdef USE_WIN32

int getConnectedTeensys (char * out)
{
	sprintf(out, "0x0001#0x0002");
	return 2;
}

#endif