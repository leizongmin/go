package localpersistence

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"
)

func generateTempPath() string {
	return path.Join(os.TempDir(), fmt.Sprintf("localpersistence-%d-%d", time.Now().Unix(), rand.Int63()))
}
