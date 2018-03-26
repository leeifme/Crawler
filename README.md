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

### 单任务版爬虫架构

<img src = "http://orj2jcr7i.bkt.clouddn.com/%E5%8D%95%E4%BB%BB%E5%8A%A1%E7%89%88%E7%88%AC%E8%99%AB%E6%9E%B6%E6%9E%84.png" alt="单任务版爬虫架构">

### 并发版爬虫框架

<img src = "http://orj2jcr7i.bkt.clouddn.com/%E5%B9%B6%E5%8F%91%E7%89%88%E7%88%AC%E8%99%AB%E6%9E%B6%E6%9E%84.png" alt="单任务版爬虫架构">


#### Scheduler实现I 

> **简单调度器**
> 所有 `Worker` 共用一个输入

<img src = "http://on-img.com/chart_image/59a84c4ce4b02082b1db046a.png" alt="简单调度器">