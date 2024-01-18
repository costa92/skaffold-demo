# 创建 k8s RBAC

## 创建 ServiceAccount
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: skaffold
```

## 创建 ClusterRole

该文件定义了一个 ClusterRole，它允许用户查看 Pod 和 Node 资源
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: skaffold-example-clusterrole
  # "namespace" 被忽略，因为 ClusterRoles 不受名字空间限制
#  namespace: default
rules:
  - apiGroups: [""]
    resources: ["pods","nodes"]
    verbs: ["get", "watch", "list"]
```

## 创建 ClusterRoleBinding

该文件定义了一个 ClusterRoleBinding，它将用户 skaffold 与 ClusterRole skaffold-example-clusterrole 绑定

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: skaffold-example-clusterrolebinding
roleRef:
  kind: ClusterRole
  name: skaffold-example-clusterrole
  apiGroup: rbac.authorization.k8s.io
subjects:
- kind: ServiceAccount
  name: skaffold
  namespace: default
```

## 创建 RBAC 需要在 Deployment 使用,所以需要配置 `serviceAccount: skaffold`
          
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-mod-image
  labels:
    app: go-mod-image
spec:
  replicas: 1
  selector:
    matchLabels:
      app: go-mod-image
  template:
    metadata:
      labels:
        app: go-mod-image
    spec:
      containers:
      - name: go-mod-image
        image: costa92/go-mod-image
      serviceAccount: skaffold  # 配置 RBAC 的账号
```

## 然后使用 项目使用的账号登录 k8s 集群

```go
func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		}

		svc, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("There are %d svc in the cluster\n", len(svc.Items))
		}
		time.Sleep(10 * time.Second)
	}
}
```

注意，上面代码 `clientset.CoreV1().Services("").List` 会报错，因为没有权限，所以需要配置 RBAC, skaffold 账号只配置了 node与pod的权限

运行代码：
```shell
skaffold run 
```

结果：
```shell
...
[go-mod-image] There are 25 pods in the cluster
[go-mod-image] services is forbidden: User "system:serviceaccount:default:skaffold" cannot list resource "services" in API group "" at the cluster scope

```

使用 InClusterConfig 查看 treafik 代码 k8s 
[treafik github](https://github.com/traefik/traefik/blob/master/pkg/provider/kubernetes/crd/client.go#L99)

```shell
func newInClusterClient(endpoint string) (*clientWrapper, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create in-cluster configuration: %w", err)
	}

	if endpoint != "" {
		config.Host = endpoint
	}

	return createClientFromConfig(config)
}
```

## 参考：
- [kubernetes RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
- [RBAC 中文文档]( https://kubernetes.io/zh-cn/docs/reference/access-authn-authz/rbac/#role-and-clusterole)