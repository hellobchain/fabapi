package models

// orderer组织信息
type OrdererOrgInfo struct {
	OrdererOrgAdminUser string   // like "Admin"   orderer组织的管理员
	OrdererOrgName      string   // like "OrdererOrg"  orderer组织的名称
	OrdererOrgEndpoint  []string // orderer组织下的orderer成员
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewOrdererOrgInfo(ordererOrgAdminUser string, ordererOrgName string, ordererOrgEndpoint []string) *OrdererOrgInfo {
	logger.Debug("NewOrdererOrgInfo enter")
	if ordererOrgName == Empty {
		ordererOrgName = DefaultOrdererOrgName
	}
	if ordererOrgAdminUser == Empty {
		ordererOrgAdminUser = DefaultOrdererOrgAdminUser
	}
	return &OrdererOrgInfo{
		OrdererOrgAdminUser: ordererOrgAdminUser,
		OrdererOrgName:      ordererOrgName,
		OrdererOrgEndpoint:  ordererOrgEndpoint,
	}
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func NewDefaultOrdererOrgInfo(ordererOrgEndpoint []string) *OrdererOrgInfo {
	logger.Debug("NewDefaultOrdererOrgInfo enter")
	return NewOrdererOrgInfo(DefaultOrdererOrgAdminUser, DefaultOrdererOrgName, ordererOrgEndpoint)
}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func (o *OrdererOrgInfo) close() {
	logger.Debug("OrdererOrgInfo close enter")
	o = nil

}

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func loopCloseO(os []*OrdererOrgInfo) {
	logger.Debug("loopCloseO enter")
	for _, o := range os {
		o.close()
	}
}
