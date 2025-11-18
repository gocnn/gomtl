//go:build darwin
// +build darwin

package main

import (
	"fmt"

	"github.com/gocnn/gomtl"
)

func main() {
	device, err := gomtl.CreateSystemDefaultDevice()
	if err != nil {
		panic(err)
	}

	fmt.Println("Preferred system default Metal device:", device.Name)

	allDevices := gomtl.CopyAllDevices()

	for _, d := range allDevices {
		fmt.Println()
		printDeviceInfo(d)
	}
}

func printDeviceInfo(d gomtl.Device) {
	fmt.Println(d.Name + ":")
	fmt.Println("• Low-power:", d.LowPower)
	fmt.Println("• Removable:", d.Removable)
	fmt.Println("• Configured as headless:", d.Headless)
	fmt.Println("• Registry ID:", d.RegistryID)
	fmt.Println()
	fmt.Println("Supports GPU Families:")
	fmt.Println("• Apple family 1 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyApple1)))
	fmt.Println("• Apple family 2 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyApple2)))
	fmt.Println("• Apple family 3 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyApple3)))
	fmt.Println("• Apple family 4 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyApple4)))
	fmt.Println("• Apple family 5 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyApple5)))
	fmt.Println("• Apple family 6 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyApple6)))
	fmt.Println("• Apple family 7 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyApple7)))
	fmt.Println("• Apple family 8 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyApple8)))
	fmt.Println("• Mac family 2 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyMac2)))
	fmt.Println("• Common family 1 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyCommon1)))
	fmt.Println("• Common family 2 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyCommon2)))
	fmt.Println("• Common family 3 GPU features:", supported(d.SupportsFamily(gomtl.GPUFamilyCommon3)))
	fmt.Println("• Metal 3 features:", supported(d.SupportsFamily(gomtl.GPUFamilyMetal3)))
}

func supported(v bool) string {
	switch v {
	case true:
		return "Supported"
	case false:
		return "Unsupported"
	}

	panic("unreachable")
}
