# GetDomainRecords200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this Domain Record was created. | [optional] [readonly] 
**Id** | Pointer to **int32** | This Record&#39;s unique ID. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of this Record. For requests, this property&#39;s actual usage and whether it is required depends on the type of record this represents:  &#x60;A&#x60; and &#x60;AAAA&#x60;: The hostname or FQDN of the Record.  &#x60;NS&#x60;: The subdomain, if any, to use with the Domain of the Record. Wildcard NS records (&#x60;*&#x60;) are not supported.  &#x60;MX&#x60;: The mail subdomain. For example, &#x60;sub&#x60; for the address &#x60;user@sub.example.com&#x60; under the &#x60;example.com&#x60; Domain.  - The left-most subdomain component may be an asterisk (&#x60;*&#x60;) to designate a wildcard subdomain. - Other subdomain components must only contain letters, digits, and hyphens, start with a letter, end with a letter or digit, and contain less than 64 characters. - Must be an empty string (&#x60;\&quot;\&quot;&#x60;) for a Null MX Record.  &#x60;CNAME&#x60;: The hostname. Must be unique. Required.  &#x60;TXT&#x60;: The hostname.  &#x60;SRV&#x60;: Unused. Use the &#x60;service&#x60; property to set the service name for this record.  &#x60;CAA&#x60;: The subdomain. Omit or enter an empty string (&#x60;\&quot;\&quot;&#x60;) to apply to the entire Domain.  &#x60;PTR&#x60;: See our guide on how to [Configure Your Linode for Reverse DNS (rDNS)](https://www.linode.com/docs/guides/configure-rdns/). | [optional] 
**Port** | Pointer to **int32** | The port this Record points to. Only valid and required for SRV record requests. | [optional] 
**Priority** | Pointer to **int32** | The priority of the target host for this Record. Lower values are preferred. Only valid for MX and SRV record requests. Required for SRV record requests.  Defaults to &#x60;0&#x60; for MX record requests. Must be &#x60;0&#x60; for Null MX records. | [optional] 
**Protocol** | Pointer to **NullableString** | The protocol this Record&#39;s service communicates with. An underscore (&#x60;_&#x60;) is prepended automatically to the submitted value for this property. Only valid for SRV record requests. | [optional] 
**Service** | Pointer to **NullableString** | The name of the service. An underscore (&#x60;_&#x60;) is prepended and a period (&#x60;.&#x60;) is appended automatically to the submitted value for this property. Only valid and required for SRV record requests. | [optional] 
**Tag** | Pointer to **NullableString** | The tag portion of a CAA record. Only valid and required for CAA record requests. | [optional] 
**Target** | Pointer to **string** | The target for this Record. For requests, this property&#39;s actual usage and whether it is required depends on the type of record this represents:  &#x60;A&#x60; and &#x60;AAAA&#x60;: The IP address. Use &#x60;[remote_addr]&#x60; to submit the IPv4 address of the request. Required.  &#x60;NS&#x60;: The name server. Must be a valid domain. Required.  &#x60;MX&#x60;: The mail server. Must be a valid domain unless creating a Null MX Record. Required.  - Must have less than 254 total characters. - The left-most domain component may be an asterisk (&#x60;*&#x60;) to designate a wildcard domain. - Other domain components must only contain letters, digits, and hyphens, start with a letter, end with a letter or digit, and contain less than 64 characters. - To create a [Null MX Record](https://datatracker.ietf.org/doc/html/rfc7505), first [remove](https://techdocs.akamai.com/linode-api/reference/delete-domain-record) any additional MX records, then create an MX record with empty strings (&#x60;\&quot;\&quot;&#x60;) for the &#x60;target&#x60; and &#x60;name&#x60;. If a Domain has a Null MX record, new MX records cannot be created.  &#x60;CNAME&#x60;: The alias. Must be a valid domain. Required.  &#x60;TXT&#x60;: The value. Required.  &#x60;SRV&#x60;: The target domain or subdomain. If a subdomain is entered, it is automatically used with the Domain. To configure for a different domain, enter a valid FQDN. For example, the value &#x60;www&#x60; with a Domain for &#x60;example.com&#x60; results in a target set to &#x60;www.example.com&#x60;, whereas the value &#x60;sample.com&#x60; results in a target set to &#x60;sample.com&#x60;. Required.  &#x60;CAA&#x60;: The value. For &#x60;issue&#x60; or &#x60;issuewild&#x60; tags, the domain of your certificate issuer. For the &#x60;iodef&#x60; tag, a contact or submission URL (domain, http, https, or mailto). Requirements depend on the tag for this record:    - &#x60;issue&#x60;: The domain of your certificate issuer. Must include a valid domain. May include additional parameters separated with semicolons (&#x60;;&#x60;), for example: &#x60;www.example.com; foo&#x3D;bar&#x60;   - &#x60;issuewild&#x60;: The domain of your wildcard certificate issuer. Must be a valid domain and must not start with an asterisk (&#x60;*&#x60;).   - &#x60;iodef&#x60;: Must be either (1) a valid domain, (2) a valid domain prepended with &#x60;http://&#x60; or &#x60;https://&#x60;, or (3) a valid email address prepended with &#x60;mailto:&#x60;.  &#x60;PTR&#x60;: Required. See our guide on how to [Configure Your Linode for Reverse DNS (rDNS)](https://www.linode.com/docs/guides/configure-rdns/).  With the exception of A, AAAA, and CAA records, this field accepts a trailing period. | [optional] 
**TtlSec** | Pointer to **int32** | \&quot;Time to Live\&quot; - the amount of time in seconds that this Domain&#39;s records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value. | [optional] 
**Type** | Pointer to **string** | The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. For more information, see the guides on [DNS Record Types](https://www.linode.com/docs/products/networking/dns-manager/guides/#dns-record-types). | [optional] 
**Updated** | Pointer to **time.Time** | When this Domain Record was last updated. | [optional] [readonly] 
**Weight** | Pointer to **int32** | The relative weight of this Record used in the case of identical priority. Higher values are preferred. Only valid and required for SRV record requests. | [optional] 

## Methods

### NewGetDomainRecords200ResponseDataInner

`func NewGetDomainRecords200ResponseDataInner() *GetDomainRecords200ResponseDataInner`

NewGetDomainRecords200ResponseDataInner instantiates a new GetDomainRecords200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDomainRecords200ResponseDataInnerWithDefaults

`func NewGetDomainRecords200ResponseDataInnerWithDefaults() *GetDomainRecords200ResponseDataInner`

NewGetDomainRecords200ResponseDataInnerWithDefaults instantiates a new GetDomainRecords200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetDomainRecords200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetDomainRecords200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetDomainRecords200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetDomainRecords200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *GetDomainRecords200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDomainRecords200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDomainRecords200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetDomainRecords200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetDomainRecords200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDomainRecords200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDomainRecords200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDomainRecords200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *GetDomainRecords200ResponseDataInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetDomainRecords200ResponseDataInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetDomainRecords200ResponseDataInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetDomainRecords200ResponseDataInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPriority

`func (o *GetDomainRecords200ResponseDataInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetDomainRecords200ResponseDataInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetDomainRecords200ResponseDataInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetDomainRecords200ResponseDataInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProtocol

`func (o *GetDomainRecords200ResponseDataInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetDomainRecords200ResponseDataInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetDomainRecords200ResponseDataInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetDomainRecords200ResponseDataInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *GetDomainRecords200ResponseDataInner) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *GetDomainRecords200ResponseDataInner) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetService

`func (o *GetDomainRecords200ResponseDataInner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetDomainRecords200ResponseDataInner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetDomainRecords200ResponseDataInner) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetDomainRecords200ResponseDataInner) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *GetDomainRecords200ResponseDataInner) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *GetDomainRecords200ResponseDataInner) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetTag

`func (o *GetDomainRecords200ResponseDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetDomainRecords200ResponseDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetDomainRecords200ResponseDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetDomainRecords200ResponseDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *GetDomainRecords200ResponseDataInner) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *GetDomainRecords200ResponseDataInner) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetTarget

