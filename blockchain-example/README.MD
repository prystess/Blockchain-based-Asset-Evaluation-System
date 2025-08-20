## 示例工程使用指南

#### 1、获取网络配置文件
    在智能合约页签中下载连接配置文件network.json，拷贝到项目resources目录下
    
#### 2、申请客户端证书
    在区块链平台申请客户端证书，下载证书和私钥，拷贝到项目resources目录下
    
#### 3、直接上链
    找到com.yonyou.example.main.ConnectBlockchainTest启动类，修改第30行和第31行的文件名
    
#### 4、以服务形式启动
    找到com.yonyou.example.service.BlockchainOperateService类
    修改第41行和第42行的文件名，然后找到BlockchainExampleApplication类，启动，默认使用80端口
    接口调用请查看blockController类