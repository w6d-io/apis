# Protocol Documentation
<a id="top"></a>

## Table of Contents

- [provider.proto](#provider.proto)
    - [CheckAccessUserInProjectRequest](#.CheckAccessUserInProjectRequest)
    - [CheckAccessUserInProjectResponse](#.CheckAccessUserInProjectResponse)
    - [DeleteWebhookRequest](#.DeleteWebhookRequest)
    - [DeleteWebhookResponse](#.DeleteWebhookResponse)
    - [ListProjectsRequest](#.ListProjectsRequest)
    - [ListProjectsResponse](#.ListProjectsResponse)
    - [Owner](#.Owner)
    - [Repository](#.Repository)
    - [SetWebhookRequest](#.SetWebhookRequest)
    - [SetWebhookResponse](#.SetWebhookResponse)
  
    - [OwnerType](#.OwnerType)
  
    - [Provider](#.Provider)
  
- [Scalar Value Types](#scalar-value-types)



<a id="provider.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## provider.proto



<a id=".CheckAccessUserInProjectRequest"></a>

### CheckAccessUserInProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| token | [string](#string) |  |  |
| userId | [string](#string) |  |  |
| project | [Repository](#Repository) |  |  |






<a id=".CheckAccessUserInProjectResponse"></a>

### CheckAccessUserInProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [bool](#bool) |  |  |
| admin | [bool](#bool) |  |  |
| push | [bool](#bool) |  |  |
| pull | [bool](#bool) |  |  |






<a id=".DeleteWebhookRequest"></a>

### DeleteWebhookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| token | [string](#string) |  |  |
| idRepo | [int32](#int32) |  |  |
| webhookId | [int32](#int32) |  |  |






<a id=".DeleteWebhookResponse"></a>

### DeleteWebhookResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [bool](#bool) |  |  |
| deleted | [bool](#bool) |  |  |






<a id=".ListProjectsRequest"></a>

### ListProjectsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| token | [string](#string) |  |  |
| page | [int32](#int32) |  |  |
| perPage | [int32](#int32) |  |  |
| idUser | [int32](#int32) |  |  |
| accessLevel | [int32](#int32) |  |  |






<a id=".ListProjectsResponse"></a>

### ListProjectsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [bool](#bool) |  |  |
| page | [int32](#int32) |  |  |
| lastPage | [int32](#int32) |  |  |
| projects | [Repository](#Repository) | repeated |  |
| latestPage | [bool](#bool) |  |  |






<a id=".Owner"></a>

### Owner



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| login | [string](#string) |  |  |
| avatarUrl | [string](#string) |  |  |
| type | [OwnerType](#OwnerType) |  |  |






<a id=".Repository"></a>

### Repository



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| url | [string](#string) |  |  |
| httpUrl | [string](#string) |  |  |
| defaultBranch | [string](#string) |  |  |
| lastActivityAt | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| private | [bool](#bool) |  |  |
| admin | [bool](#bool) |  |  |
| owner | [Owner](#Owner) |  |  |
| isEmpty | [bool](#bool) |  |  |






<a id=".SetWebhookRequest"></a>

### SetWebhookRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| token | [string](#string) |  |  |
| idRepo | [int32](#int32) |  |  |
| params | [string](#string) |  |  |






<a id=".SetWebhookResponse"></a>

### SetWebhookResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [bool](#bool) |  |  |
| webhookId | [int32](#int32) |  |  |





 <!-- end messages -->


<a id=".OwnerType"></a>

### OwnerType


| Name | Number | Description |
| ---- | ------ | ----------- |
| User | 0 |  |
| Group | 1 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->


<a id=".Provider"></a>

### Provider


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListProjects | [.ListProjectsRequest](#ListProjectsRequest) | [.ListProjectsResponse](#ListProjectsResponse) |  |
| SetWebhook | [.SetWebhookRequest](#SetWebhookRequest) | [.SetWebhookResponse](#SetWebhookResponse) |  |
| DeleteWebhook | [.DeleteWebhookRequest](#DeleteWebhookRequest) | [.DeleteWebhookResponse](#DeleteWebhookResponse) |  |
| CheckAccessUserInProject | [.CheckAccessUserInProjectRequest](#CheckAccessUserInProjectRequest) | [.CheckAccessUserInProjectResponse](#CheckAccessUserInProjectResponse) |  |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a id="double"></a> double |  | double | double | float | float64 | double | float | Float |
| <a id="float"></a> float |  | float | float | float | float32 | float | float | Float |
| <a id="int32"></a> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a id="int64"></a> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a id="uint32"></a> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a id="uint64"></a> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a id="sint32"></a> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a id="sint64"></a> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a id="fixed32"></a> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a id="fixed64"></a> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a id="sfixed32"></a> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a id="sfixed64"></a> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a id="bool"></a> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a id="string"></a> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a id="bytes"></a> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

