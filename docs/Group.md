# Group

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id | [optional] [default to null]
**Name** | **string** | Name | [optional] [default to null]
**IsLoadBalancer** | **bool** | Is load balancing enabled | [optional] [default to null]
**InstancesCount** | **int32** | Instances count in container | [optional] [default to null]
**SchedulersCount** | **int32** | Schedulers count assigned to container | [optional] [default to null]
**AffinityRuleType** | [***DictionaryItem**](DictionaryItem.md) | Instances affinity rule type | [optional] [default to null]
**AutoscalingType** | [***DictionaryItem**](DictionaryItem.md) | Autoscaling type | [optional] [default to null]
**LastChangeDate** | [**time.Time**](time.Time.md) | Last modified date | [optional] [default to null]
**CreationUser** | [***UserResource**](UserResource.md) | User who created the group | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


