/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Instance autoscaling configuration
type InstanceAutoscalingParameter struct {

	// Autoscaling parameter type
	AutoscalingParameterType *DictionaryItem `json:"AutoscalingParameterType,omitempty"`

	// Minimum RAM capacity in megabytes
	MinRam int32 `json:"MinRam,omitempty"`

	// Maximum RAM capacity in megabytes
	MaxRam int32 `json:"MaxRam,omitempty"`

	// Minimum CPU count
	MinCpu int32 `json:"MinCpu,omitempty"`

	// Maximum CPU count
	MaxCpu int32 `json:"MaxCpu,omitempty"`

	// Instance class increase agreement
	HasAgreedToIncreaseClass bool `json:"HasAgreedToIncreaseClass,omitempty"`

	// Instance class decrease agreement
	HasAgreedToDecreaseClass bool `json:"HasAgreedToDecreaseClass,omitempty"`

	// Restart agreement
	HasAgreedToRestart bool `json:"HasAgreedToRestart,omitempty"`

	// Restart time lower limit
	RestartTimeFrom *NullableTime `json:"RestartTimeFrom,omitempty"`

	// Restart time upper limit
	RestartTimeTo *NullableTime `json:"RestartTimeTo,omitempty"`

	// Time zone name (https://www.iana.org/time-zones)
	TimeZoneName string `json:"TimeZoneName,omitempty"`
}
