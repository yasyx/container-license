### 实现背景
  单位的私有化部署项目在部署完成后，总是因为各种各样的原因出现客户无法回款的问题。因此需要在无侵入的情况下对关键应用进行授权管理。
### 前提
  私有部署项目完全基于docker部署。
### 实现原理
  container-license是一个后台进程，他会代理镜像的CMD命令，并且为容器添加授权检查功能
### checker
  后台的授权日期检查进程
### generate
  生成授权文件的进程
### 不足
  建议尽量对自定义镜像进行封装，如果是通用镜像，可以通过替换通用镜像的方式绕过授权检查。
### Demo介绍
checker是一个后台进程，他会代理镜像的CMD命令，并且为容器添加授权检查功能 例如Nginx容器，checker进程会代理原Nginx镜像执行 nginx -g daemon off; 启动nginx。在启动之前他会运行授权日期检查逻辑，只有检查通过后才会启动Nginx服务。 另外，授权到期后checker进程会主动退出。
### My blog
  https://meyasy.com 
### Build the project
[build-on-linux.sh](build-on-linux.sh) 
[build-on-mac.sh](build-on-mac.sh)
