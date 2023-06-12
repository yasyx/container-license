
cp Dockerfile Dockerfile.tpl

APP_CONTAINER="nginx"
PROJECT_UUID="shdssss-asdfs-jjj"
PROJECT_LICENSE_MONTH=1
APP_CMD="nginx"
APP_ARGS="-g daemon off;"

sed -i  "s|APP_CONTAINER|${APP_CONTAINER}|g" Dockerfile.tpl
sed -i  "s|PROJECT_UUID|${PROJECT_UUID}|g" Dockerfile.tpl
sed -i  "s|PROJECT_LICENSE_MONTH|${PROJECT_LICENSE_MONTH}|g" Dockerfile.tpl
sed -i  "s|APP_CMD|${APP_CMD}|g" Dockerfile.tpl
sed -i  "s|APP_ARGS|${APP_ARGS}|g" Dockerfile.tpl

docker build -t chengzilong/license:0.1 -f Dockerfile.tpl .