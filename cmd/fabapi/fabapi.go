// Package fabapi /***************************************************************
package fabapi

import (
	"os"
)

// Main /***************************************************************
func Main() { // main的入口
	logger.Debug("Main enter")
	err := newApp().Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
