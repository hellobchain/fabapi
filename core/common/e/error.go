package e

type ErrCode int

func (e ErrCode) Error() string {
	return getErrMsg(e)
}

func getErrMsg(e ErrCode) string {
	m, ok := MsgFlags[e]
	if !ok {
		return "未知错误信息"
	}
	return m
}
