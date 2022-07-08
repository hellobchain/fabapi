/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package models

var DefaultPackagePara interface{} = nil // 默认的package的参数值

var DefaultArgs []string = nil // 默认的参数

const (
	DefaultOrgAdminUser        = "Admin"
	DefaultOrgUser             = "User1"
	DefaultOrgAnchorFile       = "anchors.tx"
	DefaultOrdererOrgAdminUser = "Admin"
	DefaultOrdererOrgName      = "ordererOrg"
	DefaultChaincodePath       = ""
	DefaultPackageId           = ""
	DefaultSequence            = -1
	DefaultChaincodeVersion    = ""
	DefaultPath                = "/fab/config/channel-artifacts"
	DefaultCcpackagePath       = "/tmp/chaincode"
	DefaultChaincodeType       = "GOLANG"
	DefaultChaincodePolicy     = ""
	DefaultEndorsementPlugin   = "escc"
	DefaultValidationPlugin    = "vscc"
	EmptyReturn                = ""
	Empty                      = ""
	SUCCESS                    = 200
	TAB                        = "    "
)
