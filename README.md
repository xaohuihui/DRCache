# DRCache
distributed-gRPC-cache

### Table of Contents

-----

- Background
- Install


#### Background
> 构建分布式缓存系统，具体框架如下
>> 1. 数据缓存使用redis，对外提供统一API查询接口 <br>
>> 2. 数据先进行本级节点查询，查询不到去远程节点获取，并缓存到本地 <br>
>> 3. 主节点与其他节点间通信使用 gRPC 服务 <br>
>> 4. 防止转发请求过多 使用singleflight 降低从节点收到主节点请求频率 <br>
>> 5. 使用一致性hash算法确定 key 对应到的节点上，可以本地缓存获取不到可以通过一致性hash计算确定转发的远程节点 <br>
