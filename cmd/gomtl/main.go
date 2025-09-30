package main

import (
	"fmt"

	"github.com/gocnn/gomtl"
)

func main() {
	fmt.Println(`
          _____                   _______                   _____                _____                    _____  
         /\    \                 /::\    \                 /\    \              /\    \                  /\    \ 
        /::\    \               /::::\    \               /::\____\            /::\    \                /::\____\
       /::::\    \             /::::::\    \             /::::|   |            \:::\    \              /:::/    /
      /::::::\    \           /::::::::\    \           /:::::|   |             \:::\    \            /:::/    / 
     /:::/\:::\    \         /:::/~~\:::\    \         /::::::|   |              \:::\    \          /:::/    /  
    /:::/  \:::\    \       /:::/    \:::\    \       /:::/|::|   |               \:::\    \        /:::/    /   
   /:::/    \:::\    \     /:::/    / \:::\    \     /:::/ |::|   |               /::::\    \      /:::/    /    
  /:::/    / \:::\    \   /:::/____/   \:::\____\   /:::/  |::|___|______        /::::::\    \    /:::/    /     
 /:::/    /   \:::\ ___\ |:::|    |     |:::|    | /:::/   |::::::::\    \      /:::/\:::\    \  /:::/    /      
/:::/____/  ___\:::|    ||:::|____|     |:::|    |/:::/    |:::::::::\____\    /:::/  \:::\____\/:::/____/       
\:::\    \ /\  /:::|____| \:::\    \   /:::/    / \::/    / ~~~~~/:::/    /   /:::/    \::/    /\:::\    \       
 \:::\    /::\ \::/    /   \:::\    \ /:::/    /   \/____/      /:::/    /   /:::/    / \/____/  \:::\    \      
  \:::\   \:::\ \/____/     \:::\    /:::/    /                /:::/    /   /:::/    /            \:::\    \     
   \:::\   \:::\____\        \:::\__/:::/    /                /:::/    /   /:::/    /              \:::\    \    
    \:::\  /:::/    /         \::::::::/    /                /:::/    /    \::/    /                \:::\    \   
     \:::\/:::/    /           \::::::/    /                /:::/    /      \/____/                  \:::\    \  
      \::::::/    /             \::::/    /                /:::/    /                                 \:::\    \ 
       \::::/    /               \::/____/                /:::/    /                                   \:::\____\
        \::/____/                 ~~                      \::/    /                                     \::/    /
                                                           \/____/                                       \/____/`)

	allDevices := gomtl.CopyAllDevices()
	deviceCount := len(allDevices)

	fmt.Printf("Detected %d Metal Capable device(s)\n\n", deviceCount)

	for i, d := range allDevices {
		printDeviceInfo(i, d)
		if i < deviceCount-1 {
			fmt.Println()
		}
	}

	fmt.Println("\nResult = PASS")
}

func printDeviceInfo(index int, d gomtl.Device) {
	fmt.Printf("Device %d: \"%s\"\n", index, d.Name)

	// General Device Information
	fmt.Println("  General Device Information:")
	fmt.Printf("    Low-power:                               %s\n", yesNo(d.LowPower))
	fmt.Printf("    Removable:                               %s\n", yesNo(d.Removable))
	fmt.Printf("    Configured as headless:                  %s\n", yesNo(d.Headless))
	fmt.Printf("    Registry ID:                             %d\n", d.RegistryID)

	// GPU Family Support
	fmt.Println("  GPU Family Support:")
	printGPUFamilySupport("Apple family 1 GPU features", d.SupportsFamily(gomtl.GPUFamilyApple1))
	printGPUFamilySupport("Apple family 2 GPU features", d.SupportsFamily(gomtl.GPUFamilyApple2))
	printGPUFamilySupport("Apple family 3 GPU features", d.SupportsFamily(gomtl.GPUFamilyApple3))
	printGPUFamilySupport("Apple family 4 GPU features", d.SupportsFamily(gomtl.GPUFamilyApple4))
	printGPUFamilySupport("Apple family 5 GPU features", d.SupportsFamily(gomtl.GPUFamilyApple5))
	printGPUFamilySupport("Apple family 6 GPU features", d.SupportsFamily(gomtl.GPUFamilyApple6))
	printGPUFamilySupport("Apple family 7 GPU features", d.SupportsFamily(gomtl.GPUFamilyApple7))
	printGPUFamilySupport("Apple family 8 GPU features", d.SupportsFamily(gomtl.GPUFamilyApple8))
	printGPUFamilySupport("Mac family 2 GPU features", d.SupportsFamily(gomtl.GPUFamilyMac2))
	printGPUFamilySupport("Common family 1 GPU features", d.SupportsFamily(gomtl.GPUFamilyCommon1))
	printGPUFamilySupport("Common family 2 GPU features", d.SupportsFamily(gomtl.GPUFamilyCommon2))
	printGPUFamilySupport("Common family 3 GPU features", d.SupportsFamily(gomtl.GPUFamilyCommon3))
	printGPUFamilySupport("Metal 3 features", d.SupportsFamily(gomtl.GPUFamilyMetal3))
}

func printGPUFamilySupport(feature string, supported bool) {
	status := "No"
	if supported {
		status = "Yes"
	}
	fmt.Printf("    %-40s %s\n", feature+":", status)
}

func yesNo(v bool) string {
	if v {
		return "Yes"
	}
	return "No"
}
