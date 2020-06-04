# config

## 示例
```go
// 在 init 中初始化
config.Init(config.FromFile(),config.FromETCD())


var cfg interface{}
config.Unmarshal("abc.json",&cfg)
```