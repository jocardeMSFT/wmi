// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_SyntheticDisplayControllerSettingData struct
type Msvm_SyntheticDisplayControllerSettingData struct {
	*CIM_ResourceAllocationSettingData

	//
	HorizontalResolution uint16

	//
	ResolutionType uint8

	//
	VerticalResolution uint16
}

func NewMsvm_SyntheticDisplayControllerSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_SyntheticDisplayControllerSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticDisplayControllerSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_SyntheticDisplayControllerSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SyntheticDisplayControllerSettingData, err error) {
	tmp, err := NewCIM_ResourceAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SyntheticDisplayControllerSettingData{
		CIM_ResourceAllocationSettingData: tmp,
	}
	return
}

// SetHorizontalResolution sets the value of HorizontalResolution for the instance
func (instance *Msvm_SyntheticDisplayControllerSettingData) SetPropertyHorizontalResolution(value uint16) (err error) {
	return instance.SetProperty("HorizontalResolution", value)
}

// GetHorizontalResolution gets the value of HorizontalResolution for the instance
func (instance *Msvm_SyntheticDisplayControllerSettingData) GetPropertyHorizontalResolution() (value uint16, err error) {
	retValue, err := instance.GetProperty("HorizontalResolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResolutionType sets the value of ResolutionType for the instance
func (instance *Msvm_SyntheticDisplayControllerSettingData) SetPropertyResolutionType(value uint8) (err error) {
	return instance.SetProperty("ResolutionType", value)
}

// GetResolutionType gets the value of ResolutionType for the instance
func (instance *Msvm_SyntheticDisplayControllerSettingData) GetPropertyResolutionType() (value uint8, err error) {
	retValue, err := instance.GetProperty("ResolutionType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVerticalResolution sets the value of VerticalResolution for the instance
func (instance *Msvm_SyntheticDisplayControllerSettingData) SetPropertyVerticalResolution(value uint16) (err error) {
	return instance.SetProperty("VerticalResolution", value)
}

// GetVerticalResolution gets the value of VerticalResolution for the instance
func (instance *Msvm_SyntheticDisplayControllerSettingData) GetPropertyVerticalResolution() (value uint16, err error) {
	retValue, err := instance.GetProperty("VerticalResolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
func (instance *Msvm_SyntheticDisplayControllerSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}