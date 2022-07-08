package models

// 链码信息
type ChaincodeInfo struct {
	// 链码信息
	ChaincodeType    string      // 链码类型 golang java nodejs
	ChaincodePolicy  string      // 链码策略
	ChaincodeID      string      // 链码id
	ChaincodePath    string      // 链码路径
	ChaincodeVersion string      // 链码版本
	PackageID        string      // packageid
	PackagePara      interface{} // package的参数
	Sequence         int64       // 序列号
	IsInit           bool        // 是否初始化
	Args             []string    // 链码宿参数
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewChaincodeInfo(isInit bool, chaincodeType string, chaincodePolicy string, chaincodeId string, chaincodePath string, chaincodeVersion string, packageID string, PackagePara interface{}, sequence int64, args []string) *ChaincodeInfo {
	logger.Debug("NewChaincodeInfo enter")
	return &ChaincodeInfo{
		ChaincodeType:    chaincodeType,
		ChaincodePolicy:  chaincodePolicy,
		ChaincodeID:      chaincodeId,
		ChaincodePath:    chaincodePath,
		ChaincodeVersion: chaincodeVersion,
		PackageID:        packageID,
		PackagePara:      PackagePara,
		Sequence:         sequence,
		IsInit:           isInit,
		Args:             args,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (c *ChaincodeInfo) close() {
	logger.Debug("ChaincodeInfo close enter")
	c = nil

}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func loopCloseChaincodeInfo(cs []*ChaincodeInfo) {
	logger.Debug("loopCloseChaincodeInfo enter")
	for _, c := range cs {
		c.close()
	}
}
