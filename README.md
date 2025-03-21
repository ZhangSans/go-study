# go练习项目

## 1，[牛客题霸-go入门到实践四十招](./牛客题霸-go入门到实践四十招)

* 功能：牛客语法题
* 涉及知识点：基础语法知识

## 2，[猜数字游戏](./猜数字游戏)

* 功能：开发一个猜数字命令行游戏。
* 涉及知识点：循环、条件判断、随机数生成。

## 3，[天气查询工具](./天气预报)

* 功能：通过调用天气API，查询并显示当前天气。
* 涉及知识点：HTTP请求、JSON解析、API调用。

## 3，[简易词典](./简易词典)

* 功能：通过命令行输入单词，调用翻译 API 返回翻译结果。
* 涉及知识点：net/http 包、json 包、goroutine 并发调用 API。





1. 命令行工具
* 功能：开发一个简单的命令行工具，如待办事项管理器、文件重命名工具等。
* 涉及知识点：flag包、文件操作、字符串处理。
2. Web服务器
* 功能：创建一个基础的HTTP服务器，处理GET/POST请求，返回静态文件或JSON数据。
* 涉及知识点：net/http包、路由处理、JSON编解码。
3. RESTful API
* 功能：开发一个简单的RESTful API，支持CRUD操作，数据可以存储在内存或文件中。
* 涉及知识点：HTTP方法、路由、JSON编解码、错误处理。
4. 爬虫
* 功能：编写一个爬虫，抓取网页内容并提取特定信息，如新闻标题或图片链接。
* 涉及知识点：net/http包、HTML解析（如goquery）、正则表达式。
5. 聊天应用
* 功能：实现一个简单的命令行聊天应用，支持多用户聊天。
* 涉及知识点：goroutine、channel、TCP/UDP通信。
6. 文件同步工具
* 功能：开发一个工具，同步两个文件夹的内容。
* 涉及知识点：文件操作、递归遍历、并发处理。
7. 密码管理器
* 功能：创建一个命令行密码管理器，支持密码的加密存储和检索。
* 涉及知识点：文件操作、加密解密、命令行参数解析。
8. 简单的游戏
* 功能：开发一个命令行游戏，如猜数字、井字棋等。
* 涉及知识点：循环、条件判断、随机数生成。
9. 日志分析工具
* 功能：编写一个工具，分析日志文件并生成统计信息，如错误次数、访问量等。
* 涉及知识点：文件操作、字符串处理、正则表达式。
10. 简单的数据库操作
* 功能：使用Go操作SQLite或MySQL数据库，实现增删改查功能。
* 涉及知识点：database/sql包、SQL查询、错误处理。
11. 并发下载器
* 功能：开发一个支持并发下载文件的工具。
* 涉及知识点：goroutine、channel、HTTP请求。
12. 天气查询工具
* 功能：通过调用天气API，查询并显示当前天气。
* 涉及知识点：HTTP请求、JSON解析、API调用。
13. 简单的博客系统
* 功能：创建一个命令行或Web端的博客系统，支持文章的发布、编辑和查看。
* 涉及知识点：文件操作、Web开发、数据库操作。
14. URL短链生成器
* 功能：开发一个工具，将长URL转换为短链并存储。
* 涉及知识点：字符串处理、文件操作、Web开发。
15. 简单的监控工具
* 功能：编写一个工具，监控服务器的CPU、内存使用情况并生成报告。
* 涉及知识点：系统信息获取、文件操作、定时任务。
16. Markdown解析器
* 功能：开发一个简单的Markdown解析器，将Markdown文本转换为HTML。
* 涉及知识点：字符串处理、正则表达式、文件操作。
17. 简单的邮件客户端
* 功能：编写一个命令行工具，发送和接收邮件。
* 涉及知识点：SMTP/IMAP协议、邮件处理。
18. 任务调度器
* 功能：开发一个简单的任务调度器，支持定时执行任务。
* 涉及知识点：time包、并发处理。
19. 简单的缓存系统
* 功能：实现一个基于内存的缓存系统，支持设置、获取和删除缓存项。
* 涉及知识点：map、并发控制、sync包。
20. 简单的编译器或解释器
* 功能：编写一个简单的编译器或解释器，解析并执行简单的脚本语言。
* 涉及知识点：词法分析、语法分析、解释执行。

1. 并发聊天室
功能：实现一个简单的并发聊天室，支持用户上下线广播、私聊、用户改名、超时强踢等功能。

技术点：goroutine、channel、net 包、map 数据结构。

代码量：约 100 行。

学习价值：理解 Go 的并发模型和网络编程基础。

参考项目：[100 行 Go 代码实现并发聊天室]1011。
https://www.jb51.net/article/168041.htm
https://blog.csdn.net/weixin_42940826/article/details/82386275


2. Git HTTP 服务器
功能：用 Go 实现一个简单的 Git HTTP 服务器，支持通过 HTTP 协议访问 Git 仓库。

技术点：net/http 包、os/exec 包、CGI 协议。

代码量：约 100 行。

学习价值：了解 Git 的工作原理和 Go 的 Web 编程。

参考项目：[100 行 Go 代码实现 Git HTTP 服务器]6。
https://blog.csdn.net/ruzhila/article/details/139848816


5. 简易代理服务器
功能：实现一个基于 Socks5 协议的简易代理服务器，支持客户端通过代理访问目标服务器。

技术点：net 包、goroutine、Socks5 协议。

代码量：约 100 行。

学习价值：理解网络协议和代理服务器的工作原理。

参考项目：[简易代理服务器]15。
https://www.cnblogs.com/yltfy1998/p/16261276.html

6. 静态网站生成器
功能：读取 Markdown 文件，生成静态 HTML 页面。

技术点：os 包、html/template 包、文件操作。

代码量：约 100 行。

学习价值：学习文件操作和模板渲染。

参考项目：[Hugo 静态网站生成器]14。
https://www.magedu.com/wenzhang/golang/35477.html
7. 任务管理器
功能：实现一个简单的命令行任务管理器，支持添加、查看、删除任务。

技术点：slice、map、命令行参数解析。

代码量：约 100 行。

学习价值：掌握 Go 的数据结构和命令行工具开发。

参考项目：[任务管理器]7。
https://news.qq.com/rain/a/20250204A05A6000