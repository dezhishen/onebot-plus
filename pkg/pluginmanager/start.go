package pluginmanager

import (
	"context"
)

func getRootDir() string {
	return "./plugins"
}

func Init(ctx context.Context) error {
	// scan and log all plugins
	return nil
}
