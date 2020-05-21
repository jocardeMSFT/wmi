// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerManagerRequestGuid struct
type MSFT_ServerManagerRequestGuid struct {
	*cim.WmiInstance

	//
	HighHalf uint64

	//
	LowHalf uint64
}

func NewMSFT_ServerManagerRequestGuidEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerManagerRequestGuid, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerRequestGuid{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerManagerRequestGuidEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerManagerRequestGuid, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerManagerRequestGuid{
		WmiInstance: tmp,
	}
	return
}

// SetHighHalf sets the value of HighHalf for the instance
func (instance *MSFT_ServerManagerRequestGuid) SetPropertyHighHalf(value uint64) (err error) {
	return instance.SetProperty("HighHalf", value)
}

// GetHighHalf gets the value of HighHalf for the instance
func (instance *MSFT_ServerManagerRequestGuid) GetPropertyHighHalf() (value uint64, err error) {
	retValue, err := instance.GetProperty("HighHalf")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLowHalf sets the value of LowHalf for the instance
func (instance *MSFT_ServerManagerRequestGuid) SetPropertyLowHalf(value uint64) (err error) {
	return instance.SetProperty("LowHalf", value)
}

// GetLowHalf gets the value of LowHalf for the instance
func (instance *MSFT_ServerManagerRequestGuid) GetPropertyLowHalf() (value uint64, err error) {
	retValue, err := instance.GetProperty("LowHalf")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}