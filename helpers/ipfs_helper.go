package helpers

import (
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

// IPFSConnect -
func IPFSConnect() (*shell.Shell, error) {
	url := os.Getenv("IPFS_URL")
	if url == "" {
		url = "localhost:5001"
	}

	return shell.NewShell(url), nil
}
