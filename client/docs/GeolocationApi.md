# {{classname}}

All URIs are relative to *http://127.0.0.1:1323/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGeoLocationFromIp**](GeolocationApi.md#GetGeoLocationFromIp) | **Get** /geolocate/{ip} | Gets the location data of an IPv4 or Ipv6.

# **GetGeoLocationFromIp**
> map[string]interface{} GetGeoLocationFromIp(ctx, ip)
Gets the location data of an IPv4 or Ipv6.

Requests an IP address and returns the location data using the ipgeolocation.io API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ip** | **string**| IP Address | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

