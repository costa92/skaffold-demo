# skaffold build

## 使用 skaffold buildpacks 构建镜像 [文档](https://skaffold-latest.firebaseapp.com/docs/pipeline-stages/builders/buildpacks/)

```yaml
# 设置构建配置  
build:
  platforms: ["linux/amd64", "linux/arm64"]
  artifacts:
    - image: costa92/go-mod-image
      buildpacks:
        builder: gcr.io/buildpacks/builder:v1
```

## 使用 skaffold docker 构建镜像 [文档](https://skaffold-latest.firebaseapp.com/docs/pipeline-stages/builders/docker/)

```yaml
# 设置构建配置
build:
  artifacts:
    - image: costa92/go-mod-image
      docker:
        cacheFrom: # 从哪些镜像缓存中拉取
          - gcr.io/distroless/static:nonroot
        dockerfile: Dockerfile
```
## 使用 skaffold kaniko 构建镜像 [文档](https://skaffold-latest.firebaseapp.com/docs/pipeline-stages/builders/ko/)

```yaml
# 设置构建配置  
build:
  platforms: ["linux/amd64", "linux/arm64"]
  artifacts:
    - image: costa92/go-mod-image
      ko:
        fromImage: gcr.io/distroless/base:debug-nonroot
```
