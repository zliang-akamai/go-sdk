# GetSshKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetSshKeys200ResponseDataInner**](GetSshKeys200ResponseDataInner.md) |  | [optional] 
**Page** | Pointer to **int32** | The current [page](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Pages** | Pointer to **int32** | The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Results** | Pointer to **int32** | The total number of results. | [optional] [readonly] 

## Methods

### NewGetSshKeys200Response

`func NewGetSshKeys200Response() *GetSshKeys200Response`

NewGetSshKeys200Response instantiates a new GetSshKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSshKeys200ResponseWithDefaults

`func NewGetSshKeys200ResponseWithDefaults() *GetSshKeys200Response`

NewGetSshKeys200ResponseWithDefaults instantiates a new GetSshKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSshKeys200Response) GetData() []GetSshKeys200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSshKeys200Response) GetDataOk() (*[]GetSshKeys200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSshKeys200Response) SetData(v []GetSshKeys200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetSshKeys200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPage

`func (o *GetSshKeys200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetSshKeys200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetSshKeys200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetSshKeys200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *GetSshKeys200Response) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetSshKeys200Response) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetSshKeys200Response) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GetSshKeys200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetResults

`func (o *GetSshKeys200Response) GetResults() int32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetSshKeys200Response) GetResultsOk() (*int32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetSshKeys200Response) SetResults(v int32)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetSshKeys200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


