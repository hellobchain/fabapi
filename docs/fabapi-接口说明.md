# fabapi API文档


<a name="overview"></a>
## 概览
描述:fabapi 接口文档.


### 版本信息
*版本* : 1.1.0


### URI scheme
*域名* : localhost:6922  
*基础路径* : /fab




<a name="paths"></a>
## 路径

<a name="chaincode-approvecc-post"></a>
### 批准链码管理
```
POST /chaincode/approvecc
```


#### 描述
功能：批准链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|批准链码管理|[chaincode.approveChaincodeRequest](#chaincode-approvechaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 批准链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/approvecc
```


##### 请求 body
```json
{
  "chaincodeid" : "test",
  "chaincodepolicy" : "string",
  "chaincodeversion" : "v1.0",
  "channelid" : "mychanel",
  "isinit" : true,
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ],
  "packageid" : "string",
  "sequence" : 1
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-commitcc-post"></a>
### 提交链码管理
```
POST /chaincode/commitcc
```


#### 描述
功能：提交链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|提交链码管理|[chaincode.commitChaincodeRequest](#chaincode-commitchaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 提交链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/commitcc
```


##### 请求 body
```json
{
  "chaincodeid" : "test",
  "chaincodepolicy" : "string",
  "chaincodeversion" : "v1.0",
  "channelid" : "mychannel",
  "isinit" : true,
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ],
  "packageid" : "string",
  "sequence" : 1
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-initcc-post"></a>
### 初始化链码管理
```
POST /chaincode/initcc
```


#### 描述
功能：初始化链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|初始化链码管理|[chaincode.initChaincodeRequest](#chaincode-initchaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 初始化链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/initcc
```


##### 请求 body
```json
{
  "args" : [ "save", "a", "b" ],
  "chaincodeid" : "test",
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-installcc-post"></a>
### 安装链码管理
```
POST /chaincode/installcc
```


#### 描述
功能：安装链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|安装链码管理|[chaincode.installChaincodeRequest](#chaincode-installchaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 安装链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/installcc
```


##### 请求 body
```json
{
  "ccpack" : "object",
  "chaincodeid" : "test",
  "chaincodeversion" : "v1.0",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-invokecc-post"></a>
### invoke链码管理
```
POST /chaincode/invokecc
```


#### 描述
功能：invoke链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|invoke链码管理|[chaincode.invokeChaincodeRequest](#chaincode-invokechaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* invoke链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/invokecc
```


##### 请求 body
```json
{
  "args" : [ "save", "a", "b" ],
  "chaincodeid" : "test",
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-packagecc-post"></a>
### 打包链码管理
```
POST /chaincode/packagecc
```


#### 描述
功能：打包链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|打包链码管理|[chaincode.packageChaincodeRequest](#chaincode-packagechaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 打包链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/packagecc
```


##### 请求 body
```json
{
  "chaincodeid" : "test",
  "chaincodesrcpath" : "src/github.com/chaincode",
  "chaincodetype" : "GOLANG",
  "chaincodeversion" : "v1.0"
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-querycc-post"></a>
### query链码管理
```
POST /chaincode/querycc
```


#### 描述
功能：query链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|query链码管理|[chaincode.queryChaincodeRequest](#chaincode-querychaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* query链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/querycc
```


##### 请求 body
```json
{
  "args" : [ "query", "a" ],
  "chaincodeid" : "test",
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-upgradecc-post"></a>
### 升级链码管理
```
POST /chaincode/upgradecc
```


#### 描述
功能：升级链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|升级链码管理|[chaincode.upgradeChaincodeRequest](#chaincode-upgradechaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 升级链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/upgradecc
```


##### 请求 body
```json
{
  "chaincodeid" : "test",
  "chaincodepolicy" : "string",
  "chaincodeversion" : "v1.0",
  "channelid" : "mychanel",
  "isinit" : true,
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ],
  "packageid" : "string",
  "sequence" : 1
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="channel-createchannel-post"></a>
### 创建通道管理
```
POST /channel/createchannel
```


#### 描述
功能：创建通道管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|创建通道管理|[channel.createChannelRequest](#channel-createchannelrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 创建通道管理


#### HTTP 请求示例

##### 请求 path
```
/channel/createchannel
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orderernames" : [ "orderer0", "orderer1", "orderer2" ],
  "ordererorgname" : "orderer",
  "ordererusername" : "Admin",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="channel-joinchannel-post"></a>
### 加入通道管理
```
POST /channel/joinchannel
```


#### 描述
功能：加入通道管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|加入通道管理|[channel.joinChannelRequest](#channel-joinchannelrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 加入通道管理


#### HTTP 请求示例

##### 请求 path
```
/channel/joinchannel
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orderernames" : [ "orderer0", "orderer1", "orderer2" ],
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryblock-post"></a>
### 查询块管理
```
POST /ledger/queryblock
```


#### 描述
功能：查询块管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|查询块管理|[ledger.queryBlockRequest](#ledger-queryblockrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 查询块管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryblock
```


##### 请求 body
```json
{
  "blocknum" : 0,
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryblockbyhash-post"></a>
### 通过哈希查询块管理
```
POST /ledger/queryblockbyhash
```


#### 描述
功能：通过哈希查询块管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|通过哈希查询块管理|[ledger.queryBlockByHashRequest](#ledger-queryblockbyhashrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 通过哈希查询块管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryblockbyhash
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "hash" : "0xacedababacedababacedababacedababacedababacedababacedababacedabab",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryblockbytxid-post"></a>
### 通过交易id查询块管理
```
POST /ledger/queryblockbytxid
```


#### 描述
功能：通过交易id查询块管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|通过交易id查询块管理|[ledger.queryBlockByTxIdRequest](#ledger-queryblockbytxidrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 通过交易id查询块管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryblockbytxid
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ],
  "txid" : "0xacedababacedababacedababacedababacedababacedababacedababacedabab"
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryblocknum-post"></a>
### 查询块高管理
```
POST /ledger/queryblocknum
```


#### 描述
功能：查询块高管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|查询块高管理|[ledger.queryBlockNumRequest](#ledger-queryblocknumrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 查询块高


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryblocknum
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryconfig-post"></a>
### 查询配置块管理
```
POST /ledger/queryconfig
```


#### 描述
功能：查询配置块管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|查询配置块管理|[ledger.queryConfigRequest](#ledger-queryconfigrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 查询配置块管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryconfig
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryledger-post"></a>
### 查询账本管理
```
POST /ledger/queryledger
```


#### 描述
功能：查询账本管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|查询账本管理|[ledger.queryLedgerRequest](#ledger-queryledgerrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 查询账本管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryledger
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-querytx-post"></a>
### 查询交易管理
```
POST /ledger/querytx
```


#### 描述
功能：查询交易管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|查询交易管理|[ledger.queryTxRequest](#ledger-querytxrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 查询交易管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/querytx
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ],
  "txid" : "0xacedababacedababacedababacedababacedababacedababacedababacedabab"
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="log-setloglevel-post"></a>
### 设置日志级别
```
POST /log/setloglevel
```


#### 描述
功能：设置日志级别


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|设置日志级别|[log.setLogLevelRequest](#log-setloglevelrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 设置日志级别


#### HTTP 请求示例

##### 请求 path
```
/log/setloglevel
```


##### 请求 body
```json
{
  "loglevel" : "info"
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```




<a name="paths"></a>
## 路径

<a name="chaincode-approvecc-post"></a>
### 批准链码管理
```
POST /chaincode/approvecc
```


#### 描述
功能：批准链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|批准链码管理|[chaincode.approveChaincodeRequest](#chaincode-approvechaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 批准链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/approvecc
```


##### 请求 body
```json
{
  "chaincodeid" : "test",
  "chaincodepolicy" : "string",
  "chaincodeversion" : "v1.0",
  "channelid" : "mychanel",
  "isinit" : true,
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ],
  "packageid" : "string",
  "sequence" : 1
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-commitcc-post"></a>
### 提交链码管理
```
POST /chaincode/commitcc
```


#### 描述
功能：提交链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|提交链码管理|[chaincode.commitChaincodeRequest](#chaincode-commitchaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 提交链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/commitcc
```


##### 请求 body
```json
{
  "chaincodeid" : "test",
  "chaincodepolicy" : "string",
  "chaincodeversion" : "v1.0",
  "channelid" : "mychannel",
  "isinit" : true,
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ],
  "packageid" : "string",
  "sequence" : 1
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-initcc-post"></a>
### 初始化链码管理
```
POST /chaincode/initcc
```


#### 描述
功能：初始化链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|初始化链码管理|[chaincode.initChaincodeRequest](#chaincode-initchaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 初始化链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/initcc
```


##### 请求 body
```json
{
  "args" : [ "save", "a", "b" ],
  "chaincodeid" : "test",
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-installcc-post"></a>
### 安装链码管理
```
POST /chaincode/installcc
```


#### 描述
功能：安装链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|安装链码管理|[chaincode.installChaincodeRequest](#chaincode-installchaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 安装链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/installcc
```


##### 请求 body
```json
{
  "ccpack" : "object",
  "chaincodeid" : "test",
  "chaincodeversion" : "v1.0",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-invokecc-post"></a>
### invoke链码管理
```
POST /chaincode/invokecc
```


#### 描述
功能：invoke链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|invoke链码管理|[chaincode.invokeChaincodeRequest](#chaincode-invokechaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* invoke链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/invokecc
```


##### 请求 body
```json
{
  "args" : [ "save", "a", "b" ],
  "chaincodeid" : "test",
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-packagecc-post"></a>
### 打包链码管理
```
POST /chaincode/packagecc
```


#### 描述
功能：打包链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|打包链码管理|[chaincode.packageChaincodeRequest](#chaincode-packagechaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 打包链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/packagecc
```


##### 请求 body
```json
{
  "chaincodeid" : "test",
  "chaincodesrcpath" : "src/github.com/chaincode",
  "chaincodetype" : "GOLANG",
  "chaincodeversion" : "v1.0"
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-querycc-post"></a>
### query链码管理
```
POST /chaincode/querycc
```


#### 描述
功能：query链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|query链码管理|[chaincode.queryChaincodeRequest](#chaincode-querychaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* query链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/querycc
```


##### 请求 body
```json
{
  "args" : [ "query", "a" ],
  "chaincodeid" : "test",
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="chaincode-upgradecc-post"></a>
### 升级链码管理
```
POST /chaincode/upgradecc
```


#### 描述
功能：升级链码管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|升级链码管理|[chaincode.upgradeChaincodeRequest](#chaincode-upgradechaincoderequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 升级链码管理


#### HTTP 请求示例

##### 请求 path
```
/chaincode/upgradecc
```


##### 请求 body
```json
{
  "chaincodeid" : "test",
  "chaincodepolicy" : "string",
  "chaincodeversion" : "v1.0",
  "channelid" : "mychanel",
  "isinit" : true,
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ],
  "packageid" : "string",
  "sequence" : 1
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="channel-createchannel-post"></a>
### 创建通道管理
```
POST /channel/createchannel
```


#### 描述
功能：创建通道管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|创建通道管理|[channel.createChannelRequest](#channel-createchannelrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 创建通道管理


#### HTTP 请求示例

##### 请求 path
```
/channel/createchannel
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orderernames" : [ "orderer0", "orderer1", "orderer2" ],
  "ordererorgname" : "orderer",
  "ordererusername" : "Admin",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="channel-joinchannel-post"></a>
### 加入通道管理
```
POST /channel/joinchannel
```


#### 描述
功能：加入通道管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|加入通道管理|[channel.joinChannelRequest](#channel-joinchannelrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 加入通道管理


#### HTTP 请求示例

##### 请求 path
```
/channel/joinchannel
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orderernames" : [ "orderer0", "orderer1", "orderer2" ],
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryblock-post"></a>
### 查询块管理
```
POST /ledger/queryblock
```


#### 描述
功能：查询块管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|查询块管理|[ledger.queryBlockRequest](#ledger-queryblockrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 查询块管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryblock
```


##### 请求 body
```json
{
  "blocknum" : 0,
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryblockbyhash-post"></a>
### 通过哈希查询块管理
```
POST /ledger/queryblockbyhash
```


#### 描述
功能：通过哈希查询块管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|通过哈希查询块管理|[ledger.queryBlockByHashRequest](#ledger-queryblockbyhashrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 通过哈希查询块管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryblockbyhash
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "hash" : "0xacedababacedababacedababacedababacedababacedababacedababacedabab",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryblockbytxid-post"></a>
### 通过交易id查询块管理
```
POST /ledger/queryblockbytxid
```


#### 描述
功能：通过交易id查询块管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|通过交易id查询块管理|[ledger.queryBlockByTxIdRequest](#ledger-queryblockbytxidrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 通过交易id查询块管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryblockbytxid
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ],
  "txid" : "0xacedababacedababacedababacedababacedababacedababacedababacedabab"
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryblocknum-post"></a>
### 查询块高管理
```
POST /ledger/queryblocknum
```


#### 描述
功能：查询块高管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|查询块高管理|[ledger.queryBlockNumRequest](#ledger-queryblocknumrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 查询块高


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryblocknum
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryconfig-post"></a>
### 查询配置块管理
```
POST /ledger/queryconfig
```


#### 描述
功能：查询配置块管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|查询配置块管理|[ledger.queryConfigRequest](#ledger-queryconfigrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 查询配置块管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryconfig
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-queryledger-post"></a>
### 查询账本管理
```
POST /ledger/queryledger
```


#### 描述
功能：查询账本管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|查询账本管理|[ledger.queryLedgerRequest](#ledger-queryledgerrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 查询账本管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/queryledger
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ]
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="ledger-querytx-post"></a>
### 查询交易管理
```
POST /ledger/querytx
```


#### 描述
功能：查询交易管理


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|查询交易管理|[ledger.queryTxRequest](#ledger-querytxrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 查询交易管理


#### HTTP 请求示例

##### 请求 path
```
/ledger/querytx
```


##### 请求 body
```json
{
  "channelid" : "mychannel",
  "orgmsps" : [ "Org1MSP", "Org2MSP", "Org3MSP" ],
  "orgnames" : [ "org1", "org2", "org3" ],
  "txid" : "0xacedababacedababacedababacedababacedababacedababacedababacedabab"
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


<a name="log-setloglevel-post"></a>
### 设置日志级别
```
POST /log/setloglevel
```


#### 描述
功能：设置日志级别


#### 参数

|类型|名称|描述|Schema|
|---|---|---|---|
|**Body**|**resource**  <br>*必填*|设置日志级别|[log.setLogLevelRequest](#log-setloglevelrequest)|


#### 响应

|HTTP 代码|描述|Schema|
|---|---|---|
|**200**|OK|[gintool.ApiResponse](#gintool-apiresponse)|
|**400**|Bad Request|[gintool.ApiResponse](#gintool-apiresponse)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### 标签

* 设置日志级别


#### HTTP 请求示例

##### 请求 path
```
/log/setloglevel
```


##### 请求 body
```json
{
  "loglevel" : "info"
}
```


#### HTTP 响应示例

##### 响应 200
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```


##### 响应 400
```json
{
  "code" : 0,
  "err_msg" : "string",
  "message" : "string",
  "result" : "object"
}
```




<a name="definitions"></a>
## 定义

<a name="chaincode-approvechaincoderequest"></a>
### chaincode.approveChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**chaincodepolicy**  <br>*可选*|**示例** : `"string"`|string|
|**chaincodeversion**  <br>*必填*|**示例** : `"v1.0"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychanel"`|string|
|**isinit**  <br>*可选*|**示例** : `true`|boolean|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|
|**packageid**  <br>*必填*|**示例** : `"string"`|string|
|**sequence**  <br>*必填*|**示例** : `1`|integer|


<a name="chaincode-commitchaincoderequest"></a>
### chaincode.commitChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**chaincodepolicy**  <br>*可选*|**示例** : `"string"`|string|
|**chaincodeversion**  <br>*必填*|**示例** : `"v1.0"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**isinit**  <br>*可选*|**示例** : `true`|boolean|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|
|**packageid**  <br>*必填*|**示例** : `"string"`|string|
|**sequence**  <br>*必填*|**示例** : `1`|integer|


<a name="chaincode-initchaincoderequest"></a>
### chaincode.initChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**args**  <br>*可选*|**示例** : `[ "save", "a", "b" ]`|< string > array|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="chaincode-installchaincoderequest"></a>
### chaincode.installChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**ccpack**  <br>*可选*|**示例** : `"object"`|object|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**chaincodeversion**  <br>*必填*|**示例** : `"v1.0"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="chaincode-invokechaincoderequest"></a>
### chaincode.invokeChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**args**  <br>*必填*|**示例** : `[ "save", "a", "b" ]`|< string > array|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="chaincode-packagechaincoderequest"></a>
### chaincode.packageChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**chaincodesrcpath**  <br>*必填*|**示例** : `"src/github.com/chaincode"`|string|
|**chaincodetype**  <br>*必填*|**示例** : `"GOLANG"`|string|
|**chaincodeversion**  <br>*必填*|**示例** : `"v1.0"`|string|


<a name="chaincode-querychaincoderequest"></a>
### chaincode.queryChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**args**  <br>*必填*|**示例** : `[ "query", "a" ]`|< string > array|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="chaincode-upgradechaincoderequest"></a>
### chaincode.upgradeChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**chaincodepolicy**  <br>*可选*|**示例** : `"string"`|string|
|**chaincodeversion**  <br>*必填*|**示例** : `"v1.0"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychanel"`|string|
|**isinit**  <br>*可选*|**示例** : `true`|boolean|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|
|**packageid**  <br>*必填*|**示例** : `"string"`|string|
|**sequence**  <br>*必填*|**示例** : `1`|integer|


<a name="channel-createchannelrequest"></a>
### channel.createChannelRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*可选*|**示例** : `"mychannel"`|string|
|**orderernames**  <br>*必填*|**示例** : `[ "orderer0", "orderer1", "orderer2" ]`|< string > array|
|**ordererorgname**  <br>*可选*|**示例** : `"orderer"`|string|
|**ordererusername**  <br>*可选*|**示例** : `"Admin"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="channel-joinchannelrequest"></a>
### channel.joinChannelRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*可选*|**示例** : `"mychannel"`|string|
|**orderernames**  <br>*必填*|**示例** : `[ "orderer0", "orderer1", "orderer2" ]`|< string > array|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="gintool-apiresponse"></a>
### gintool.ApiResponse

|名称|描述|Schema|
|---|---|---|
|**code**  <br>*可选*|状态码  <br>**示例** : `0`|integer|
|**err_msg**  <br>*可选*|内部错误详情  <br>**示例** : `"string"`|string|
|**message**  <br>*可选*|状态短语  <br>**示例** : `"string"`|string|
|**result**  <br>*可选*|数据结果集  <br>**示例** : `"object"`|object|


<a name="ledger-queryblockbyhashrequest"></a>
### ledger.queryBlockByHashRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**hash**  <br>*必填*|**示例** : `"0xacedababacedababacedababacedababacedababacedababacedababacedabab"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="ledger-queryblockbytxidrequest"></a>
### ledger.queryBlockByTxIdRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|
|**txid**  <br>*必填*|**示例** : `"0xacedababacedababacedababacedababacedababacedababacedababacedabab"`|string|


<a name="ledger-queryblocknumrequest"></a>
### ledger.queryBlockNumRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*可选*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="ledger-queryblockrequest"></a>
### ledger.queryBlockRequest

|名称|描述|Schema|
|---|---|---|
|**blocknum**  <br>*可选*|**示例** : `0`|integer|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="ledger-queryconfigrequest"></a>
### ledger.queryConfigRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="ledger-queryledgerrequest"></a>
### ledger.queryLedgerRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="ledger-querytxrequest"></a>
### ledger.queryTxRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*可选*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|
|**txid**  <br>*必填*|**示例** : `"0xacedababacedababacedababacedababacedababacedababacedababacedabab"`|string|


<a name="log-setloglevelrequest"></a>
### log.setLogLevelRequest

|名称|描述|Schema|
|---|---|---|
|**loglevel**  <br>*必填*|**示例** : `"info"`|string|






<a name="definitions"></a>
## 定义

<a name="chaincode-approvechaincoderequest"></a>
### chaincode.approveChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**chaincodepolicy**  <br>*可选*|**示例** : `"string"`|string|
|**chaincodeversion**  <br>*必填*|**示例** : `"v1.0"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychanel"`|string|
|**isinit**  <br>*可选*|**示例** : `true`|boolean|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|
|**packageid**  <br>*必填*|**示例** : `"string"`|string|
|**sequence**  <br>*必填*|**示例** : `1`|integer|


<a name="chaincode-commitchaincoderequest"></a>
### chaincode.commitChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**chaincodepolicy**  <br>*可选*|**示例** : `"string"`|string|
|**chaincodeversion**  <br>*必填*|**示例** : `"v1.0"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**isinit**  <br>*可选*|**示例** : `true`|boolean|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|
|**packageid**  <br>*必填*|**示例** : `"string"`|string|
|**sequence**  <br>*必填*|**示例** : `1`|integer|


<a name="chaincode-initchaincoderequest"></a>
### chaincode.initChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**args**  <br>*可选*|**示例** : `[ "save", "a", "b" ]`|< string > array|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="chaincode-installchaincoderequest"></a>
### chaincode.installChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**ccpack**  <br>*可选*|**示例** : `"object"`|object|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**chaincodeversion**  <br>*必填*|**示例** : `"v1.0"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="chaincode-invokechaincoderequest"></a>
### chaincode.invokeChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**args**  <br>*必填*|**示例** : `[ "save", "a", "b" ]`|< string > array|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="chaincode-packagechaincoderequest"></a>
### chaincode.packageChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**chaincodesrcpath**  <br>*必填*|**示例** : `"src/github.com/chaincode"`|string|
|**chaincodetype**  <br>*必填*|**示例** : `"GOLANG"`|string|
|**chaincodeversion**  <br>*必填*|**示例** : `"v1.0"`|string|


<a name="chaincode-querychaincoderequest"></a>
### chaincode.queryChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**args**  <br>*必填*|**示例** : `[ "query", "a" ]`|< string > array|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="chaincode-upgradechaincoderequest"></a>
### chaincode.upgradeChaincodeRequest

|名称|描述|Schema|
|---|---|---|
|**chaincodeid**  <br>*必填*|**示例** : `"test"`|string|
|**chaincodepolicy**  <br>*可选*|**示例** : `"string"`|string|
|**chaincodeversion**  <br>*必填*|**示例** : `"v1.0"`|string|
|**channelid**  <br>*必填*|**示例** : `"mychanel"`|string|
|**isinit**  <br>*可选*|**示例** : `true`|boolean|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|
|**packageid**  <br>*必填*|**示例** : `"string"`|string|
|**sequence**  <br>*必填*|**示例** : `1`|integer|


<a name="channel-createchannelrequest"></a>
### channel.createChannelRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*可选*|**示例** : `"mychannel"`|string|
|**orderernames**  <br>*必填*|**示例** : `[ "orderer0", "orderer1", "orderer2" ]`|< string > array|
|**ordererorgname**  <br>*可选*|**示例** : `"orderer"`|string|
|**ordererusername**  <br>*可选*|**示例** : `"Admin"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="channel-joinchannelrequest"></a>
### channel.joinChannelRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*可选*|**示例** : `"mychannel"`|string|
|**orderernames**  <br>*必填*|**示例** : `[ "orderer0", "orderer1", "orderer2" ]`|< string > array|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="gintool-apiresponse"></a>
### gintool.ApiResponse

|名称|描述|Schema|
|---|---|---|
|**code**  <br>*可选*|状态码  <br>**示例** : `0`|integer|
|**err_msg**  <br>*可选*|内部错误详情  <br>**示例** : `"string"`|string|
|**message**  <br>*可选*|状态短语  <br>**示例** : `"string"`|string|
|**result**  <br>*可选*|数据结果集  <br>**示例** : `"object"`|object|


<a name="ledger-queryblockbyhashrequest"></a>
### ledger.queryBlockByHashRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**hash**  <br>*必填*|**示例** : `"0xacedababacedababacedababacedababacedababacedababacedababacedabab"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="ledger-queryblockbytxidrequest"></a>
### ledger.queryBlockByTxIdRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|
|**txid**  <br>*必填*|**示例** : `"0xacedababacedababacedababacedababacedababacedababacedababacedabab"`|string|


<a name="ledger-queryblocknumrequest"></a>
### ledger.queryBlockNumRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*可选*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="ledger-queryblockrequest"></a>
### ledger.queryBlockRequest

|名称|描述|Schema|
|---|---|---|
|**blocknum**  <br>*可选*|**示例** : `0`|integer|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="ledger-queryconfigrequest"></a>
### ledger.queryConfigRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="ledger-queryledgerrequest"></a>
### ledger.queryLedgerRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*必填*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|


<a name="ledger-querytxrequest"></a>
### ledger.queryTxRequest

|名称|描述|Schema|
|---|---|---|
|**channelid**  <br>*可选*|**示例** : `"mychannel"`|string|
|**orgmsps**  <br>*必填*|**示例** : `[ "Org1MSP", "Org2MSP", "Org3MSP" ]`|< string > array|
|**orgnames**  <br>*必填*|**示例** : `[ "org1", "org2", "org3" ]`|< string > array|
|**txid**  <br>*必填*|**示例** : `"0xacedababacedababacedababacedababacedababacedababacedababacedabab"`|string|


<a name="log-setloglevelrequest"></a>
### log.setLogLevelRequest

|名称|描述|Schema|
|---|---|---|
|**loglevel**  <br>*必填*|**示例** : `"info"`|string|





