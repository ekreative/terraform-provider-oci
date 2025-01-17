// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/v42/common"
)

// Forecast Forecast configuration of usage/cost.
type Forecast struct {

	// forecast end time.
	TimeForecastEnded *common.SDKTime `mandatory:"true" json:"timeForecastEnded"`

	// BASIC uses ETS to project future usage/cost based on history data. The basis for projections will be a rolling set of equivalent historical days for which projection is being made.
	ForecastType ForecastForecastTypeEnum `mandatory:"false" json:"forecastType,omitempty"`

	// forecast start time. Will default to UTC-1 if not specified
	TimeForecastStarted *common.SDKTime `mandatory:"false" json:"timeForecastStarted"`
}

func (m Forecast) String() string {
	return common.PointerString(m)
}

// ForecastForecastTypeEnum Enum with underlying type: string
type ForecastForecastTypeEnum string

// Set of constants representing the allowable values for ForecastForecastTypeEnum
const (
	ForecastForecastTypeBasic ForecastForecastTypeEnum = "BASIC"
)

var mappingForecastForecastType = map[string]ForecastForecastTypeEnum{
	"BASIC": ForecastForecastTypeBasic,
}

// GetForecastForecastTypeEnumValues Enumerates the set of values for ForecastForecastTypeEnum
func GetForecastForecastTypeEnumValues() []ForecastForecastTypeEnum {
	values := make([]ForecastForecastTypeEnum, 0)
	for _, v := range mappingForecastForecastType {
		values = append(values, v)
	}
	return values
}
