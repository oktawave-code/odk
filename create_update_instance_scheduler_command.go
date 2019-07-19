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

// Create/update instance scheduler command
type CreateUpdateInstanceSchedulerCommand struct {

	// Name
	Name string `json:"Name"`

	// Type id
	TypeId int32 `json:"TypeId"`

	// Start date
	StartDate time.Time `json:"StartDate"`

	// Cycle type id
	CycleTypeId int32 `json:"CycleTypeId,omitempty"`

	// Cycle number beetwen scheduler launch
	CycleNumber int32 `json:"CycleNumber,omitempty"`

	// Action type
	ActionTypeId int32 `json:"ActionTypeId"`

	// New instance type. In case of \"Instance type change\" action type.
	NewInstanceTypeId int32 `json:"NewInstanceTypeId,omitempty"`

	// Snapshot name. In case of \"Create snapshot\" action type.
	SnapshotName string `json:"SnapshotName,omitempty"`

	// Snapshot description. In case of \"Create snapshot\" action type.
	SnapshotDescription string `json:"SnapshotDescription,omitempty"`

	// Snashot. In case of \"Delete snapshot\" action type.
	SnapshotId int32 `json:"SnapshotId,omitempty"`

	// Clone name. In case of \"Clone machine\" action type.
	CloneName string `json:"CloneName,omitempty"`

	// Storage path. In case of \"Create backup\" action type.
	StoragePath string `json:"StoragePath,omitempty"`

	// Storage project id. In case of \"Create backup\" action type.
	StorageProjectId string `json:"StorageProjectId,omitempty"`

	// Is backups days limit set. In case of \"Create backup\" action type.
	IsBackupsDaysLimit bool `json:"IsBackupsDaysLimit,omitempty"`

	// Backups days limit. In case of \"Create backup\" action type.
	BackupsDaysLimit int32 `json:"BackupsDaysLimit,omitempty"`
}