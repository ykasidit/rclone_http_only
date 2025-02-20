// Package all imports all the backends
package all

import (
	// Active file systems
	_ "github.com/rclone/rclone/backend/http"
	_ "github.com/rclone/rclone/backend/local"
)
