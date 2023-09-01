package pluginmanager

import (
	"fmt"
	"testing"
)

func TestScan(*testing.T) {
	err := scanPath("F:/Test/onebot/onebot-plus/wiki", func(file string) error {
		fmt.Printf("%s\n", file)
		return nil
	})
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

}
