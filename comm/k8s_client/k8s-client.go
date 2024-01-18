package k8s_client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Connect() (*kubernetes.Clientset, error) {
	// NewForConfig为给定的配置创建一个新的Clientset（如下图所示包含所有的api-versions，这样做的目的是便于其它
	// 资源类型对这个Pod进行管理和控制？）。
	configFile := getConfig()

	config, err := clientcmd.BuildConfigFromFlags("", *configFile)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
