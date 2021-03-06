/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Run import network command
type RunImportNetworkCommand struct {

	// Is public network
	Public bool `json:"Public"`

	// Id of public ip - only for public network. When null, random address will be assigned.
	PublicIpId int32 `json:"PublicIpId,omitempty"`

	// Id of Opn - only for private network
	OpnId int32 `json:"OpnId,omitempty"`
}
