# final_project-ftgo-h8

1. Dev
timothypartaliano@gmail.com
bimaputrasejati9999@gmail.com

2. Build image docker
docker build (folder) -t (nama image)

- rabbitmq
docker build -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management

- product service
docker build product-service -t fl-product-service:v1

docker tag fl-product-service:v1 gcr.io/copper-diorama-388207/fl-product-service:v1

docker push gcr.io/copper-diorama-388207/fl-product-service:v1

- notification service

docker build notification-service -t fl-notification-service:v1
docker tag fl-notification-service:v1 gcr.io/copper-diorama-388207/fl-notification-service:v1

- main api

docker build api -t fl-mainapi:v2
docker tag fl-mainapi:v2 gcr.io/copper-diorama-388207/fl-mainapi:v2

3. Login container registry gcp and push
gcloud auth login
docker push gcr.io/copper-diorama-388207/fl-mainapi:v1
docker push gcr.io/copper-diorama-388207/fl-notification-service
copper-diorama-388207
