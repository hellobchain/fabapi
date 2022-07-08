/***************************************************************
 * @file       程序文件名称
 * @brief      程序文件的功能
 * @author     wsw
 * @version    v1
 * @date       2021.12.20
 **************************************************************/
package service

import "fabapi/core/fabsdk/models"

type ChaincodeService interface {

	// package chaincode
	PackageChaincode(chaincodeInfo *models.ChaincodeInfo) (string, interface{}, string, error)

	// install chaincode
	InstallChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo) (string, error)

	// approve chaincode
	ApproveChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) error

	// upgrade chaincode
	UpgradeChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) error

	// commit chaincode
	CommitChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) error

	// init chaincode
	InitChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) (string, error)

	// invoke chaincode
	InvokeChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) (string, error)

	// query chaincode
	QueryChaincode(chaincodeInfo *models.ChaincodeInfo, orgs []*models.PeerOrgInfo, channelID string) (string, error)

	Close()
}
