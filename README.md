# skaffold Demo

## 初始代码

```sh
mkdir demo  && cd demo
go mod init github.com/costa92/skaffold-demo
```

## 初始 skaffold
```sh
skaffold init --generate-manifests
```

## 开发运行

```sh
skaffold dev
```

## 部署代码

```sh
skaffold run --tail
```

## debug

```sh
skaffold debug
```