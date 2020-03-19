package channel

import (
	"errors"
)

type ChannelHandlerContext struct {
	next     *ChannelHandlerContext
	prev     *ChannelHandlerContext
	handler  ChannelHandler
	pipeline Pipeline
	inbound  bool
	outbound bool
}

func NewChannelHandlerContext(pipeline Pipeline, handler ChannelHandler) *ChannelHandlerContext {
	ctx := &ChannelHandlerContext{
		next:     nil,
		prev:     nil,
		handler:  handler,
		pipeline: pipeline,
		inbound:  false,
		outbound: false,
	}

	if _, ok := handler.(ChannelInboundHandler); ok {
		ctx.inbound = true
	} else {
		ctx.outbound = true
	}

	return ctx
}

func (c *ChannelHandlerContext) Handler() ChannelHandler {
	return c.handler
}

func (c *ChannelHandlerContext) Pipeline() Pipeline {
	return c.pipeline
}

func (c *ChannelHandlerContext) Bind(address string) {
	c.invokeBind(c.findOutbound(), address)
}

func (c *ChannelHandlerContext) invokeBind(ctx *ChannelHandlerContext, address string) {
	handler, ok := ctx.handler.(ChannelOutboundHandler)
	if !ok {
		panic(errors.New("handler execute failed"))
	}

	handler.Bind(ctx, address)
}

func (c *ChannelHandlerContext) FireChannelActive(msg interface{}) {
	c.invokeChannelActive(c.findInbound(), msg)
}

func (c *ChannelHandlerContext) invokeChannelActive(ctx *ChannelHandlerContext, msg interface{}) {
	handler, ok := ctx.handler.(ChannelInboundHandler)

	if !ok {
		panic(errors.New("handler execute failed"))
	}
	handler.ChannelActive(ctx, msg)
}

func (c *ChannelHandlerContext) FireChannelInActive(msg interface{}) {
	c.invokeChannelActive(c.findInbound(), msg)
}

func (c *ChannelHandlerContext) invokeChannelInActive(ctx *ChannelHandlerContext, msg interface{}) {
	handler, ok := ctx.handler.(ChannelInboundHandler)

	if !ok {
		panic(errors.New("handler execute failed"))
	}
	handler.ChannelInActive(ctx, msg)
}

func (c *ChannelHandlerContext) FireChannelRead(msg interface{}) {
	c.invokeChannelRead(c.findInbound(), msg)
}

func (c *ChannelHandlerContext) findInbound() *ChannelHandlerContext {
	for ctx := c.next; ctx != nil; ctx = ctx.next {
		if ctx.inbound {
			return ctx
		}
	}
	return nil
}

func (c *ChannelHandlerContext) invokeChannelRead(ctx *ChannelHandlerContext, msg interface{}) {
	handler, ok := ctx.handler.(ChannelInboundHandler)

	if !ok {
		panic(errors.New("handler execute failed"))
	}

	handler.ChannelRead(ctx, msg)
}

func (c *ChannelHandlerContext) Write(msg interface{}) {
	c.invokeChannelWrite(c.findOutbound(), msg)
}

func (c *ChannelHandlerContext) findOutbound() *ChannelHandlerContext {
	for ctx := c.prev; ctx != nil; ctx = ctx.prev {
		if ctx.outbound {
			return ctx
		}
	}
	return nil
}

func (c *ChannelHandlerContext) invokeChannelWrite(ctx *ChannelHandlerContext, msg interface{}) {
	handler, ok := ctx.handler.(ChannelOutboundHandler)

	if !ok {
		panic(errors.New("handler execute failed"))
	}

	handler.Write(ctx, msg)
}
