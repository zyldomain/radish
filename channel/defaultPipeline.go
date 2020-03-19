package channel

import "radish/channel/iface"

type DefaultChannelPipeline struct {
	head    *ChannelHandlerContext
	tail    *ChannelHandlerContext
	channel iface.Channel
}

func NewDefaultChannelPipeline(channel iface.Channel) *DefaultChannelPipeline {

	p := &DefaultChannelPipeline{
		head:    nil,
		tail:    nil,
		channel: channel,
	}

	p.addLast0(NewHeadContext(p))
	p.addLast0(NewTailContext(p))
	return p
}

func (d *DefaultChannelPipeline) AddLast(handler iface.ChannelHandler) iface.Pipeline {
	ctx := NewChannelHandlerContext(d, handler)
	d.addLast0(ctx)

	handler.ChannelHandlerAdded(ctx)

	return d
}

func (d *DefaultChannelPipeline) addLast0(c iface.ChannelHandlerContextInvoker) {

	ctx, _ := c.(*ChannelHandlerContext)
	if d.head == nil {
		d.head = ctx
		d.tail = ctx
	} else if d.tail == d.head {
		d.tail.next = ctx
		ctx.prev = d.tail
		d.tail = ctx
	} else {
		d.tail.prev.next = ctx
		ctx.prev = d.tail.prev

		ctx.next = d.tail
		d.tail.prev = ctx
	}

}

func (d *DefaultChannelPipeline) Tail() iface.ChannelHandlerContextInvoker {
	return d.tail
}
func (d *DefaultChannelPipeline) Channel() iface.Channel {
	return d.channel
}
func (d *DefaultChannelPipeline) ChannelRead(msg interface{}) {
	d.head.FireChannelRead(msg)
}
func (d *DefaultChannelPipeline) Write(msg interface{}) {
	d.tail.Write(msg)
}
func (d *DefaultChannelPipeline) Bind(address string) {
	d.tail.Bind(address)
}
func (d *DefaultChannelPipeline) ChannelActive(msg interface{}) {
	d.head.FireChannelActive(msg)
}
func (d *DefaultChannelPipeline) ChannelInActive(msg interface{}) {
	d.head.FireChannelInActive(msg)
}
