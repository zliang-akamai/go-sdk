/*
Linode API

[Read the API documentation](https://techdocs.akamai.com/linode-api/reference/api).

API version: 4.189.2
Contact: support@linode.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// TagsAPIService TagsAPI service
type TagsAPIService service

type ApiDeleteAgRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	apiVersion string
	tagLabel string
}

func (r ApiDeleteAgRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteAgExecute(r)
}

/*
DeleteAg Delete a tag

Remove a Tag from all objects and delete it.

__Important__. You must be an unrestricted User in order to access, add, or modify Tags information.


<<LB>>

---


- __CLI__.

    ```
    linode-cli tags delete
linode-cli tags rm
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    account:read_write
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @param tagLabel
 @return ApiDeleteAgRequest
*/
func (a *TagsAPIService) DeleteAg(ctx context.Context, apiVersion string, tagLabel string) ApiDeleteAgRequest {
	return ApiDeleteAgRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
		tagLabel: tagLabel,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *TagsAPIService) DeleteAgExecute(r ApiDeleteAgRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.DeleteAg")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/tags/{tagLabel}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tagLabel"+"}", url.PathEscape(parameterValueToString(r.tagLabel, "tagLabel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GetAccountDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTaggedObjectsRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	apiVersion string
	tagLabel string
	page *int32
	pageSize *int32
}

// The page of a collection to return.
func (r ApiGetTaggedObjectsRequest) Page(page int32) ApiGetTaggedObjectsRequest {
	r.page = &page
	return r
}

// The number of items to return per page.
func (r ApiGetTaggedObjectsRequest) PageSize(pageSize int32) ApiGetTaggedObjectsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetTaggedObjectsRequest) Execute() (*GetTaggedObjects200Response, *http.Response, error) {
	return r.ApiService.GetTaggedObjectsExecute(r)
}

/*
GetTaggedObjects List tagged objects

Returns a paginated list of all objects you've tagged with the requested Tag. This is a mixed collection of all object types.

__Important__. You must be an unrestricted User in order to access, add, or modify Tags information.


<<LB>>

---


- __OAuth scopes__.

    ```
    account:read_only
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @param tagLabel
 @return ApiGetTaggedObjectsRequest
*/
func (a *TagsAPIService) GetTaggedObjects(ctx context.Context, apiVersion string, tagLabel string) ApiGetTaggedObjectsRequest {
	return ApiGetTaggedObjectsRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
		tagLabel: tagLabel,
	}
}

// Execute executes the request
//  @return GetTaggedObjects200Response
func (a *TagsAPIService) GetTaggedObjectsExecute(r ApiGetTaggedObjectsRequest) (*GetTaggedObjects200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTaggedObjects200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.GetTaggedObjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/tags/{tagLabel}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tagLabel"+"}", url.PathEscape(parameterValueToString(r.tagLabel, "tagLabel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 100
		r.pageSize = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GetAccountDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTagsRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	apiVersion string
	page *int32
	pageSize *int32
}

// The page of a collection to return.
func (r ApiGetTagsRequest) Page(page int32) ApiGetTagsRequest {
	r.page = &page
	return r
}

// The number of items to return per page.
func (r ApiGetTagsRequest) PageSize(pageSize int32) ApiGetTagsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetTagsRequest) Execute() (*GetTags200Response, *http.Response, error) {
	return r.ApiService.GetTagsExecute(r)
}

/*
GetTags List tags

Tags are User-defined labels attached to objects in your Account, such as Linodes. They are used for specifying and grouping attributes of objects that are relevant to the User.

This operation returns a paginated list of Tags on your account.

__Important__. You must be an unrestricted User in order to access, add, or modify Tags information.


<<LB>>

---


- __CLI__.

    ```
    linode-cli tags list
linode-cli tags ls
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    account:read_only
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @return ApiGetTagsRequest
*/
func (a *TagsAPIService) GetTags(ctx context.Context, apiVersion string) ApiGetTagsRequest {
	return ApiGetTagsRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
	}
}

// Execute executes the request
//  @return GetTags200Response
func (a *TagsAPIService) GetTagsExecute(r ApiGetTagsRequest) (*GetTags200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTags200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.GetTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 100
		r.pageSize = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GetAccountDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostTagRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	apiVersion string
	postTagRequest *PostTagRequest
}

// The tag to create, and optionally the objects to tag.
func (r ApiPostTagRequest) PostTagRequest(postTagRequest PostTagRequest) ApiPostTagRequest {
	r.postTagRequest = &postTagRequest
	return r
}

func (r ApiPostTagRequest) Execute() (*GetTags200ResponseDataInner, *http.Response, error) {
	return r.ApiService.PostTagExecute(r)
}

/*
PostTag Create a tag

Creates a new Tag and optionally tags requested objects with it immediately.

__Important__. You must be an unrestricted User in order to access, add, or modify Tags information.


<<LB>>

---


- __CLI__.

    ```
    linode-cli tags create \
  --label 'example tag' \
  --linodes 123 \
  --linodes 456 \
  --volumes 9082 \
  --volumes 10003
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    account:read_write
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @return ApiPostTagRequest
*/
func (a *TagsAPIService) PostTag(ctx context.Context, apiVersion string) ApiPostTagRequest {
	return ApiPostTagRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
	}
}

// Execute executes the request
//  @return GetTags200ResponseDataInner
func (a *TagsAPIService) PostTagExecute(r ApiPostTagRequest) (*GetTags200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTags200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.PostTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.postTagRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GetAccountDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
