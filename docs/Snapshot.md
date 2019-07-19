# Snapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID | [optional] [default to null]
**Name** | **string** | Name | [optional] [default to null]
**Description** | **string** | Description | [optional] [default to null]
**Instance** | [***BaseResource**](BaseResource.md) | Instance | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | Creation date | [optional] [default to null]
**LastChangeDate** | [**time.Time**](time.Time.md) | Last change date | [optional] [default to null]
**CreationUser** | [***UserResource**](UserResource.md) | User who created the snapshot | [optional] [default to null]
**IsSystem** | **bool** | Is system snapshot. Not created by the user. | [optional] [default to null]
**IsCurrent** | **bool** | Is current | [optional] [default to null]
**SnapshotParent** | [***BaseResource**](BaseResource.md) | Snapshot parent | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


