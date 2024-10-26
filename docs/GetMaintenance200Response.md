# GetMaintenance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetMaintenance200ResponseDataInner**](GetMaintenance200ResponseDataInner.md) |  | [optional] 
**Page** | Pointer to **int32** | The current [page](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Pages** | Pointer to **int32** | The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Results** | Pointer to **int32** | The total number of results. | [optional] [readonly] 

## Methods

### NewGetMaintenance200Response

`func NewGetMaintenance200Response() *GetMaintenance200Response`

NewGetMaintenance200Response instantiates a new GetMaintenance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMaintenance200ResponseWithDefaults

`func NewGetMaintenance200ResponseWithDefaults() *GetMaintenance200Response`

NewGetMaintenance200ResponseWithDefaults instantiates a new GetMaintenance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMaintenance200Response) GetData() []GetMaintenance200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMaintenance200Response) GetDataOk() (*[]GetMaintenance200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMaintenance200Response) SetData(v []GetMaintenance200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetMaintenance200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPage

`func (o *GetMaintenance200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetMaintenance200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetMaintenance200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetMaintenance200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *GetMaintenance200Response) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetMaintenance200Response) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetMaintenance200Response) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GetMaintenance200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetResults

`func (o *GetMaintenance200Response) GetResults() int32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetMaintenance200Response) GetResultsOk() (*int32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetMaintenance200Response) SetResults(v int32)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetMaintenance200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


