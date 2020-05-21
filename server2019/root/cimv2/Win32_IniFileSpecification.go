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

// Win32_IniFileSpecification struct
type Win32_IniFileSpecification struct {
	*CIM_FileSpecification

	//
	Action uint16

	//
	IniFile string

	//
	key string

	//
	Section string

	//
	Value string
}

func NewWin32_IniFileSpecificationEx1(instance *cim.WmiInstance) (newInstance *Win32_IniFileSpecification, err error) {
	tmp, err := NewCIM_FileSpecificationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_IniFileSpecification{
		CIM_FileSpecification: tmp,
	}
	return
}

func NewWin32_IniFileSpecificationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_IniFileSpecification, err error) {
	tmp, err := NewCIM_FileSpecificationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_IniFileSpecification{
		CIM_FileSpecification: tmp,
	}
	return
}

// SetAction sets the value of Action for the instance
func (instance *Win32_IniFileSpecification) SetPropertyAction(value uint16) (err error) {
	return instance.SetProperty("Action", value)
}

// GetAction gets the value of Action for the instance
func (instance *Win32_IniFileSpecification) GetPropertyAction() (value uint16, err error) {
	retValue, err := instance.GetProperty("Action")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIniFile sets the value of IniFile for the instance
func (instance *Win32_IniFileSpecification) SetPropertyIniFile(value string) (err error) {
	return instance.SetProperty("IniFile", value)
}

// GetIniFile gets the value of IniFile for the instance
func (instance *Win32_IniFileSpecification) GetPropertyIniFile() (value string, err error) {
	retValue, err := instance.GetProperty("IniFile")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *Win32_IniFileSpecification) SetPropertykey(value string) (err error) {
	return instance.SetProperty("key", value)
}

// Getkey gets the value of key for the instance
func (instance *Win32_IniFileSpecification) GetPropertykey() (value string, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSection sets the value of Section for the instance
func (instance *Win32_IniFileSpecification) SetPropertySection(value string) (err error) {
	return instance.SetProperty("Section", value)
}

// GetSection gets the value of Section for the instance
func (instance *Win32_IniFileSpecification) GetPropertySection() (value string, err error) {
	retValue, err := instance.GetProperty("Section")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValue sets the value of Value for the instance
func (instance *Win32_IniFileSpecification) SetPropertyValue(value string) (err error) {
	return instance.SetProperty("Value", value)
}

// GetValue gets the value of Value for the instance
func (instance *Win32_IniFileSpecification) GetPropertyValue() (value string, err error) {
	retValue, err := instance.GetProperty("Value")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}