## 预览打包后的效果 vue、react、uniapp

1、自行配制好`go`的运行环境

2、将打包后的文件全部丢到`dist`目录

3、可以安装`fresh`后使用命令`fresh`启动、更新文件会热更新、也可`go run main.go`

5、打包项目`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o uchat-linux main.go`

6、生成接口文档、先编写好`swag`注解、然后删除`docs`文件夹、执行命令`swag init`即可
