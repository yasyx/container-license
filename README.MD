### checker是一个为容器增加授权的工具
- 实现原理
  checker是一个后台进程，他会代理镜像的CMD命令，并且为容器添加授权检查功能
  例如Nginx容器，checker进程会代理原Nginx镜像执行 `nginx -g daemon off;` 启动nginx。在启动之前他会运行授权日期检查逻辑，只有检查通过后才会启动Nginx服务。
  另外，授权到期后checker进程会主动退出。
      
   
### build the project
docker build --build-arg="PROJECT_UUID=58f34c72-46d5-4fab-9666-b8530b0189ce" --build-arg="PROJECT_LICENSE_MONTH=36" . -t container/license:0.1