# radish

## 📖介绍

该项目仿照Netty的设计，使用Go语言实现的网络框架，通过与Golang自带的网络框架、gnet网络框架进行压测比较，该框架具有非常高的执行效率。目前支持TCP和UDP方式的NIO网络传输模型，包括客户端和服务端。非常欢迎大家Fork，一起完善该项目。

## 🎉快速使用

### 导入开发包

```shell
go get github.com/zyldomain/radish
```

### 服务端

#### 自定义服务端Handler

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



#### 服务端启动配置

##### 添加一个Handler

```go
	cg := channel.NewEpollEventGroup(4)
	pg := channel.NewEpollEventGroup(4)
	b := core.NewBootstrap().ParentGroup(pg).ChildGroup(cg).ChildHandler(&PrintHandler{})
	b.Bind("localhost:8080").Sync()
```

##### 添加多个Handler

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

	cg := loop.NewEpollEventGroup(num)
	pg := loop.NewEpollEventGroup(1)

	b := core.NewServerBootstrap().
		ServerSocketChannel(epoll.NIOServerSocket).
		NetWrok("tcp").
		ParentGroup(cg).
		ChildGroup(pg).
		ChildHandler(pipeline.NewChannelInitializer(
			func(pipeline iface.Pipeline) {
				pipeline.AddLast(&PrintHandler{})
				pipeline.AddLast(&ConvertHandler{})
			}))
	b.Bind("localhost:9001").Sync()
}

```



#### 执行效果

![image-20200319123927285](https://github.com/zyldomain/radish/blob/master/image-20200319123927285.png)



### 客户端

#### 自定义客户端Handler

```go
type ClientHandler struct {
	pipeline.ChannelInboundHandlerAdapter
}

func (p *ClientHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	b, ok := msg.([]byte)
	if !ok {
		ctx.Write([]byte("数据错误"))
		return
	}

	fmt.Println("服务端回传消息-> " + string(b))
	ctx.FireChannelRead(msg)
}
```

#### 客户端设置

```go
func main() {
	g := loop.NewEpollEventGroup(1)
	b := core.NewBootstrap().Group(g).Network("tcp").SocketChannel(epoll.NIOSocket).Handler(&ClientHandler{})

	b.Bind("localhost:9001")

	for i := 0; i < 10; i++ {
		b.Channel().Write([]byte("hello"))
	}
	select {}
}
```

