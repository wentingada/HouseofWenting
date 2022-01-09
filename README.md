# HouseWorkofWenting
server.go. --20220109
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

 1.接收客户端 request，并将 request 中带的 header 写入 response header
 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
 4.当访问 localhost/healthz 时，应返回 200
 
multiplePC.go --20220109
基于 Channel 编写一个简单的多线程生产者消费者模型：

队列：
队列长度 10，队列元素类型为 int
生产者：
每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
