
cp Dockerfile Dockerfile.tpl

APP_CONTAINER="swr.cn-north-4.myhuaweicloud.com/local-deploy/udesk_nginx:license-1.20.2"
PROJECT_UUID="96788b9e-e08c-483e-a4b5-ba3c15fc8e24"
PROJECT_LICENSE_MONTH=3
APP_CMD="nginx"
APP_ARGS="-g daemon off;"

sed -i  "s|APP_CONTAINER|${APP_CONTAINER}|g" Dockerfile.tpl
sed -i  "s|PROJECT_UUID|${PROJECT_UUID}|g" Dockerfile.tpl
sed -i  "s|PROJECT_LICENSE_MONTH|${PROJECT_LICENSE_MONTH}|g" Dockerfile.tpl
sed -i  "s|APP_CMD|${APP_CMD}|g" Dockerfile.tpl
sed -i  "s|APP_ARGS|${APP_ARGS}|g" Dockerfile.tpl

docker build -t swr.cn-north-4.myhuaweicloud.com/local-deploy/udesk_nginx:license-96788b9e -f Dockerfile.tpl .