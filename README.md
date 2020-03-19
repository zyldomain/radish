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

#### 添加一个Handler

```go
	cg := channel.NewEpollEventGroup(4)
	pg := channel.NewEpollEventGroup(4)
	b := core.NewBootstrap().ParentGroup(pg).ChildGroup(cg).ChildHandler(&PrintHandler{})
	b.Bind("localhost:8080").Sync()
```

#### 添加多个Handler

```go
type PrintHandler struct {
	channel.ChannelInboundHandlerAdapter
}

func (p *PrintHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	b, ok := msg.([]byte)

	if !ok {
		ctx.Write([]byte("数据错误"))
		return
	}

	fmt.Println("客户端发送消息-> " + string(b))
	ctx.FireChannelRead(msg)
}

type ConvertHandler struct {
	channel.ChannelInboundHandlerAdapter
}

func (p *ConvertHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	b, ok := msg.([]byte)

	if !ok {
		ctx.Write([]byte("数据错误"))
		return
	}

	ctx.Write([]byte("服务端收到消息-> " + string(b) + "\n"))

}

func main() {

	cg := channel.NewEpollEventGroup(4)
	pg := channel.NewEpollEventGroup(4)

	b := core.NewBootstrap().
		ParentGroup(pg).
		ChildGroup(cg).
		ChildHandler(channel.NewChannelInitializer(
			func(pipeline iface.Pipeline) {
				pipeline.AddLast(&PrintHandler{})
				pipeline.AddLast(&ConvertHandler{})
			}))
	b.Bind("localhost:8080").Sync()
}

```



### 执行效果

![image-20200319123927285](https://github.com/zyldomain/radish/blob/master/image-20200319123927285.png)
