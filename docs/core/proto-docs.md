<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [iou/iou.proto](#iou/iou.proto)
    - [IouNamespace](#ioux.iou.IouNamespace)
    - [MsgBurnIou](#ioux.iou.MsgBurnIou)
    - [MsgCreateIouNamespace](#ioux.iou.MsgCreateIouNamespace)
    - [MsgIssueIou](#ioux.iou.MsgIssueIou)
    - [MsgUpdateIouNamespaceAdmin](#ioux.iou.MsgUpdateIouNamespaceAdmin)
    - [MsgUpdateIouNamespaceIssuer](#ioux.iou.MsgUpdateIouNamespaceIssuer)
    - [Params](#ioux.iou.Params)
  
- [iou/genesis.proto](#iou/genesis.proto)
    - [GenesisState](#ioux.iou.GenesisState)
  
- [iou/query.proto](#iou/query.proto)
    - [QueryAllIouNamespaceRequest](#ioux.iou.QueryAllIouNamespaceRequest)
    - [QueryAllIouNamespaceResponse](#ioux.iou.QueryAllIouNamespaceResponse)
    - [QueryGetIouNamespaceRequest](#ioux.iou.QueryGetIouNamespaceRequest)
    - [QueryGetIouNamespaceResponse](#ioux.iou.QueryGetIouNamespaceResponse)
    - [QueryParamsRequest](#ioux.iou.QueryParamsRequest)
    - [QueryParamsResponse](#ioux.iou.QueryParamsResponse)
  
    - [Query](#ioux.iou.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="iou/iou.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## iou/iou.proto



<a name="ioux.iou.IouNamespace"></a>

### IouNamespace



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `admin` | [string](#string) |  |  |
| `issuer` | [string](#string) |  |  |






<a name="ioux.iou.MsgBurnIou"></a>

### MsgBurnIou



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `namespace_id` | [string](#string) |  |  |
| `prefix` | [string](#string) |  |  |
| `base_denom` | [string](#string) |  |  |
| `holder` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="ioux.iou.MsgCreateIouNamespace"></a>

### MsgCreateIouNamespace



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `admin` | [string](#string) |  |  |
| `issuer` | [string](#string) |  |  |






<a name="ioux.iou.MsgIssueIou"></a>

### MsgIssueIou



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `namespace_id` | [string](#string) |  |  |
| `prefix` | [string](#string) |  |  |
| `base_denom` | [string](#string) |  |  |
| `issuer` | [string](#string) |  |  |
| `destination` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="ioux.iou.MsgUpdateIouNamespaceAdmin"></a>

### MsgUpdateIouNamespaceAdmin



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `admin_before` | [string](#string) |  |  |
| `admin_after` | [string](#string) |  |  |






<a name="ioux.iou.MsgUpdateIouNamespaceIssuer"></a>

### MsgUpdateIouNamespaceIssuer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `admin` | [string](#string) |  |  |
| `issuer` | [string](#string) |  |  |






<a name="ioux.iou.Params"></a>

### Params






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="iou/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## iou/genesis.proto



<a name="ioux.iou.GenesisState"></a>

### GenesisState
GenesisState defines the auction module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ioux.iou.Params) |  |  |
| `iou_namespaces` | [IouNamespace](#ioux.iou.IouNamespace) | repeated | this line is used by starport scaffolding # genesis/proto/state

this line is used by starport scaffolding # genesis/proto/stateField |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="iou/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## iou/query.proto



<a name="ioux.iou.QueryAllIouNamespaceRequest"></a>

### QueryAllIouNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="ioux.iou.QueryAllIouNamespaceResponse"></a>

### QueryAllIouNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `iou_namespaces` | [IouNamespace](#ioux.iou.IouNamespace) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="ioux.iou.QueryGetIouNamespaceRequest"></a>

### QueryGetIouNamespaceRequest
this line is used by starport scaffolding # 3


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |






<a name="ioux.iou.QueryGetIouNamespaceResponse"></a>

### QueryGetIouNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `iou_namepsace` | [IouNamespace](#ioux.iou.IouNamespace) |  |  |






<a name="ioux.iou.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="ioux.iou.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#ioux.iou.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ioux.iou.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#ioux.iou.QueryParamsRequest) | [QueryParamsResponse](#ioux.iou.QueryParamsResponse) |  | GET|/ioux/iou/params|
| `IouNamespace` | [QueryGetIouNamespaceRequest](#ioux.iou.QueryGetIouNamespaceRequest) | [QueryGetIouNamespaceResponse](#ioux.iou.QueryGetIouNamespaceResponse) | this line is used by starport scaffolding # 2 | GET|/ioux/iou/iou_namespaces/{id}|
| `IouNamespaceAll` | [QueryAllIouNamespaceRequest](#ioux.iou.QueryAllIouNamespaceRequest) | [QueryAllIouNamespaceResponse](#ioux.iou.QueryAllIouNamespaceResponse) |  | GET|/ioux/iou/iou_namespaces|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
