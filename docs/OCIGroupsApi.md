# odk\OCIGroupsApi

All URIs are relative to *https://api.oktawave.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsChangeAssignmentsInGroup**](OCIGroupsApi.md#GroupsChangeAssignmentsInGroup) | **Put** /groups/{id}/assignments | Changes group assignments
[**GroupsCreate**](OCIGroupsApi.md#GroupsCreate) | **Post** /groups | Creates group
[**GroupsCreateContainerScheduler**](OCIGroupsApi.md#GroupsCreateContainerScheduler) | **Post** /groups/{id}/schedulers | Creates a group scheduler
[**GroupsDelete**](OCIGroupsApi.md#GroupsDelete) | **Delete** /groups/{id} | Deletes group
[**GroupsDeleteGroupScheduler**](OCIGroupsApi.md#GroupsDeleteGroupScheduler) | **Delete** /groups/schedulers/{id} | Deletes group scheduler
[**GroupsGetAssignmentsInGroup**](OCIGroupsApi.md#GroupsGetAssignmentsInGroup) | **Get** /groups/{id}/assignments | Returns group assignments
[**GroupsGetGroup**](OCIGroupsApi.md#GroupsGetGroup) | **Get** /groups/{id} | Returns group
[**GroupsGetGroupAutoscaler**](OCIGroupsApi.md#GroupsGetGroupAutoscaler) | **Get** /groups/{id}/autoscaler | Returns group autoscaler settings
[**GroupsGetGroupScheduler**](OCIGroupsApi.md#GroupsGetGroupScheduler) | **Get** /groups/schedulers/{id} | Returns group scheduler
[**GroupsGetGroupSchedulers**](OCIGroupsApi.md#GroupsGetGroupSchedulers) | **Get** /groups/{id}/schedulers | Returns group schedulers
[**GroupsGetGroups**](OCIGroupsApi.md#GroupsGetGroups) | **Get** /groups | Returns a list of groups
[**GroupsGetLoadBalancers**](OCIGroupsApi.md#GroupsGetLoadBalancers) | **Get** /groups/load_balancers | Gets load balancers
[**GroupsSetGroupAutoscaler**](OCIGroupsApi.md#GroupsSetGroupAutoscaler) | **Post** /groups/{id}/autoscaler | Sets group autoscaler
[**GroupsTurnoffGroupAutoscaler**](OCIGroupsApi.md#GroupsTurnoffGroupAutoscaler) | **Delete** /groups/{id}/autoscaler | Turns off group autoscaler
[**GroupsUpdate**](OCIGroupsApi.md#GroupsUpdate) | **Put** /groups/{id} | Updates group
[**GroupsUpdateGroupScheduler**](OCIGroupsApi.md#GroupsUpdateGroupScheduler) | **Put** /groups/schedulers/{id} | Updates a group scheduler
[**InstancesGetGroups**](OCIGroupsApi.md#InstancesGetGroups) | **Get** /instances/{id}/groups | Returns a list of instance groups
[**LoadBalancersChangeServiceStatus**](OCIGroupsApi.md#LoadBalancersChangeServiceStatus) | **Put** /groups/{id}/load_balancer/services | Changes load balancer service state
[**LoadBalancersCreate**](OCIGroupsApi.md#LoadBalancersCreate) | **Post** /groups/{id}/load_balancer | Create load balancer for group
[**LoadBalancersDelete**](OCIGroupsApi.md#LoadBalancersDelete) | **Delete** /groups/{id}/load_balancer | Delete load balancer
[**LoadBalancersGetLoadBalancer**](OCIGroupsApi.md#LoadBalancersGetLoadBalancer) | **Get** /groups/{id}/load_balancer | Gets load balancer for group
[**LoadBalancersGetLoadBalancerDetails**](OCIGroupsApi.md#LoadBalancersGetLoadBalancerDetails) | **Get** /groups/{id}/load_balancer/details | Gets load balancer detail for group
[**LoadBalancersUpdate**](OCIGroupsApi.md#LoadBalancersUpdate) | **Put** /groups/{id}/load_balancer | Update load balancer for group


# **GroupsChangeAssignmentsInGroup**
> ApiCollectionGroupAssignment GroupsChangeAssignmentsInGroup(ctx, id, command)
Changes group assignments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a group | 
  **command** | [**ChangeContainerAssignmentsCommand**](ChangeContainerAssignmentsCommand.md)| Change group assignments command | 

### Return type

[**ApiCollectionGroupAssignment**](ApiCollection[GroupAssignment].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsCreate**
> Group GroupsCreate(ctx, command)
Creates group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **command** | [**CreateGroupCommand**](CreateGroupCommand.md)| Create group command | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsCreateContainerScheduler**
> GroupScheduler GroupsCreateContainerScheduler(ctx, id, command)
Creates a group scheduler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Group Id | 
  **command** | [**CreateUpdateGroupSchedulerCommand**](CreateUpdateGroupSchedulerCommand.md)| Create group scheduler command | 

### Return type

[**GroupScheduler**](GroupScheduler.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsDelete**
> Object GroupsDelete(ctx, id)
Deletes group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a group | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsDeleteGroupScheduler**
> Object GroupsDeleteGroupScheduler(ctx, id)
Deletes group scheduler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a group scheduler | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGetAssignmentsInGroup**
> ApiCollectionGroupAssignment GroupsGetAssignmentsInGroup(ctx, id, optional)
Returns group assignments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a group | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a group | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionGroupAssignment**](ApiCollection[GroupAssignment].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGetGroup**
> Group GroupsGetGroup(ctx, id, optional)
Returns group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of group | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of group | 
 **fields** | **string**| Response fields filter | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGetGroupAutoscaler**
> GroupAutoscaler GroupsGetGroupAutoscaler(ctx, id, optional)
Returns group autoscaler settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a group | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a group | 
 **fields** | **string**| Response fields filter | 

### Return type

[**GroupAutoscaler**](GroupAutoscaler.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGetGroupScheduler**
> GroupScheduler GroupsGetGroupScheduler(ctx, id, optional)
Returns group scheduler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Group scheduler Id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Group scheduler Id | 
 **fields** | **string**| Response fields filter | 

### Return type

[**GroupScheduler**](GroupScheduler.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGetGroupSchedulers**
> ApiCollectionGroupScheduler GroupsGetGroupSchedulers(ctx, id, optional)
Returns group schedulers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a group | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Id of a group | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionGroupScheduler**](ApiCollection[GroupScheduler].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGetGroups**
> ApiCollectionGroup GroupsGetGroups(ctx, optional)
Returns a list of groups

Acceptable order values are: Name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionGroup**](ApiCollection[Group].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsGetLoadBalancers**
> ApiCollectionLoadBalancer GroupsGetLoadBalancers(ctx, optional)
Gets load balancers

Acceptable order values are: GroupName

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32**| Page size | 
 **pageNumber** | **int32**| Page number | 
 **orderBy** | **string**| Order by | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionLoadBalancer**](ApiCollection[LoadBalancer].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsSetGroupAutoscaler**
> GroupAutoscaler GroupsSetGroupAutoscaler(ctx, id, command)
Sets group autoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a group | 
  **command** | [**SetGroupAutoscalerCommand**](SetGroupAutoscalerCommand.md)| Set group autoscaler command | 

### Return type

[**GroupAutoscaler**](GroupAutoscaler.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsTurnoffGroupAutoscaler**
> Object GroupsTurnoffGroupAutoscaler(ctx, id)
Turns off group autoscaler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a group | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsUpdate**
> Group GroupsUpdate(ctx, id, command)
Updates group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a group | 
  **command** | [**CreateGroupCommand**](CreateGroupCommand.md)| Update group command | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsUpdateGroupScheduler**
> GroupScheduler GroupsUpdateGroupScheduler(ctx, id, command)
Updates a group scheduler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Group scheduler Id | 
  **command** | [**CreateUpdateGroupSchedulerCommand**](CreateUpdateGroupSchedulerCommand.md)| Create group scheduler command | 

### Return type

[**GroupScheduler**](GroupScheduler.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstancesGetGroups**
> ApiCollectionGroup InstancesGetGroups(ctx, id, optional)
Returns a list of instance groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Instance id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Instance id | 
 **fields** | **string**| Response fields filter | 

### Return type

[**ApiCollectionGroup**](ApiCollection[Group].md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadBalancersChangeServiceStatus**
> LoadBalancer LoadBalancersChangeServiceStatus(ctx, id, command)
Changes load balancer service state

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Id of a group | 
  **command** | [**ChangeContainerServiceStateCommand**](ChangeContainerServiceStateCommand.md)| Change service status command | 

### Return type

[**LoadBalancer**](LoadBalancer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadBalancersCreate**
> LoadBalancer LoadBalancersCreate(ctx, id, command)
Create load balancer for group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Group id | 
  **command** | [**SetLoadBalancerCommand**](SetLoadBalancerCommand.md)| Create load balancer command | 

### Return type

[**LoadBalancer**](LoadBalancer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadBalancersDelete**
> Object LoadBalancersDelete(ctx, id)
Delete load balancer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Group id | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadBalancersGetLoadBalancer**
> LoadBalancer LoadBalancersGetLoadBalancer(ctx, id, optional)
Gets load balancer for group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Group id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Group id | 
 **fields** | **string**| Response fields filter | 

### Return type

[**LoadBalancer**](LoadBalancer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadBalancersGetLoadBalancerDetails**
> LoadBalancer LoadBalancersGetLoadBalancerDetails(ctx, id, optional)
Gets load balancer detail for group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Group id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Group id | 
 **fields** | **string**| Response fields filter | 

### Return type

[**LoadBalancer**](LoadBalancer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadBalancersUpdate**
> LoadBalancer LoadBalancersUpdate(ctx, id, command)
Update load balancer for group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Group id | 
  **command** | [**SetLoadBalancerCommand**](SetLoadBalancerCommand.md)| Update load balancer command | 

### Return type

[**LoadBalancer**](LoadBalancer.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/xml, text/xml, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

