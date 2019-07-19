# CreateUpdateHealthCheckFullPageCommand

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpMethodId** | **int32** | Health check http method type (Dictionary 166) | [default to 1440]
**ContentRegularExpression** | **string** | The content has to match the expression (GET and POST methods only) | [optional] [default to null]
**ContentNegativeRegularExpression** | **string** | The content cannot match the expression (GET and POST methods only) | [optional] [default to null]
**Port** | **int32** | Port | [default to 80]
**PageTimeout** | **int32** | Time limit for the main page body [ms] | [default to 7000]
**ElementTimeout** | **int32** | Time limit for each page elements [ms] | [default to 5000]
**ElementsTotalTimeout** | **int32** | Time limit for all page elements [ms] | [default to 10000]
**FetchPageElements** | **bool** | Fetch page elements | [default to null]
**MaxRedirects** | **int32** | Maximum length of HTTP redirects sequence | [default to 5]
**MaxParallelRequests** | **int32** | Maximum number of HTTP requests run in parallel | [default to 6]
**GenerateHar** | **bool** | Generate a HAR file for each check | [default to null]
**AllowedElementErrorCount** | **int32** | Number of elements that may not be fetched correctly | [optional] [default to 0]
**ContentSizeLimit** | **int32** | Content size limit (bytes) | [optional] [default to 2097152]
**IgnoreHtmlParsingTime** | **bool** | Ignore HTML code processing time in results | [default to null]
**SaveCookies** | **bool** | Save cookies between checks | [default to null]
**DisableContentEncoding** | **bool** | Disable HTTP transfer compression | [default to null]
**Content** | **string** | Content | [optional] [default to null]
**ContentType** | **string** | Content type | [optional] [default to null]
**IgnoredElementsFilter** | **string** | Ignore errors for elements with URLs matching the expression (only if page elements are fetched) | [optional] [default to null]
**ElementsFilter** | **string** | Do not fetch elements with URLs that match the expression | [optional] [default to null]
**ErrorTolerance** | **int32** | How many (%) locations have to report an error to consider it a breakdown | [default to 51]
**Name** | **string** | Health check name | [default to null]
**Address** | **string** | URL or IP address of the monitored service | [default to null]
**Interval** | **int32** | Time interval between health checks, in seconds | [default to 60]
**Paused** | **bool** | Is paused | [default to null]
**LocationsFailoverEnabled** | **bool** | Use random substitute locations in case of location breakdown | [default to null]
**NotificationTypeIds** | **[]int32** | Notification types enabled for a health check | [optional] [default to null]
**NotificationEventTypeIds** | **[]int32** | Event types with enabled notification | [optional] [default to null]
**NotificationTimeId** | **int32** | Time when notification is sent | [default to 1594]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


