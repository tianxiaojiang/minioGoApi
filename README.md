
### 向 minio-server 访问图片的api

使用gin框架，实现上传、读取、删除三个接口

- [PUT]/:bucket
- [GET]/:bucket/:file
- [DELETE]/:bucket/:file


### 使用步骤

- 编译

在当前目录下执行命令```make```, 会将源码编译成二进制文件,位于 ```/Users/matin/projects/test/Docker/MinIO/service/api/bin```

- 启动服务
  
与 minio-server 在一个 docker-compose 中启动；开发环境，可以直接使用 docker-compose 来启动。

```/Users/matin/projects/test/Docker/MinIO```

