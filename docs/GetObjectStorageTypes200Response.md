# GetObjectStorageTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetObjectStorageTypes200ResponseDataInner**](GetObjectStorageTypes200ResponseDataInner.md) | The Object Storage types. | [optional] 
**Page** | Pointer to **int32** | The current [page](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Pages** | Pointer to **int32** | The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Results** | Pointer to **int32** | The total number of results. | [optional] [readonly] 

## Methods

### NewGetObjectStorageTypes200Response

`func NewGetObjectStorageTypes200Response() *GetObjectStorageTypes200Response`

NewGetObjectStorageTypes200Response instantiates a new GetObjectStorageTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageTypes200ResponseWithDefaults

`func NewGetObjectStorageTypes200ResponseWithDefaults() *GetObjectStorageTypes200Response`

NewGetObjectStorageTypes200ResponseWithDefaults instantiates a new GetObjectStorageTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetObjectStorageTypes200Response) GetData() []GetObjectStorageTypes200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetObjectStorageTypes200Response) GetDataOk() (*[]GetObjectStorageTypes200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetObjectStorageTypes200Response) SetData(v []GetObjectStorageTypes200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetObjectStorageTypes200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPage

`func (o *GetObjectStorageTypes200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetObjectStorageTypes200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetObjectStorageTypes200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetObjectStorageTypes200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *GetObjectStorageTypes200Response) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetObjectStorageTypes200Response) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetObjectStorageTypes200Response) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GetObjectStorageTypes200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetResults

`func (o *GetObjectStorageTypes200Response) GetResults() int32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetObjectStorageTypes200Response) GetResultsOk() (*int32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetObjectStorageTypes200Response) SetResults(v int32)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetObjectStorageTypes200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


