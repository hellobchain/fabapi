package pprof

import (
	"fmt"
	"net/http"

	_ "net/http/pprof"
)

var logger = wlogging.MustGetLoggerWithoutName()

/***************************************************************
 *  @brief     函数作用
 *  @param     参数
 *  @note      备注
 *  @Sample usage:     函数的使用方法
**************************************************************/
func StartPProf(port string) {
	go func() {
		addr := fmt.Sprintf(":%v", port)
		logger.Debugf("pprof start at [%s]", addr)
		logger.Debugf("http://127.0.0.1:%s/debug/pprof/", port)
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			logger.Error(err)
		}
	}()
}
