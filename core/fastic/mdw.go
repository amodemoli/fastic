// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package fastic

import (
	"github.com/valyala/fasthttp"
)

type Mdw struct {
	SChain func(handler fasthttp.RequestHandler, middlewares ...func(fasthttp.RequestHandler) fasthttp.RequestHandler) fasthttp.RequestHandler
	Chain  func(handler RequestHandler, middlewares ...func(RequestHandler) RequestHandler) RequestHandler
}

func MdwNew() *Mdw {
	return &Mdw{
		SChain: ChainMdw,
		Chain:  ChainCtx,
	}
}
