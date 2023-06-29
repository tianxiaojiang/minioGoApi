
### 向 minio-server 访问图片的api

使用gin框架，实现上传、读取、删除、打标签四个接口

- [PUT]/:bucket
- [GET]/:bucket/:file
- [DELETE]/:bucket/:file
- [PTCH]/:bucket/:file

调用说明：

上传：{{host}}/{{bucketName}}, contentType:form-data, 上传文件名: upload    
读取： {{host}}/{{bucketName}}{{fileName}}  
删除： {{host}}/{{bucketName}}{{fileName}}  
打标签： {{host}}/{{bucketName}}{{fileName}}?tk={{tagName}}&tv={{tagValue}}    

### 使用步骤

- 编译说明

在当前目录下执行命令```make```, 会将源码编译成二进制文件,位于 ```/Users/matin/projects/test/Docker/MinIO/service/api/bin```

- 启动服务
  
与 minio-server 在一个 docker-compose 中启动；开发环境，可以直接使用 docker-compose 来启动。

```/Users/matin/projects/test/Docker/MinIO```

