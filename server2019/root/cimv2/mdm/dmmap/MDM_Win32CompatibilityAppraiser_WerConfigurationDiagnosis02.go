// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02 struct
type MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	MostRestrictiveSetting int32

	//
	ParentID string

	//
	WerTelemetryOptIn int32
}

func NewMDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMostRestrictiveSetting sets the value of MostRestrictiveSetting for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02) SetPropertyMostRestrictiveSetting(value int32) (err error) {
	return instance.SetProperty("MostRestrictiveSetting", value)
}

// GetMostRestrictiveSetting gets the value of MostRestrictiveSetting for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02) GetPropertyMostRestrictiveSetting() (value int32, err error) {
	retValue, err := instance.GetProperty("MostRestrictiveSetting")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWerTelemetryOptIn sets the value of WerTelemetryOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02) SetPropertyWerTelemetryOptIn(value int32) (err error) {
	return instance.SetProperty("WerTelemetryOptIn", value)
}

// GetWerTelemetryOptIn gets the value of WerTelemetryOptIn for the instance
func (instance *MDM_Win32CompatibilityAppraiser_WerConfigurationDiagnosis02) GetPropertyWerTelemetryOptIn() (value int32, err error) {
	retValue, err := instance.GetProperty("WerTelemetryOptIn")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}