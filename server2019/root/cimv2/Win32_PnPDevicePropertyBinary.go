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

// Win32_PnPDevicePropertyBinary struct
type Win32_PnPDevicePropertyBinary struct {
	*Win32_PnPDeviceProperty

	//
	Data []uint8
}

func NewWin32_PnPDevicePropertyBinaryEx1(instance *cim.WmiInstance) (newInstance *Win32_PnPDevicePropertyBinary, err error) {
	tmp, err := NewWin32_PnPDevicePropertyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PnPDevicePropertyBinary{
		Win32_PnPDeviceProperty: tmp,
	}
	return
}

func NewWin32_PnPDevicePropertyBinaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PnPDevicePropertyBinary, err error) {
	tmp, err := NewWin32_PnPDevicePropertyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PnPDevicePropertyBinary{
		Win32_PnPDeviceProperty: tmp,
	}
	return
}

// SetData sets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyBinary) SetPropertyData(value []uint8) (err error) {
	return instance.SetProperty("Data", value)
}

// GetData gets the value of Data for the instance
func (instance *Win32_PnPDevicePropertyBinary) GetPropertyData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}