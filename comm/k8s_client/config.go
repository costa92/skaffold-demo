package k8s_client

import (
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

// Path: comm/k8s_client/config.go
func getConfig() *string {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig := filepath.Join(home, ".kube", "config")
		return &kubeconfig
	}
	return nil
}
