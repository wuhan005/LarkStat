# 📈 LarkStat [![Go Report Card](https://goreportcard.com/badge/github.com/wuhan005/LarkStat)](https://goreportcard.com/report/github.com/wuhan005/LarkStat) [![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?logo=sourcegraph)](https://sourcegraph.com/github.com/wuhan005/LarkStat)

飞书文档支持以卡片形式展示插入的链接，其后端将解析目标链接页面中的 `<meta>`
标签，提取页面的标题、描述、图片等信息。

目标链接的图片会在用户浏览器中跨域加载，因此我们可以动态生成每次加载的图片，以达到统计页面访问次数的效果。配合对 `Referer`
请求头的解析，可以得到具体的飞书租户子域名。

> [!WARNING]
> 注意：该项目功能很可能由于未来飞书调整页面 CORS 策略或业务逻辑后变得不再可用。届时该仓库将被归档。

## 在线体验

https://7ie-tech.feishu.cn/docx/Gbb7dyaAXodDXmxdSykc9qDvnZb

## 开始使用

替换 `https://larkstat.7ie.tech/<your-uri>` 中的 `<your-uri>`
为其他字符，即可获得一个新的统计链接。（如：`https://larkstat.7ie.tech/larkstat-test`
）将其插入到你的飞书文档，再将其配置为`卡片视图`展示即可。

## 私有部署

本项目使用 Go 开发，基于 https://github.com/syumai/workers 打包为 WASM 部署于 Cloudflare，并使用 Cloudflare KV 存储统计数据。

需要安装：

* Go 1.22+
* tinygo
* wrangler

在 Cloudflare 中创建名为 `larkstat` 的 KV 命名空间后，本地执行命令部署：

```bash
make deploy
```

## License

MIT License
