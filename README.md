# Go Short URL

Go 语言实现的短网址平台。

## 关于 `vanjs-router` 的改进建议

- 支持延时加载路由，将需要加载的 () => Element 作为参数存储在 Route 对象上，然后通过 route.load() 载入 Element

  ```ts
  // 如果启用 delayed，则需要显示调用 show() 方法，否则就算命中路由，也不会显示页面。
  new Handler({
    name: "login",
    delayed: true,
    loader() {
      return div("Hello World");
    },
    async onFirst() {
      await new Promise();
      this.show();
    },
    async onLoad() {},
  });
  // 页面加载器，会被作为路由的根元素
  handler.loader = () => div("Hello World");
  // show() 方法必须确保当前路由是命中状态，才能显示，因为可能 Promise 启动时是路由命中状态，但是当 Promise 结束时，命中状态可能中途就不再维持。
  // onLoad 和 onFirst 都必须是 Promise
  // 必须要 await onFirst() 再 await onLoad()
  ```
