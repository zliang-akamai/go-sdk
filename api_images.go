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


// ImagesAPIService ImagesAPI service
type ImagesAPIService service

type ApiDeleteImageRequest struct {
	ctx context.Context
	ApiService *ImagesAPIService
	apiVersion string
	imageId string
}

func (r ApiDeleteImageRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteImageExecute(r)
}

/*
DeleteImage Delete an image

Deletes a private image you have permission to `read_write`.

> 🚧
>
> - You can't undo this delete action.
>
> - When you delete an image, all [replicated instances](https://techdocs.akamai.com/linode-api/reference/post-replicate-image) of that image are also deleted.


<<LB>>

---


- __CLI__.

    ```
    linode-cli images delete 12345
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    images:read_write
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @param imageId The unique identifier assigned to the image after creation.
 @return ApiDeleteImageRequest
*/
func (a *ImagesAPIService) DeleteImage(ctx context.Context, apiVersion string, imageId string) ApiDeleteImageRequest {
	return ApiDeleteImageRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
		imageId: imageId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ImagesAPIService) DeleteImageExecute(r ApiDeleteImageRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesAPIService.DeleteImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/images/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

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

type ApiGetImageRequest struct {
	ctx context.Context
	ApiService *ImagesAPIService
	apiVersion string
	imageId string
}

func (r ApiGetImageRequest) Execute() (*GetImages200ResponseDataInner, *http.Response, error) {
	return r.ApiService.GetImageExecute(r)
}

/*
GetImage Get an image

Get information about a single image. An image can be one of two types:

- **Public image**. The `id` for these images begins with `linode/`. These images are generally available to all users. To limit the response to public images, don't include [authentication](https://techdocs.akamai.com/linode-api/reference/get-started#authentication) when calling this operation.

- **Private image**. The `id` for these images begins with `private/`. These images are account-specific and only accessible to users with appropriate [grants](https://techdocs.akamai.com/linode-api/reference/get-user-grants). To view private images, you need to include authentication when calling this operation. The response will also include public images.


<<LB>>

---


- __CLI__.

    ```
    linode-cli images view linode/debian9
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    images:read_only
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @param imageId The unique identifier assigned to the image after creation.
 @return ApiGetImageRequest
*/
func (a *ImagesAPIService) GetImage(ctx context.Context, apiVersion string, imageId string) ApiGetImageRequest {
	return ApiGetImageRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
		imageId: imageId,
	}
}

// Execute executes the request
//  @return GetImages200ResponseDataInner
func (a *ImagesAPIService) GetImageExecute(r ApiGetImageRequest) (*GetImages200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetImages200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesAPIService.GetImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/images/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

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

type ApiGetImagesRequest struct {
	ctx context.Context
	ApiService *ImagesAPIService
	apiVersion string
	page *int32
	pageSize *int32
}

// The page of a collection to return.
func (r ApiGetImagesRequest) Page(page int32) ApiGetImagesRequest {
	r.page = &page
	return r
}

// The number of items to return per page.
func (r ApiGetImagesRequest) PageSize(pageSize int32) ApiGetImagesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetImagesRequest) Execute() (*GetImages200Response, *http.Response, error) {
	return r.ApiService.GetImagesExecute(r)
}

/*
GetImages List images

Returns a paginated list of images. An image can be one of two types:

- **Public image**. The `id` for these images begins with `linode/`. These images are generally available to all users. To limit the response to public images, don't include [authentication](https://techdocs.akamai.com/linode-api/reference/get-started#authentication) when calling this operation.

- **Private image**. The `id` for these images begins with `private/`. These images are account-specific and only accessible to users with appropriate [grants](https://techdocs.akamai.com/linode-api/reference/get-user-grants). To view private images, you need to include authentication when calling this operation. The response includes both private and public images.


<<LB>>

---


- __CLI__.

    ```
    linode-cli images list
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    images:read_only
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @return ApiGetImagesRequest
*/
func (a *ImagesAPIService) GetImages(ctx context.Context, apiVersion string) ApiGetImagesRequest {
	return ApiGetImagesRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
	}
}

