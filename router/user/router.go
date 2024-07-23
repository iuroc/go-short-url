package user

import (
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/register", registerHandler)
	return router
}

/*
如何理解 ServeHTTP？

正常来说，通过 http.NewServeMux() 创建了一个路由，路由上面有很多条规则，
如果不将路由作为一个 Handler 传给 http.ListenAndServe() 或者 router.Handle()，他就不会被激活。

Handler 是一个包含了 ServeHTTP 方法的任意类型，ServeHTTP 是所属的 Handler 被触发时，执行的回调函数。

例如：

	router := http.NewServeMux()
	router.Handle("/xx", func(w, r) {})

这里传入的函数，相当于直接就是一个 ServeHTTP，只不过它还不是一个 Handler，
当 Handle() 执行时，会创建一个 Handler 然后把这个 ServeHTTP 放进去。

首先，在 http.ListenAndServe() 中，需要传入一个根 Handler，每一次的请求，都首先执行这个 Handler 里面的 ServeHTTP，
我们看 [http.ListenAndServe]() 的签名：

	func ListenAndServe(addr string, handler Handler) error

这里的 Handler 签名是：

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}

好，我们现在发来一个请求，我们就进入了这个 ServeHTTP，并且我们得到了 w 和 r。

而路由本身也是一个 Handler，那么它的 ServeHTTP 是怎么做的呢？

它先通过多条 router.Handle("/", func(w, r) {}) 来记录多条属于该 router 的规则，然后在执行 ServeHTTP 时，
会在里面去根据刚刚路由本身记录的那些规则去执行处理函数，或者不是函数而是新的 Handler，那就执行这个 Handler 里面的 ServeHTTP。

所以说，一个没有被 http.ListenAndServe() 或者 router.Handle 参数记录的 Handler，它的 ServeHTTP 函数是永远不会被执行的。

因为这些 Handler 永远不会被轮得到执行其中的 ServeHTTP 函数。

除非，我们手动触发它的 ServeHTTP。

所以我们现在创建 2 个 router，其中 routerA 记录了 N 条规则，而 routerB 只记录了一条规则，
我们现在只让 routerB 传入 http.ListenAndServe()，也就是说，routerA 的 ServeHTTP 永远不会被自动触发。

但是，我现在想要 routerB 能用 routerA 的能力去处理请求，
所以 routerB 的 ServeHTTP 执行时，先去找 routerB 挂载的规则，根据当前的 Path，找到了 routerB 的一条规则，比如 /，
我们 Path 可能是 /abc，但是用 / 依然可以匹配到，只需要我们的规则是 / 结尾的就行。

这个时候，我们把收到的 w 和 r 转发给 routerA 去处理就行了，我们手动去触发 routerA 的 ServeHTTP，然后把 w 和 r 传进去。

这样的话，routerA 的 ServerHTTP 就接收到了 w 和 r，就可以进行处理了。

总结，一个 Handler 只有 1 个 ServeHTTP，Handler 被触发，指的就是其 ServeHTTP 函数被触发，
并且 ServeHTTP 可以接收到来自 Handler 的触发者传来的 w 和 r，比如 http.ListenAndServe(address, handler)，
handler 里的 w 和 r 就是来自 http 的。

而 router.HandleFunc("/", func(w, r) {}) 里的 w 和 r 是来自当前 router 的。

那如何理解中间件的实现呢？

我们可以这么理解，那就是多一层代理，把上一层的 w 和 r 传给下一层的 ServeHTTP，
而不是直接将下一层 Handler 和代理层的 handler 一起挂载。

比如正常 router 根据其功能，命名为 userRouter，现在加上 ParseForm 校验，
那我们加一个函数：

	func ParseFormMiddleware(handler Handler) Handler {
		middleware := http.NewServeMux()
		middleware.HandleFunc("/", func (w, r) {
			if r.ParseForm() != nil {
				w.Write([]byte("参数错误"))
			} else {
				handler.ServeHTTP(w, r)
			}
		})
		return middleware
	}

所以，中间件就是接收一个 Handler，然后创建一个新的 Handler，
并在这个新的 Handler 的规则里面，条件性地触发接收到的 Handler 的 ServeHTTP 函数，
最后，返回这个新的 Handler。
所以，中间件是一个返回值为 Handler 的函数，它代理传入函数的 Handler 去执行其 ServeHTTP 函数。
*/
const M = 123