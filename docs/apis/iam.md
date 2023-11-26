# Protocol Documentation
<a id="top"></a>

## Table of Contents

- [iam.proto](#iam.proto)
    - [Input](#permission.Input)
    - [Reply](#permission.Reply)
  
    - [Mode](#permission.Mode)
  
    - [Iam](#permission.Iam)
  
- [Scalar Value Types](#scalar-value-types)



<a id="iam.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## iam.proto



<a id="permission.Input"></a>

### Input



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| perm_type | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| value | [string](#string) |  |  |
| mode | [Mode](#permission.Mode) |  |  |






<a id="permission.Reply"></a>

### Reply






 <!-- end messages -->


<a id="permission.Mode"></a>

### Mode


| Name | Number | Description |
| ---- | ------ | ----------- |
| Admin | 0 |  |
| Public | 1 |  |
| Trait | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->


<a id="permission.Iam"></a>

### Iam
Service for handling admin (internal) projectmgt action

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddPermission | [Input](#permission.Input) | [Reply](#permission.Reply) |  |
| RemovePermission | [Input](#permission.Input) | [Reply](#permission.Reply) |  |
| ReplacePermission | [Input](#permission.Input) | [Reply](#permission.Reply) |  |

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