// Execute executes the request
//  @return GetImages200Response
func (a *ImagesAPIService) GetImagesExecute(r ApiGetImagesRequest) (*GetImages200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetImages200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesAPIService.GetImages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/images"
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

type ApiPostImageRequest struct {
	ctx context.Context
	ApiService *ImagesAPIService
	apiVersion string
	postImageRequest *PostImageRequest
}

func (r ApiPostImageRequest) PostImageRequest(postImageRequest PostImageRequest) ApiPostImageRequest {
	r.postImageRequest = &postImageRequest
	return r
}

func (r ApiPostImageRequest) Execute() (*GetImages200ResponseDataInner, *http.Response, error) {
	return r.ApiService.PostImageExecute(r)
}

/*
PostImage Create an image

Captures a private, gold-master image from a Linode disk.

> 📘
>
> - Captured images are stored using our Object Storage service. The `region` where the target image exists determines where the [resulting image is stored](https://techdocs.akamai.com/cloud-computing/docs/images#regions-and-captured-custom-images).
>
> - If you create an image from an encrypted disk, the API doesn't encrypt the image. When you rebuild a compute instance from that image, the resulting disk will be automatically encrypted.


<<LB>>

---


- __CLI__.

    ```
    linode-cli images create \
  --label this_is_a_label \
  --description "A longer description \
    of the image" \
  --disk_id 123
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    images:read_write
linodes:read_only
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @return ApiPostImageRequest
*/
func (a *ImagesAPIService) PostImage(ctx context.Context, apiVersion string) ApiPostImageRequest {
	return ApiPostImageRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
	}
}

// Execute executes the request
//  @return GetImages200ResponseDataInner
func (a *ImagesAPIService) PostImageExecute(r ApiPostImageRequest) (*GetImages200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetImages200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesAPIService.PostImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/images"
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
	localVarPostBody = r.postImageRequest
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

type ApiPostReplicateImageRequest struct {
	ctx context.Context
	ApiService *ImagesAPIService
	apiVersion string
	imageId string
	postReplicateImageRequest *PostReplicateImageRequest
}

func (r ApiPostReplicateImageRequest) PostReplicateImageRequest(postReplicateImageRequest PostReplicateImageRequest) ApiPostReplicateImageRequest {
	r.postReplicateImageRequest = &postReplicateImageRequest
	return r
}

func (r ApiPostReplicateImageRequest) Execute() (*GetImages200ResponseDataInner, *http.Response, error) {
	return r.ApiService.PostReplicateImageExecute(r)
}

/*
PostReplicateImage Replicate an image

__Limited availability__ Target an existing image and replicate it to another compute region.

- This operation is in Limited Availability. Talk to your account team about access to it.

- This is only available in a `region` that supports Object Storage, which stores the replicated image. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation to review a region's `capabilities`.

- To replicate an image, it needs to have a `status` of `available`. Run the [List images](https://techdocs.akamai.com/linode-api/reference/get-images) operation to view an image's `status`.

- To also keep the target image in its original compute region, you need to include that `region` in the request's data. If you leave it out, the API removes the image from that region. Run the [Get an image](https://techdocs.akamai.com/linode-api/reference/get-image) operation to see the `regions` where an image currently exists.

- You can't include an empty array to delete all images. You need to provide at least one compute `region` where the image is `available`. Use the [Delete an image](https://techdocs.akamai.com/linode-api/reference/delete-image) operation.


<<LB>>

---


- __CLI__.

    ```
    linode-cli images replicate private/12345 \
  --regions "us-mia" \
  --regions "us-east"
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    images:read_write
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @param imageId The unique identifier assigned to the image after creation.
 @return ApiPostReplicateImageRequest
*/
func (a *ImagesAPIService) PostReplicateImage(ctx context.Context, apiVersion string, imageId string) ApiPostReplicateImageRequest {
	return ApiPostReplicateImageRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
		imageId: imageId,
	}
}

// Execute executes the request
//  @return GetImages200ResponseDataInner
func (a *ImagesAPIService) PostReplicateImageExecute(r ApiPostReplicateImageRequest) (*GetImages200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetImages200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesAPIService.PostReplicateImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/images/{imageId}/regions"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postReplicateImageRequest == nil {
		return localVarReturnValue, nil, reportError("postReplicateImageRequest is required and must be specified")
	}

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
	localVarPostBody = r.postReplicateImageRequest
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

type ApiPostUploadImageRequest struct {
	ctx context.Context
	ApiService *ImagesAPIService
	apiVersion string
	postUploadImageRequest *PostUploadImageRequest
}

// The uploaded image details.
func (r ApiPostUploadImageRequest) PostUploadImageRequest(postUploadImageRequest PostUploadImageRequest) ApiPostUploadImageRequest {
	r.postUploadImageRequest = &postUploadImageRequest
	return r
}

func (r ApiPostUploadImageRequest) Execute() (*PostUploadImage200Response, *http.Response, error) {
	return r.ApiService.PostUploadImageExecute(r)
}

