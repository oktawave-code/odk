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

// Private IP address
type PrivateIp struct {

	// Interface id
	InterfaceId int32 `json:"InterfaceId,omitempty"`

	// The MAC address of the network card associated with that IP address
	MacAddress string `json:"MacAddress,omitempty"`

	// Adres IPv4
	Address string `json:"Address,omitempty"`

	// Adres IPv6
	AddressV6 string `json:"AddressV6,omitempty"`

	// Instance
	Instance *BaseResource `json:"Instance,omitempty"`

	// Creation date
	CreationDate time.Time `json:"CreationDate,omitempty"`
}