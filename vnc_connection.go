/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

import (
	"time"
)

// VNC connection properties
type VncConnection struct {

	// Instance
	Instance *BaseResource `json:"Instance,omitempty"`

	// Source ip
	SourceIp string `json:"SourceIp,omitempty"`

	// Address
	Address string `json:"Address,omitempty"`

	// Port
	Port int32 `json:"Port,omitempty"`

	// Password
	Password string `json:"Password,omitempty"`

	// Open until
	OpenUntil time.Time `json:"OpenUntil,omitempty"`
}
