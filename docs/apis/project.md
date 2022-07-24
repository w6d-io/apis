# Protocol Documentation
<a id="top"></a>

## Table of Contents

- [project.proto](#project.proto)
    - [Component](#.Component)
    - [GetProjectRequest](#.GetProjectRequest)
    - [GetProjectResponse](#.GetProjectResponse)
    - [LinkProjectRequest](#.LinkProjectRequest)
    - [LinkProjectResponse](#.LinkProjectResponse)
    - [Project](#.Project)
    - [ProjectStack](#.ProjectStack)
    - [ProjectStackError](#.ProjectStackError)
    - [ProjectType](#.ProjectType)
    - [ProjectsRequest](#.ProjectsRequest)
    - [ProjectsResponse](#.ProjectsResponse)
    - [RepositoryOwner](#.RepositoryOwner)
    - [Search](#.Search)
    - [UpdateProjectResponse](#.UpdateProjectResponse)
  
    - [ProjectExternal](#.ProjectExternal)
  
- [Scalar Value Types](#scalar-value-types)



<a id="project.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## project.proto



<a id=".Component"></a>

### Component



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| applicationType | [string](#string) |  |  |
| packageMgr | [string](#string) |  |  |
| path | [string](#string) |  |  |
| name | [string](#string) |  |  |
| type | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a id=".GetProjectRequest"></a>

### GetProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [int32](#int32) |  |  |






<a id=".GetProjectResponse"></a>

### GetProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [bool](#bool) |  |  |






<a id=".LinkProjectRequest"></a>

### LinkProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project | [Project](#Project) |  |  |






<a id=".LinkProjectResponse"></a>

### LinkProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [bool](#bool) |  |  |
| project | [Project](#Project) |  |  |






<a id=".Project"></a>

### Project



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| owner | [RepositoryOwner](#RepositoryOwner) |  |  |
| admin | [bool](#bool) |  |  |
| linked | [bool](#bool) |  |  |
| components | [Component](#Component) | repeated |  |
| providerId | [string](#string) |  |  |
| providerName | [string](#string) |  |  |
| providerType | [string](#string) |  |  |
| providerApiUrl | [string](#string) |  |  |
| providerRepoUrl | [string](#string) |  |  |
| repositoryId | [int32](#int32) |  |  |
| repositoryAdmin | [bool](#bool) |  |  |
| repositoryDefaultBranch | [string](#string) |  |  |
| repositoryLastActivityAt | [string](#string) |  |  |
| repositoryPrivate | [bool](#bool) |  |  |
| repositoryPathWithNamespace | [string](#string) |  |  |
| repositoryAvatarUrl | [string](#string) |  |  |
| repositoryIsEmpty | [bool](#bool) |  |  |
| type | [ProjectType](#ProjectType) |  |  |
| favorite | [bool](#bool) |  |  |






<a id=".ProjectStack"></a>

### ProjectStack



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projectId | [string](#string) |  |  |
| components | [Component](#Component) | repeated |  |






<a id=".ProjectStackError"></a>

### ProjectStackError



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projectId | [string](#string) |  |  |
| code | [string](#string) |  |  |
| message | [string](#string) |  |  |






<a id=".ProjectType"></a>

### ProjectType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| idBoundary | [int32](#int32) |  | id from scope/organisation |
| type | [string](#string) |  | individual/scope/organisation |






<a id=".ProjectsRequest"></a>

### ProjectsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| per_page | [int32](#int32) |  |  |






<a id=".ProjectsResponse"></a>

### ProjectsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projects | [Project](#Project) | repeated |  |
| page | [int32](#int32) |  |  |
| per_page | [int32](#int32) |  |  |
| eof | [bool](#bool) |  |  |






<a id=".RepositoryOwner"></a>

### RepositoryOwner



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| login | [string](#string) |  |  |
| type | [string](#string) |  |  |






<a id=".Search"></a>

### Search



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| provider | [string](#string) |  |  |






<a id=".UpdateProjectResponse"></a>

### UpdateProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [bool](#bool) |  |  |
| project | [Project](#Project) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a id=".ProjectExternal"></a>

### ProjectExternal


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetProject | [.GetProjectRequest](#GetProjectRequest) | [.GetProjectResponse](#GetProjectResponse) |  |
| GetProjects | [.ProjectsRequest](#ProjectsRequest) | [.ProjectsResponse](#ProjectsResponse) |  |
| LinkProject | [.LinkProjectRequest](#LinkProjectRequest) | [.LinkProjectResponse](#LinkProjectResponse) |  |
| UpdateProject | [.Project](#Project) | [.UpdateProjectResponse](#UpdateProjectResponse) |  |

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

