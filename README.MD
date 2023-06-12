### checker是一个为容器增加授权的工具
- 实现原理
  checker是一个后台进程，他会代理镜像的CMD命令，并且为容器添加授权检查功能
  例如Nginx容器，checker进程会代理原Nginx镜像执行 `nginx -g daemon off;` 启动nginx。在启动之前他会运行授权日期检查逻辑，只有检查通过后才会启动Nginx服务。
  另外，授权到期后checker进程会主动退出。
      
   
### build the project
[build-on-linux.sh](build-on-linux.sh) 
[build-on-mac.sh](build-on-mac.sh)