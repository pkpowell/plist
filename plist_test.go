package plist

import (
	"fmt"
	"strings"
	"testing"

	plist "github.com/micromdm/plist"
	hplist "howett.net/plist"
)

const data = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<array>
	<dict>
		<key>_SPCommandLineArguments</key>
		<array>
			<string>/usr/sbin/system_profiler</string>
			<string>-nospawn</string>
			<string>-xml</string>
			<string>SPEthernetDataType</string>
			<string>-detailLevel</string>
			<string>full</string>
		</array>
		<key>_SPCompletionInterval</key>
		<real>0.038819074630737305</real>
		<key>_SPResponseTime</key>
		<real>0.064825057983398438</real>
		<key>_dataType</key>
		<string>SPEthernetDataType</string>
		<key>_detailLevel</key>
		<integer>-1</integer>
		<key>_items</key>
		<array/>
		<key>_parentDataType</key>
		<string>SPHardwareDataType</string>
		<key>_properties</key>
		<dict>
			<key>_name</key>
			<dict>
				<key>_isColumn</key>
				<string>YES</string>
				<key>_order</key>
				<string>0</string>
			</dict>
			<key>spethernet_BSD_Device_Name</key>
			<dict>
				<key>_order</key>
				<string>120</string>
			</dict>
			<key>spethernet_avb_support</key>
			<dict>
				<key>_order</key>
				<string>140</string>
			</dict>
			<key>spethernet_bus</key>
			<dict>
				<key>_isColumn</key>
				<string>YES</string>
				<key>_order</key>
				<string>5</string>
			</dict>
			<key>spethernet_device-id</key>
			<dict>
				<key>_order</key>
				<string>20</string>
			</dict>
			<key>spethernet_driver</key>
			<dict>
				<key>_order</key>
				<string>110</string>
			</dict>
			<key>spethernet_firmware_version</key>
			<dict>
				<key>_order</key>
				<string>105</string>
			</dict>
			<key>spethernet_mac_address</key>
			<dict>
				<key>_order</key>
				<string>130</string>
			</dict>
			<key>spethernet_max_link_speed</key>
			<dict>
				<key>_order</key>
				<string>150</string>
			</dict>
			<key>spethernet_pcie_link-speed</key>
			<dict>
				<key>_order</key>
				<string>90</string>
			</dict>
			<key>spethernet_pcie_link-width</key>
			<dict>
				<key>_order</key>
				<string>100</string>
			</dict>
			<key>spethernet_pcie_slot</key>
			<dict>
				<key>_isColumn</key>
				<string>YES</string>
				<key>_order</key>
				<string>103</string>
			</dict>
			<key>spethernet_product-id</key>
			<dict>
				<key>_order</key>
				<string>30</string>
			</dict>
			<key>spethernet_product_name</key>
			<dict>
				<key>_isColumn</key>
				<string>NO</string>
				<key>_order</key>
				<string>10</string>
			</dict>
			<key>spethernet_revision-id</key>
			<dict>
				<key>_order</key>
				<string>60</string>
			</dict>
			<key>spethernet_rom-revision</key>
			<dict>
				<key>_order</key>
				<string>80</string>
			</dict>
			<key>spethernet_subsystem-id</key>
			<dict>
				<key>_order</key>
				<string>50</string>
			</dict>
			<key>spethernet_subsystem-vendor-id</key>
			<dict>
				<key>_order</key>
				<string>40</string>
			</dict>
			<key>spethernet_usb_device_speed</key>
			<dict>
				<key>_order</key>
				<string>85</string>
			</dict>
			<key>spethernet_vendor-id</key>
			<dict>
				<key>_order</key>
				<string>15</string>
			</dict>
			<key>spethernet_vendor_name</key>
			<dict>
				<key>_isColumn</key>
				<string>NO</string>
				<key>_order</key>
				<string>8</string>
			</dict>
			<key>spethernet_version</key>
			<dict>
				<key>_order</key>
				<string>70</string>
			</dict>
			<key>volumes</key>
			<dict>
				<key>_detailLevel</key>
				<string>0</string>
			</dict>
		</dict>
		<key>_timeStamp</key>
		<date>2024-12-22T20:51:33Z</date>
		<key>_versionInfo</key>
		<dict>
			<key>com.apple.SystemProfiler.SPEthernetReporter</key>
			<string>200</string>
		</dict>
	</dict>
</array>
</plist>`

type SPEthernetDataType struct {
	Name                      string `json:"_name"`
	EthernetAVBSupport        string `json:"spethernet_avb_support"`
	EthernetBSDDeviceName     string `json:"spethernet_BSD_Device_Name"`
	EthernetBus               string `json:"spethernet_bus"`
	EthernetDeviceID          string `json:"spethernet_device-id"`
	EthernetDriver            string `json:"spethernet_driver-id"`
	EthernetMACAddress        string `json:"spethernet_mac_address-id"`
	EthernetMaxLinkSpeed      string `json:"spethernet_max_link_speed"`
	EthernetPCIELinkSpeed     string `json:"spethernet_pcie_link-speed"`
	EthernetPCIELinkWidth     string `json:"spethernet_pcie_link-width"`
	EthernetRevisionID        string `json:"spethernet_revision-id"`
	EthernetSubsystemID       string `json:"spethernet_subsystem-id"`
	EthernetSubsystemVendorID string `json:"spethernet_subsystem-vendor-id"`
	EthernetVendorID          string `json:"spethernet_vendor-id"`
}

// ExampleUnmarshaler demonstrates using structs that use the Unmarshaler interface.
func MUnmarshaler() {
	var d []SPEthernetDataType
	err := plist.Unmarshal([]byte(data), &d)
	if err != nil {
		fmt.Println(err)
		return
	}

	// typeA, ok := decider.ActualType.(*TypeA)
	// if !ok {
	// 	fmt.Println("actual type is not TypeA")
	// 	return
	// }

	fmt.Println(d)
	// Output: VALUE-A
}

func HUnmarshaler() {
	var d []SPEthernetDataType
	err := hplist.NewDecoder(strings.NewReader(data)).Decode(&d)
	if err != nil {
		fmt.Println("getDSCL Unmarshal error", err.Error())
		// return nil
	}
	fmt.Println(d)
}

func BenchmarkMunmarhaler_test(b *testing.B) {
	for range b.N {
		MUnmarshaler()
	}
}

func BenchmarkHunmarhaler(b *testing.B) {
	for range b.N {
		HUnmarshaler()
	}
}
