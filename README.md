# smartlab_backend_go

XDU 物理实验计算器 的 Golang 后端

## Swagger

/swagger/index.html

使用了 [gin-swagger](https://github.com/swaggo/gin-swagger) 库，每次 api 变更后应当手动更新。

## Config

将 config.yaml.example 命名为 config.yaml,并填入值

程序会优先读取 /config.xaml 再读取 /data/config.xaml