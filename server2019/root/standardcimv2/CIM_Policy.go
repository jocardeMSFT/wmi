// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_Policy struct
type CIM_Policy struct {
	*CIM_ManagedElement

	//
	CommonName string

	//
	PolicyKeywords []string
}

func NewCIM_PolicyEx1(instance *cim.WmiInstance) (newInstance *CIM_Policy, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Policy{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_PolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Policy, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Policy{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetCommonName sets the value of CommonName for the instance
func (instance *CIM_Policy) SetPropertyCommonName(value string) (err error) {
	return instance.SetProperty("CommonName", value)
}

// GetCommonName gets the value of CommonName for the instance
func (instance *CIM_Policy) GetPropertyCommonName() (value string, err error) {
	retValue, err := instance.GetProperty("CommonName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyKeywords sets the value of PolicyKeywords for the instance
func (instance *CIM_Policy) SetPropertyPolicyKeywords(value []string) (err error) {
	return instance.SetProperty("PolicyKeywords", value)
}

// GetPolicyKeywords gets the value of PolicyKeywords for the instance
func (instance *CIM_Policy) GetPropertyPolicyKeywords() (value []string, err error) {
	retValue, err := instance.GetProperty("PolicyKeywords")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}