// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_QueryNicSwitchHardwareCapabilities struct
type MSNdis_QueryNicSwitchHardwareCapabilities struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string
}

func NewMSNdis_QueryNicSwitchHardwareCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSNdis_QueryNicSwitchHardwareCapabilities, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_QueryNicSwitchHardwareCapabilities{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_QueryNicSwitchHardwareCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_QueryNicSwitchHardwareCapabilities, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_QueryNicSwitchHardwareCapabilities{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_QueryNicSwitchHardwareCapabilities) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_QueryNicSwitchHardwareCapabilities) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_QueryNicSwitchHardwareCapabilities) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_QueryNicSwitchHardwareCapabilities) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

//

// <param name="Header" type="MSNdis_WmiMethodHeader "></param>

// <param name="NicSwitchHardwareCapabilities" type="MSNdis_NicSwitchCapabilities "></param>
func (instance *MSNdis_QueryNicSwitchHardwareCapabilities) WmiQueryNICSwitchHardwareCapabilities( /* IN */ Header MSNdis_WmiMethodHeader,
	/* OUT */ NicSwitchHardwareCapabilities MSNdis_NicSwitchCapabilities) (err error) {
	_, err = instance.InvokeMethod("WmiQueryNICSwitchHardwareCapabilities", Header)
	if err != nil {
		return
	}
	return

}
