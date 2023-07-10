package utils

import (
	"fmt"
	"time"
)

func GenerateUniqueFilename() string {
	timestamp := time.Now().UnixNano()
	filename := fmt.Sprintf("image_%d.png", timestamp)
	return filename
}