`func (o *GetDomainRecords200ResponseDataInner) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GetDomainRecords200ResponseDataInner) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GetDomainRecords200ResponseDataInner) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GetDomainRecords200ResponseDataInner) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTtlSec

`func (o *GetDomainRecords200ResponseDataInner) GetTtlSec() int32`

GetTtlSec returns the TtlSec field if non-nil, zero value otherwise.

### GetTtlSecOk

`func (o *GetDomainRecords200ResponseDataInner) GetTtlSecOk() (*int32, bool)`

GetTtlSecOk returns a tuple with the TtlSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSec

`func (o *GetDomainRecords200ResponseDataInner) SetTtlSec(v int32)`

SetTtlSec sets TtlSec field to given value.

### HasTtlSec

`func (o *GetDomainRecords200ResponseDataInner) HasTtlSec() bool`

HasTtlSec returns a boolean if a field has been set.

### GetType

`func (o *GetDomainRecords200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDomainRecords200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDomainRecords200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDomainRecords200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *GetDomainRecords200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetDomainRecords200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetDomainRecords200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetDomainRecords200ResponseDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetWeight

`func (o *GetDomainRecords200ResponseDataInner) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetDomainRecords200ResponseDataInner) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetDomainRecords200ResponseDataInner) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GetDomainRecords200ResponseDataInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


