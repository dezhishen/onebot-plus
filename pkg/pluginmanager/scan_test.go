package pluginmanager

import (
	"fmt"
	"testing"
)

func TestScan(*testing.T) {
	err := ScanInPath("F:/Test/onebot/onebot-plus/wiki", func(files []string) error {

		for _, file := range files {
			fmt.Printf("%s\n", file)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

}
