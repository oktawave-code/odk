# InstanceAutoscalingParameter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoscalingParameterType** | [***DictionaryItem**](DictionaryItem.md) | Autoscaling parameter type | [optional] [default to null]
**MinRam** | **int32** | Minimum RAM capacity in megabytes | [optional] [default to null]
**MaxRam** | **int32** | Maximum RAM capacity in megabytes | [optional] [default to null]
**MinCpu** | **int32** | Minimum CPU count | [optional] [default to null]
**MaxCpu** | **int32** | Maximum CPU count | [optional] [default to null]
**HasAgreedToIncreaseClass** | **bool** | Instance class increase agreement | [optional] [default to null]
**HasAgreedToDecreaseClass** | **bool** | Instance class decrease agreement | [optional] [default to null]
**HasAgreedToRestart** | **bool** | Restart agreement | [optional] [default to null]
**RestartTimeFrom** | [***NullableTime**](Nullable[Time].md) | Restart time lower limit | [optional] [default to null]
**RestartTimeTo** | [***NullableTime**](Nullable[Time].md) | Restart time upper limit | [optional] [default to null]
**TimeZoneName** | **string** | Time zone name (https://www.iana.org/time-zones) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


