# Instance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID | [optional] [default to null]
**Name** | **string** | Name | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | Creation date | [optional] [default to null]
**CreationUser** | [***UserResource**](UserResource.md) | User who created the instance | [optional] [default to null]
**IsLocked** | **bool** | If the instance is locked by a running operation | [optional] [default to null]
**LockingDate** | [**time.Time**](time.Time.md) | Locking date | [optional] [default to null]
**Template** | [***BaseResource**](BaseResource.md) | Template from which the instance was created | [optional] [default to null]
**Subregion** | [***BaseResource**](BaseResource.md) | Subregion that is running the instance | [optional] [default to null]
**Type_** | [***DictionaryItem**](DictionaryItem.md) | Instance type. Defines the configuration of CPU and RAM | [optional] [default to null]
**Status** | [***DictionaryItem**](DictionaryItem.md) | Status | [optional] [default to null]
**SystemCategory** | [***DictionaryItem**](DictionaryItem.md) | Operating system category installed on the instance | [optional] [default to null]
**AutoscalingType** | [***DictionaryItem**](DictionaryItem.md) | Autoscaling type | [optional] [default to null]
**VmWareToolsStatus** | [***DictionaryItem**](DictionaryItem.md) | VMware Tools status | [optional] [default to null]
**MonitStatus** | [***DictionaryItem**](DictionaryItem.md) | Monitoring status | [optional] [default to null]
**TemplateType** | [***DictionaryItem**](DictionaryItem.md) | Template type eg marketplace, oci instance | [optional] [default to null]
**IpAddress** | **string** | IP address | [optional] [default to null]
**PrivateIpAddress** | **string** | Private IP address | [optional] [default to null]
**DnsAddress** | **string** | DNS address | [optional] [default to null]
**TotalDisksCapacity** | **int32** | Total disks capacity in GB | [optional] [default to null]
**PaymentType** | [***DictionaryItem**](DictionaryItem.md) | Payment type | [optional] [default to null]
**HealthCheck** | [***BaseResource**](BaseResource.md) | Health check | [optional] [default to null]
**ScsiControllerType** | [***DictionaryItem**](DictionaryItem.md) | SCSI controller type | [optional] [default to null]
**IsFreemium** | **bool** | Is freemium | [optional] [default to null]
**CpuNumber** | **int32** | Number of CPUs | [optional] [default to null]
**RamMb** | **int32** | Memory in MB | [optional] [default to null]
**SupportType** | [***Software**](Software.md) | Support type | [optional] [default to null]
**IsHotPlugEnabled** | **bool** | Is hot plug enabled | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


