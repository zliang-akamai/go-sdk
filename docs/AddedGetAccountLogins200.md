# AddedGetAccountLogins200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetAccountLogins200ResponseDataInner**](GetAccountLogins200ResponseDataInner.md) |  | [optional] 
**Page** | Pointer to **int32** | The current [page](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Pages** | Pointer to **int32** | The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination). | [optional] [readonly] 
**Results** | Pointer to **int32** | The total number of results. | [optional] [readonly] 

## Methods

### NewAddedGetAccountLogins200

`func NewAddedGetAccountLogins200() *AddedGetAccountLogins200`

NewAddedGetAccountLogins200 instantiates a new AddedGetAccountLogins200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedGetAccountLogins200WithDefaults

`func NewAddedGetAccountLogins200WithDefaults() *AddedGetAccountLogins200`

NewAddedGetAccountLogins200WithDefaults instantiates a new AddedGetAccountLogins200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AddedGetAccountLogins200) GetData() []GetAccountLogins200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddedGetAccountLogins200) GetDataOk() (*[]GetAccountLogins200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddedGetAccountLogins200) SetData(v []GetAccountLogins200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AddedGetAccountLogins200) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPage

`func (o *AddedGetAccountLogins200) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AddedGetAccountLogins200) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AddedGetAccountLogins200) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *AddedGetAccountLogins200) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPages

`func (o *AddedGetAccountLogins200) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *AddedGetAccountLogins200) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *AddedGetAccountLogins200) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *AddedGetAccountLogins200) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetResults

`func (o *AddedGetAccountLogins200) GetResults() int32`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AddedGetAccountLogins200) GetResultsOk() (*int32, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AddedGetAccountLogins200) SetResults(v int32)`

SetResults sets Results field to given value.

### HasResults

`func (o *AddedGetAccountLogins200) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


