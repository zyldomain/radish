# radish

## 快速使用

### 导入开发包

```shell
go get github.com/zyldomain/radish
```

### 自定义服务端Handler

```go
type PrintHandler struct {
	channel.ChannelInboundHandlerAdapter
}

func (p *PrintHandler) ChannelRead(ctx *channel.ChannelHandlerContext, msg interface{}) {
	b, ok := msg.([]byte)

	if !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(string(b))
	}

	ctx.Write([]byte("服务端的消息-> " + string(b) + "\n"))
}
```



### 开启服务端

```go
	cg := channel.NewEpollEventGroup(4)
	pg := channel.NewEpollEventGroup(4)

	b := core.NewBootstrap().ParentGroup(pg).ChildGroup(cg).ChildHandler(&PrintHandler{})
	b.Bind("localhost:8080").Sync()
```

### 执行效果

![image-20200319123927285](http://github.com/zyldomain/radish/image-20200319123927285.png)