/*
PostUploadImage Upload an image

Creates a new private image container and returns a URL as the `upload_to` object in the response. Use this URL to upload your own disk image to the container.

1. Ensure the disk image is raw disk image (`.img`) format.

2. Compress the disk image using gzip (`.gz`) format. Compressed, the file can be up to 5 GB and decompressed it can be up to 6 GB.

3. Upload the file in a separate PUT request that includes the `Content-type: application/octet-stream` header:

  ```
  curl -v \
    -H "Content-Type: application/octet-stream" \
    --upload-file example.img.gz \
    $UPLOAD_URL \
    --progress-bar \
    --output /dev/null
  ```

> 📘
>
> - You need to upload image data within 24 hours of creation or the API cancels the upload and deletes the image container.
>
> - Only core regions that support our [Object Storage](https://techdocs.akamai.com/cloud-computing/reference/how-to-choose-a-data-center#product-availability) service can store an uploaded image.
>
> - If you create an image from an encrypted disk, the API doesn't encrypt the image. When you rebuild a compute instance from that image, the resulting disk will be automatically encrypted.
>
> - You can create a new image and upload image data using a single process through [Cloud Manager](https://www.linode.com/docs/products/tools/images/guides/upload-an-image/#uploading-an-image-file-through-the-cloud-manager) or the [Linode CLI](https://www.linode.com/docs/products/tools/images/guides/upload-an-image/#uploading-an-image-file-through-the-linode-cli).


<<LB>>

---


- __CLI__.

    ```
    # Run the operation to just get the upload_to URL
linode-cli images upload \
  --description "Optional details about the Image" \
  --label "Example Image" \
  --region us-east

# Upload the image file in a single step
linode-cli image-upload \
  --description "Optional details about the Image" \
  --label "Example Image" \
  --region us-east \
  /path/to/image-file.img.gz
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    images:read_write
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @return ApiPostUploadImageRequest
*/
func (a *ImagesAPIService) PostUploadImage(ctx context.Context, apiVersion string) ApiPostUploadImageRequest {
	return ApiPostUploadImageRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
	}
}

// Execute executes the request
//  @return PostUploadImage200Response
func (a *ImagesAPIService) PostUploadImageExecute(r ApiPostUploadImageRequest) (*PostUploadImage200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostUploadImage200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesAPIService.PostUploadImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/images/upload"
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
	localVarPostBody = r.postUploadImageRequest
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

type ApiPutImageRequest struct {
	ctx context.Context
	ApiService *ImagesAPIService
	apiVersion string
	imageId string
	getImages200ResponseDataInner *GetImages200ResponseDataInner
}

func (r ApiPutImageRequest) GetImages200ResponseDataInner(getImages200ResponseDataInner GetImages200ResponseDataInner) ApiPutImageRequest {
	r.getImages200ResponseDataInner = &getImages200ResponseDataInner
	return r
}

func (r ApiPutImageRequest) Execute() (*GetImages200ResponseDataInner, *http.Response, error) {
	return r.ApiService.PutImageExecute(r)
}

/*
PutImage Update an image

Updates a private image that you have permission to `read_write`.

> 📘
>
> You can't update the `regions` with this operation. Use the [Replicate an image](https://techdocs.akamai.com/linode-api/reference/post-replicate-image) operation to modify the existing regions for your image.


<<LB>>

---


- __CLI__.

    ```
    linode-cli images update private/12345 \
  --label "My gold-master image" \
  --description "The detailed description \
    of my image."
    ```

    [Learn more...](https://www.linode.com/docs/products/tools/cli/get-started/)

- __OAuth scopes__.

    ```
    images:read_write
    ```

    [Learn more...](https://techdocs.akamai.com/linode-api/reference/get-started#oauth)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param apiVersion __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta.
 @param imageId The unique identifier assigned to the image after creation.
 @return ApiPutImageRequest
*/
func (a *ImagesAPIService) PutImage(ctx context.Context, apiVersion string, imageId string) ApiPutImageRequest {
	return ApiPutImageRequest{
		ApiService: a,
		ctx: ctx,
		apiVersion: apiVersion,
		imageId: imageId,
	}
}

// Execute executes the request
//  @return GetImages200ResponseDataInner
func (a *ImagesAPIService) PutImageExecute(r ApiPutImageRequest) (*GetImages200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetImages200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImagesAPIService.PutImage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{apiVersion}/images/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", url.PathEscape(parameterValueToString(r.apiVersion, "apiVersion")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", url.PathEscape(parameterValueToString(r.imageId, "imageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getImages200ResponseDataInner == nil {
		return localVarReturnValue, nil, reportError("getImages200ResponseDataInner is required and must be specified")
	}

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
	localVarPostBody = r.getImages200ResponseDataInner
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
