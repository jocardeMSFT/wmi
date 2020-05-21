// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_SystemSlot struct
type Win32_SystemSlot struct {
	*CIM_Slot

	//
	BusNumber uint32

	//
	CurrentUsage uint16

	//
	DeviceNumber uint32

	//
	FunctionNumber uint32

	//
	PMESignal bool

	//
	SegmentGroupNumber uint32

	//
	Shared bool

	//
	SlotDesignation string
}

func NewWin32_SystemSlotEx1(instance *cim.WmiInstance) (newInstance *Win32_SystemSlot, err error) {
	tmp, err := NewCIM_SlotEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_SystemSlot{
		CIM_Slot: tmp,
	}
	return
}

func NewWin32_SystemSlotEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_SystemSlot, err error) {
	tmp, err := NewCIM_SlotEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_SystemSlot{
		CIM_Slot: tmp,
	}
	return
}

// SetBusNumber sets the value of BusNumber for the instance
func (instance *Win32_SystemSlot) SetPropertyBusNumber(value uint32) (err error) {
	return instance.SetProperty("BusNumber", value)
}

// GetBusNumber gets the value of BusNumber for the instance
func (instance *Win32_SystemSlot) GetPropertyBusNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentUsage sets the value of CurrentUsage for the instance
func (instance *Win32_SystemSlot) SetPropertyCurrentUsage(value uint16) (err error) {
	return instance.SetProperty("CurrentUsage", value)
}

// GetCurrentUsage gets the value of CurrentUsage for the instance
func (instance *Win32_SystemSlot) GetPropertyCurrentUsage() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentUsage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceNumber sets the value of DeviceNumber for the instance
func (instance *Win32_SystemSlot) SetPropertyDeviceNumber(value uint32) (err error) {
	return instance.SetProperty("DeviceNumber", value)
}

// GetDeviceNumber gets the value of DeviceNumber for the instance
func (instance *Win32_SystemSlot) GetPropertyDeviceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFunctionNumber sets the value of FunctionNumber for the instance
func (instance *Win32_SystemSlot) SetPropertyFunctionNumber(value uint32) (err error) {
	return instance.SetProperty("FunctionNumber", value)
}

// GetFunctionNumber gets the value of FunctionNumber for the instance
func (instance *Win32_SystemSlot) GetPropertyFunctionNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("FunctionNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPMESignal sets the value of PMESignal for the instance
func (instance *Win32_SystemSlot) SetPropertyPMESignal(value bool) (err error) {
	return instance.SetProperty("PMESignal", value)
}

// GetPMESignal gets the value of PMESignal for the instance
func (instance *Win32_SystemSlot) GetPropertyPMESignal() (value bool, err error) {
	retValue, err := instance.GetProperty("PMESignal")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSegmentGroupNumber sets the value of SegmentGroupNumber for the instance
func (instance *Win32_SystemSlot) SetPropertySegmentGroupNumber(value uint32) (err error) {
	return instance.SetProperty("SegmentGroupNumber", value)
}

// GetSegmentGroupNumber gets the value of SegmentGroupNumber for the instance
func (instance *Win32_SystemSlot) GetPropertySegmentGroupNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("SegmentGroupNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShared sets the value of Shared for the instance
func (instance *Win32_SystemSlot) SetPropertyShared(value bool) (err error) {
	return instance.SetProperty("Shared", value)
}

// GetShared gets the value of Shared for the instance
func (instance *Win32_SystemSlot) GetPropertyShared() (value bool, err error) {
	retValue, err := instance.GetProperty("Shared")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSlotDesignation sets the value of SlotDesignation for the instance
func (instance *Win32_SystemSlot) SetPropertySlotDesignation(value string) (err error) {
	return instance.SetProperty("SlotDesignation", value)
}

// GetSlotDesignation gets the value of SlotDesignation for the instance
func (instance *Win32_SystemSlot) GetPropertySlotDesignation() (value string, err error) {
	retValue, err := instance.GetProperty("SlotDesignation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}