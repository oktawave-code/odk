/*
 * Oktawave API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package odk

// Create disk command
type CreateDiskCommand struct {

	// Name of disk
	DiskName string `json:"DiskName"`

	// Space capacity in GB
	SpaceCapacity int32 `json:"SpaceCapacity"`

	// Tier id
	TierId int32 `json:"TierId"`

	// Shared disk type, null if disk is not shared
	SharedDiskTypeId int32 `json:"SharedDiskTypeId,omitempty"`

	// Subregion identifier
	SubregionId int32 `json:"SubregionId,omitempty"`

	// Instance ids list
	InstanceIdsList []int32 `json:"InstanceIdsList,omitempty"`
}
