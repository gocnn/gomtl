//go:build !darwin
// +build !darwin

package gomtl

import "errors"

// Device is an abstract representation of the GPU and serves as the primary interface for a Metal app.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice
type Device struct {
	device uintptr

	// Headless indicates whether a device is configured as headless.
	Headless bool

	// LowPower indicates whether a device is low-power.
	LowPower bool

	// Removable determines whether or not a GPU is removable.
	Removable bool

	// RegistryID is the registry ID value for the device.
	RegistryID uint64

	// Name is the name of the device.
	Name string
}

// CreateSystemDefaultDevice returns the preferred system default Metal device.
//
// Reference: https://developer.apple.com/documentation/metal/1433401-mtlcreatesystemdefaultdevice
func CreateSystemDefaultDevice() (Device, error) {
	return Device{}, errors.New("gomtl: Metal is only supported on macOS")
}

// CopyAllDevices returns all Metal devices in the system.
//
// Reference: https://developer.apple.com/documentation/metal/1433367-mtlcopyalldevices
func CopyAllDevices() []Device {
	return nil
}

// Device returns the underlying id<MTLDevice> pointer.
func (d Device) Device() uintptr {
	return d.device
}

// SupportsFamily reports whether the device supports the feature set of the GPU family.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/3143473-supportsfamily
func (d Device) SupportsFamily(gf GPUFamily) bool {
	return false
}
