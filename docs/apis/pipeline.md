# Protocol Documentation
<a id="top"></a>

## Table of Contents

- [pipeline.proto](#pipeline.proto)
    - [Actions](#.Actions)
    - [Actions.DataEntry](#.Actions.DataEntry)
    - [Actions.EnvironmentsEntry](#.Actions.EnvironmentsEntry)
    - [Actions.ParamsEntry](#.Actions.ParamsEntry)
    - [Commit](#.Commit)
    - [Conditions](#.Conditions)
    - [Conditions.Condition](#.Conditions.Condition)
    - [Pipeline](#.Pipeline)
    - [Stage](#.Stage)
    - [Tasks](#.Tasks)
    - [Trigger](#.Trigger)
    - [Trigger.DataEntry](#.Trigger.DataEntry)
  
- [Scalar Value Types](#scalar-value-types)



<a id="pipeline.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pipeline.proto



<a id=".Actions"></a>

### Actions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| componentId | [string](#string) |  |  |
| ref | [string](#string) |  |  |
| data | [Actions.DataEntry](#Actions.DataEntry) | repeated |  |
| params | [Actions.ParamsEntry](#Actions.ParamsEntry) | repeated |  |
| environments | [Actions.EnvironmentsEntry](#Actions.EnvironmentsEntry) | repeated |  |
| startTime | [int64](#int64) |  |  |
| endTime | [int64](#int64) |  |  |
| status | [string](#string) |  |  |






<a id=".Actions.DataEntry"></a>

### Actions.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a id=".Actions.EnvironmentsEntry"></a>

### Actions.EnvironmentsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a id=".Actions.ParamsEntry"></a>

### Actions.ParamsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a id=".Commit"></a>

### Commit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| ref | [string](#string) |  |  |
| message | [string](#string) |  |  |






<a id=".Conditions"></a>

### Conditions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Conditions.Condition](#Conditions.Condition) | repeated |  |






<a id=".Conditions.Condition"></a>

### Conditions.Condition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| ref | [string](#string) |  |  |
| type | [string](#string) |  |  |
| when | [string](#string) |  |  |






<a id=".Pipeline"></a>

### Pipeline



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [string](#string) |  |  |
| pipeline_id_number | [string](#string) |  |  |
| project_id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| startTime | [int64](#int64) |  |  |
| end_time | [int64](#int64) |  |  |
| log_uri | [string](#string) |  |  |
| complete | [bool](#bool) |  |  |
| force | [bool](#bool) |  |  |
| artifacts | [bool](#bool) |  |  |
| trigger_id | [string](#string) |  |  |
| commit | [Commit](#Commit) |  |  |
| event_id | [string](#string) |  |  |
| triggers | [Trigger](#Trigger) | repeated |  |
| stages | [Stage](#Stage) | repeated |  |






<a id=".Stage"></a>

### Stage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| tasks | [Tasks](#Tasks) | repeated |  |
| status | [string](#string) |  |  |
| endTime | [int64](#int64) |  |  |
| startTime | [int64](#int64) |  |  |






<a id=".Tasks"></a>

### Tasks



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| skipOnFailure | [bool](#bool) |  |  |
| conditions | [Conditions](#Conditions) | repeated |  |
| actions | [Actions](#Actions) | repeated |  |
| startTime | [int64](#int64) |  |  |
| endTime | [int64](#int64) |  |  |
| status | [string](#string) |  |  |






<a id=".Trigger"></a>

### Trigger



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| ref | [string](#string) |  |  |
| type | [string](#string) |  |  |
| componentId | [string](#string) |  |  |
| data | [Trigger.DataEntry](#Trigger.DataEntry) | repeated |  |






<a id=".Trigger.DataEntry"></a>

### Trigger.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

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

