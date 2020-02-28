# HealthCheckSip

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SipUserName** | **string** | Sip user name | [optional] [default to null]
**SipPassword** | **string** | Sip password | [optional] [default to null]
**SipDomain** | **string** | Sip domain | [optional] [default to null]
**Timeout** | **int32** | Timeout | [optional] [default to null]
**NotificationTypes** | [**[]DictionaryItem**](DictionaryItem.md) | Notification types enabled for a health check | [optional] [default to null]
**NotificationEventTypes** | [**[]DictionaryItem**](DictionaryItem.md) | Event types with enabled notification | [optional] [default to null]
**NotificationTime** | [***DictionaryItem**](DictionaryItem.md) | Time when notification is sent | [optional] [default to null]
**LocationsFailoverEnabled** | **bool** | Use random substitute locations in case of location breakdown | [optional] [default to null]
**Id** | **int32** | Id | [optional] [default to null]
**Interval** | **int32** | Interval | [optional] [default to null]
**Name** | **string** | Name | [optional] [default to null]
**Address** | **string** | Address | [optional] [default to null]
**ServiceType** | [***DictionaryItem**](DictionaryItem.md) | Type | [optional] [default to null]
**State** | [***DictionaryItem**](DictionaryItem.md) | State | [optional] [default to null]
**DetailsLocation** | **string** | Details url | [optional] [default to null]
**Paused** | **bool** | Is paused | [optional] [default to null]
**Suspended** | **bool** | Is suspended | [optional] [default to null]
**LastInvalidCheck** | [**time.Time**](time.Time.md) | Last invalid check | [optional] [default to null]
**LastValidCheck** | [**time.Time**](time.Time.md) | Last valid check | [optional] [default to null]
**Description** | **string** | Description | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


