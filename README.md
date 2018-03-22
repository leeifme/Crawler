## Go分布式爬虫
> 处理和分发爬取珍爱网用户信息的数据任务,并提供关键字收索以及简单前端页面数据展示

- 基于 `Go` 语言标准库开发网络爬虫,不使用现有的爬虫库/框架
- 使用 `Docker` 容器运行 `ElasticSearch` 作为数据存储以及全文收索引擎
- 使用 `Go` 语言标准模板库实现http数据展示部分
- 爬虫框架由单任务版过渡到并发版(多个 `Goroutine}` ),直至分布式爬取数据

### 爬虫总体算法

#### 解析器<Parser>

- 输入：`utf-8` 编码的文本
- 输出：`Request{URL，对应Parser}` 列表，`Item` 列表

<img src = "http://orj2jcr7i.bkt.clouddn.com/Parser.png" alt="爬虫总体算法框架">