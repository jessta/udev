package main

import "github.com/jessta/udev"
import "os"
import "fmt"

func main() {
	u := udev.NewUdev()
	defer u.Unref()

	e := u.NewEnumerate()
	defer e.Unref()

	e.AddMatchSubsystem("hidraw")
	e.ScanDevices()

	for device := e.First(); !device.IsNil(); device = device.Next() {
		path := device.Name()
		dev := u.DeviceFromSysPath(path)

		/* usb_device_get_devnode() returns the path to the device node
		   itself in /dev. */
		fmt.Printf("Device Node Path: %s\n", dev.DevNode())

		/* The device pointed to by dev contains information about
		   the hidraw device. In order to get information about the
		   USB device, get the parent device with the
		   subsystem/devtype pair of "usb"/"usb_device". This will
		   be several levels up the tree, but the function will find
		   it.*/
		dev = dev.ParentWithSubsystemDevType("usb", "usb_device")

		if dev.IsNil() {
			fmt.Println("Unable to find parent usb device.")
			os.Exit(1)
		}

		/* From here, we can call get_sysattr_value() for each file
		   in the device's /sys entry. The strings passed into these
		   functions (idProduct, idVendor, serial, etc.) correspond
		   directly to the files in the directory which represents
		   the USB device. Note that USB strings are Unicode, UCS2
		   encoded, but the strings returned from
		   udev_device_get_sysattr_value() are UTF-8 encoded. */

		fmt.Printf("  VID/PID: %s %s\n", dev.SysAttrValue("idVendor"), dev.SysAttrValue("idProduct"))

		fmt.Printf("  %s\n  %s\n",
			dev.SysAttrValue("manufacturer"),
			dev.SysAttrValue("product"))

		fmt.Printf("  serial: %s\n", dev.SysAttrValue("serial"))
	}
}
