# 分页查询自己账户下的域名列表文档示例

该项目为调用QueryDomainList分页查询自己账户下的域名列表。文档示例，该示例**无法在线调试**，如需调试可下载到本地后替换 [AK](https://usercenter.console.aliyun.com/#/manage/ak) 以及参数后进行调试。

## 运行条件

- 下载并解压需要语言的代码;

- 在阿里云帐户中获取您的 [凭证](https://usercenter.console.aliyun.com/#/manage/ak)并通过它替换下载后代码中的 ACCESS_KEY_ID 以及 ACCESS_KEY_SECRET;

- 执行对应语言的构建及运行语句

## 执行步骤
下载的代码包，在根据自己需要更改代码中的参数和 AK 以后，可以在**解压代码所在目录下**按如下的步骤执行

- Java
*JDK 版本要求 1.8 及以上*
```sh
mvn clean package
java -jar target/sample-1.0.0-jar-with-dependencies.jar
```

- Java 异步
*JDK 版本要求 1.8 及以上*
```sh
mvn clean install compile
mvn exec:java -Dexec.mainClass=demo.${apiName} -Dexec.cleanupDaemonThreads=false
```

- TypeScript
*Node 版本要求 10.x 及以上*
```sh
npm install --registry=https://registry.npm.taobao.org && tsc && node ./dist/client.js
```

- Go
*Golang 版本要求 1.13 及以上*
```sh
GOPROXY=https://goproxy.cn,direct go run ./main
```

- PHP
*PHP 版本要求 7.2 及以上*
```sh
composer install && php src/Sample.php
```

- Python
*Python 版本要求 Python3*
```sh
python3 setup.py install && python ./alibabacloud_sample/sample.py
```

- C#
*.NETCORE 版本要求 2.1及以上*
```sh
cd ./core && dotnet run
```

## 使用的 API

-  QueryDomainList 调用QueryDomainList分页查询自己账户下的域名列表。文档示例，可以参考：[文档](https://next.api.aliyun.com/document/Domain/2018-01-29/QueryDomainList)



## 返回示例

*实际输出结构可能稍有不同，属于正常返回；下列输出值仅作为参考，以实际调用为准*


- JSON 格式 
```js
{
  "Data": {
    "Domain": [
      {
        "ExpirationDate": "Nov 02,2019 04:00:45",
        "InstanceId": "ST2015110212003800001928",
        "RegistrationDate": "Nov 02,2017 04:00:45",
        "DomainName": "fds234sdf.net",
        "DomainType": "gTLD"
      }
    ]
  },
  "TotalItemNum": 1,
  "PageSize": 5,
  "CurrentPageNum": 0,
  "RequestId": "77F90DA6-89AB-4074-80F3-E595CB57DF98",
  "PrePage": false,
  "TotalPageNum": 1,
  "NextPage": false
}
```
- XML 格式 
```xml
<?xml version='1.0' encoding='UTF-8'?>
<QueryDomainListResponse>
    <Data>
        <Domain>
            <ExpirationDate>Nov 02,2019 04:00:45</expirationDate>
            <InstanceId>ST20151102120031118</SaleId>
            <RegistrationDate>Nov 02,2017 04:00:45</registrationDate>
            <DomainName>test.com</DomainName>
            <DomainType>gTLD</domainType>
        </Domain>
    </Data>
    <TotalItemNum>1</TotalItemNum>
    <PageSize>5</PageSize>
    <CurrentPageNum>0</CurrentPageNum>
    <RequestId>B7AB5469-5E38-4AA9-A920-C65B7A9C8E6E</RequestId>
    <PrePage>false</PrePage>
    <TotalPageNum>1</TotalPageNum>
    <NextPage>false</NextPage>
</QueryDomainListResponse>
```


