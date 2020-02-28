/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Create/update imap health check command
type CreateUpdateHealthCheckSipCommand struct {

	// Sip user name
	SipUserName string `json:"SipUserName,omitempty"`

	// Sip password
	SipPassword string `json:"SipPassword,omitempty"`

	// Sip domain
	SipDomain string `json:"SipDomain,omitempty"`

	// Time the server has to complete the request in [ms]
	Timeout int32 `json:"Timeout,omitempty"`

	// Health check name
	Name string `json:"Name"`

	// URL or IP address of the monitored service
	Address string `json:"Address"`

	// Time interval between health checks, in seconds
	Interval int32 `json:"Interval"`

	// Is paused
	Paused bool `json:"Paused"`

	// Use random substitute locations in case of location breakdown
	LocationsFailoverEnabled bool `json:"LocationsFailoverEnabled"`

	// Notification types enabled for a health check
	NotificationTypeIds []int32 `json:"NotificationTypeIds,omitempty"`

	// Event types with enabled notification
	NotificationEventTypeIds []int32 `json:"NotificationEventTypeIds,omitempty"`

	// Time when notification is sent
	NotificationTimeId int32 `json:"NotificationTimeId"`

	// Description
	Description string `json:"Description,omitempty"`
}
