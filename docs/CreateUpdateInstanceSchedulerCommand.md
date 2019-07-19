# CreateUpdateInstanceSchedulerCommand

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | [default to null]
**TypeId** | **int32** | Type id | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Start date | [default to null]
**CycleTypeId** | **int32** | Cycle type id | [optional] [default to null]
**CycleNumber** | **int32** | Cycle number beetwen scheduler launch | [optional] [default to null]
**ActionTypeId** | **int32** | Action type | [default to null]
**NewInstanceTypeId** | **int32** | New instance type. In case of \&quot;Instance type change\&quot; action type. | [optional] [default to null]
**SnapshotName** | **string** | Snapshot name. In case of \&quot;Create snapshot\&quot; action type. | [optional] [default to null]
**SnapshotDescription** | **string** | Snapshot description. In case of \&quot;Create snapshot\&quot; action type. | [optional] [default to null]
**SnapshotId** | **int32** | Snashot. In case of \&quot;Delete snapshot\&quot; action type. | [optional] [default to null]
**CloneName** | **string** | Clone name. In case of \&quot;Clone machine\&quot; action type. | [optional] [default to null]
**StoragePath** | **string** | Storage path. In case of \&quot;Create backup\&quot; action type. | [optional] [default to null]
**StorageProjectId** | **string** | Storage project id. In case of \&quot;Create backup\&quot; action type. | [optional] [default to null]
**IsBackupsDaysLimit** | **bool** | Is backups days limit set. In case of \&quot;Create backup\&quot; action type. | [optional] [default to null]
**BackupsDaysLimit** | **int32** | Backups days limit. In case of \&quot;Create backup\&quot; action type. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


