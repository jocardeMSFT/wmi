// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_IDEControllerDevice struct
type Win32_IDEControllerDevice struct {
	*CIM_ControlledBy
}

func NewWin32_IDEControllerDeviceEx1(instance *cim.WmiInstance) (newInstance *Win32_IDEControllerDevice, err error) {
	tmp, err := NewCIM_ControlledByEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_IDEControllerDevice{
		CIM_ControlledBy: tmp,
	}
	return
}

func NewWin32_IDEControllerDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_IDEControllerDevice, err error) {
	tmp, err := NewCIM_ControlledByEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_IDEControllerDevice{
		CIM_ControlledBy: tmp,
	}
	return
}
