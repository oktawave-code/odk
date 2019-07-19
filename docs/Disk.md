# Disk

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id | [optional] [default to null]
**Name** | **string** | Name | [optional] [default to null]
**SpaceCapacity** | **int32** | Space capacity in GB | [optional] [default to null]
**Tier** | [***DictionaryItem**](DictionaryItem.md) | Tier | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | Creation date | [optional] [default to null]
**CreationUser** | [***UserResource**](UserResource.md) | User who created the disk | [optional] [default to null]
**IsShared** | **bool** | If disk is shared | [optional] [default to null]
**SharedDiskType** | [***DictionaryItem**](DictionaryItem.md) | Shared disk type, null if disk is not shared | [optional] [default to null]
**Subregion** | [***BaseResource**](BaseResource.md) | Subregion | [optional] [default to null]
**IsLocked** | **bool** | If the disk is locked by a running operation | [optional] [default to null]
**LockingDate** | [**time.Time**](time.Time.md) | Locking date | [optional] [default to null]
**Connections** | [**[]DiskConnection**](DiskConnection.md) | Connections to instances | [optional] [default to null]
**IsFreemium** | **bool** | Is freemium | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


