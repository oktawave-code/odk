# Template

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id | [optional] [default to null]
**Name** | **string** | Name | [optional] [default to null]
**Description** | **string** | Description | [optional] [default to null]
**Version** | **string** | Version | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | Creation date | [optional] [default to null]
**LastChangeDate** | [**time.Time**](time.Time.md) | Last change date | [optional] [default to null]
**CreationUser** | [***UserResource**](UserResource.md) | User who created the template | [optional] [default to null]
**DefaultInstanceType** | [***DictionaryItem**](DictionaryItem.md) | Default instance type | [optional] [default to null]
**MinimumInstanceType** | [***DictionaryItem**](DictionaryItem.md) | Minimum instance type | [optional] [default to null]
**EthernetControllersNumber** | **int32** | Ethernet controllers number | [optional] [default to null]
**EthernetControllersType** | [***DictionaryItem**](DictionaryItem.md) | Ethernet controllers type | [optional] [default to null]
**SystemCategory** | [***DictionaryItem**](DictionaryItem.md) | System category | [optional] [default to null]
**OwnerAccount** | [***BaseResource**](BaseResource.md) | Owner account | [optional] [default to null]
**PublicationStatus** | [***DictionaryItem**](DictionaryItem.md) | Publication status | [optional] [default to null]
**Disks** | [**[]TemplateDisk**](TemplateDisk.md) | Disks | [optional] [default to null]
**Software** | [**[]Software**](Software.md) | Softwares | [optional] [default to null]
**TemplateType** | [***DictionaryItem**](DictionaryItem.md) | Template type | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


