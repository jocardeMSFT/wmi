// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_Setting struct
type CIM_Setting struct {
	*cim.WmiInstance

	//
	Caption string

	//
	Description string

	//
	SettingID string
}

func NewCIM_SettingEx1(instance *cim.WmiInstance) (newInstance *CIM_Setting, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_Setting{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_SettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Setting, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Setting{
		WmiInstance: tmp,
	}
	return
}

// SetCaption sets the value of Caption for the instance
func (instance *CIM_Setting) SetPropertyCaption(value string) (err error) {
	return instance.SetProperty("Caption", value)
}

// GetCaption gets the value of Caption for the instance
func (instance *CIM_Setting) GetPropertyCaption() (value string, err error) {
	retValue, err := instance.GetProperty("Caption")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *CIM_Setting) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *CIM_Setting) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingID sets the value of SettingID for the instance
func (instance *CIM_Setting) SetPropertySettingID(value string) (err error) {
	return instance.SetProperty("SettingID", value)
}

// GetSettingID gets the value of SettingID for the instance
func (instance *CIM_Setting) GetPropertySettingID() (value string, err error) {
	retValue, err := instance.GetProperty("SettingID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}