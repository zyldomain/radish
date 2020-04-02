package pipeline

import (
	"errors"
	"radish/channel/iface"
)

type ChannelHandlerContext struct {
	next     *ChannelHandlerContext
	prev     *ChannelHandlerContext
	handler  iface.ChannelHandler
	pipeline iface.Pipeline
	inbound  bool
	outbound bool
}

func NewChannelHandlerContext(pipeline iface.Pipeline, handler iface.ChannelHandler) iface.ChannelHandlerContextInvoker {
	ctx := &ChannelHandlerContext{
		next:     nil,
		prev:     nil,
		handler:  handler,
		pipeline: pipeline,
		inbound:  false,
		outbound: false,
	}

	if _, ok := handler.(iface.ChannelInboundHandler); ok {
		ctx.inbound = true
	} else {
		ctx.outbound = true
	}

	return ctx
}

func (c *ChannelHandlerContext) Handler() iface.ChannelHandler {
	return c.handler
}

func (c *ChannelHandlerContext) Pipeline() iface.Pipeline {
	return c.pipeline
}

func (c *ChannelHandlerContext) Bind(address string) {
	c.invokeBind(c.findOutbound(), address)
}

func (c *ChannelHandlerContext) invokeBind(ctx iface.ChannelHandlerContextInvoker, address string) {
	handler, ok := ctx.Handler().(iface.ChannelOutboundHandler)
	if !ok {
		panic(errors.New("handler execute failed"))
	}

	handler.Bind(ctx, address)
}

func (c *ChannelHandlerContext) FireChannelActive(msg interface{}) {
	c.invokeChannelActive(c.findInbound(), msg)
}

func (c *ChannelHandlerContext) FireChannelUserEventTrigger(msg interface{}) {
	c.invokeChannelUserEventTrigger(c.findInbound(), msg)
}

func (c *ChannelHandlerContext) invokeChannelUserEventTrigger(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	handler, ok := ctx.Handler().(iface.ChannelInboundHandler)
	if !ok {
		panic(errors.New("handler execute failed"))
	}

	handler.UserEventTrigger(ctx, msg)
}

func (c *ChannelHandlerContext) invokeChannelActive(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	handler, ok := ctx.Handler().(iface.ChannelInboundHandler)

	if !ok {
		panic(errors.New("handler execute failed"))
	}
	handler.ChannelActive(ctx, msg)
}

func (c *ChannelHandlerContext) FireChannelInActive(msg interface{}) {
	c.invokeChannelActive(c.findInbound(), msg)
}

func (c *ChannelHandlerContext) invokeChannelInActive(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	handler, ok := ctx.Handler().(iface.ChannelInboundHandler)

	if !ok {
		panic(errors.New("handler execute failed"))
	}
	handler.ChannelInActive(ctx, msg)
}

func (c *ChannelHandlerContext) FireChannelRead(msg interface{}) {
	c.invokeChannelRead(c.findInbound(), msg)
}

func (c *ChannelHandlerContext) findInbound() iface.ChannelHandlerContextInvoker {
	for ctx := c.next; ctx != nil; ctx = ctx.next {
		if ctx.inbound {
			return ctx
		}
	}
	return nil
}

func (c *ChannelHandlerContext) invokeChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	handler, ok := ctx.Handler().(iface.ChannelInboundHandler)

	if !ok {
		panic(errors.New("handler execute failed"))
	}

	handler.ChannelRead(ctx, msg)
}

func (c *ChannelHandlerContext) Write(msg interface{}) {
	c.invokeChannelWrite(c.findOutbound(), msg)
}

func (c *ChannelHandlerContext) findOutbound() iface.ChannelHandlerContextInvoker {
	for ctx := c.prev; ctx != nil; ctx = ctx.prev {
		if ctx.outbound {
			return ctx
		}
	}
	return nil
}

func (c *ChannelHandlerContext) invokeChannelWrite(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	handler, ok := ctx.Handler().(iface.ChannelOutboundHandler)

	if !ok {
		panic(errors.New("handler execute failed"))
	}

	handler.Write(ctx, msg)
}

func (c *ChannelHandlerContext) RemoveSelf() {
	c.next.prev = c.prev
	c.prev.next = c.next
	c.prev = nil
	c.next = nil
}

func (c *ChannelHandlerContext) Channel() iface.Channel {
	return c.pipeline.Channel()
}
func (c *ChannelHandlerContext) Close() {
	c.invokeClose(c.findOutbound())
}

func (c *ChannelHandlerContext) invokeClose(ctx iface.ChannelHandlerContextInvoker) {
	handler, ok := ctx.Handler().(iface.ChannelOutboundHandler)

	if !ok {
		panic(errors.New("handler execute failed"))
	}

	handler.Close(ctx)
}
