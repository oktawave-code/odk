# InstanceScheduler

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID | [optional] [default to null]
**Name** | **string** | Name | [optional] [default to null]
**Instance** | [***BaseResource**](BaseResource.md) | Instance | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | Creation date | [optional] [default to null]
**LastChangeDate** | [**time.Time**](time.Time.md) | Last change date | [optional] [default to null]
**CreationUser** | [***UserResource**](UserResource.md) | User who created the scheduler | [optional] [default to null]
**Type_** | [***DictionaryItem**](DictionaryItem.md) | Type | [optional] [default to null]
**Status** | [***DictionaryItem**](DictionaryItem.md) | Status | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | Start date | [optional] [default to null]
**TimeZoneName** | **string** | Time zone name | [optional] [default to null]
**CycleType** | [***DictionaryItem**](DictionaryItem.md) | Cycle type | [optional] [default to null]
**CycleNumber** | **int32** | Cycle number beetwen scheduler launch | [optional] [default to null]
**ActionType** | [***DictionaryItem**](DictionaryItem.md) | Action type | [optional] [default to null]
**NewInstanceType** | [***DictionaryItem**](DictionaryItem.md) | New instance type. In case of \&quot;Instance type change\&quot; action type. | [optional] [default to null]
**SnapshotName** | **string** | Snapshot name. In case of \&quot;Create snapshot\&quot; action type. | [optional] [default to null]
**SnapshotDescription** | **string** | Snapshot description. In case of \&quot;Create snapshot\&quot; action type. | [optional] [default to null]
**Snapshot** | [***BaseResource**](BaseResource.md) | Snashot. In case of \&quot;Delete snapshot\&quot; action type. | [optional] [default to null]
**CloneName** | **string** | Clone name. In case of \&quot;Clone machine\&quot; action type. | [optional] [default to null]
**StoragePath** | **string** | Storage path. In case of \&quot;Create backup\&quot; action type. | [optional] [default to null]
**StorageProjectId** | **string** | Storage project id. In case of \&quot;Create backup\&quot; action type. | [optional] [default to null]
**IsBackupsDaysLimit** | **bool** | Is backups days limit set. In case of \&quot;Create backup\&quot; action type. | [optional] [default to null]
**BackupsDaysLimit** | **int32** | Backups days limit. In case of \&quot;Create backup\&quot; action type. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


