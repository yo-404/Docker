### Docker bind mounts and volumes

docker bind mounts are used to connect a storage directory from the host to the container . Usually the container mount directory and the directory selected within the host are the same . 

## Benefits of using Bind Mounts and volumes

Bind mounts can help user in retaining data even after the container is destroyed . Since we know that containers are ephemeral in nature and susceptible to issues like failures . In default way , with the containers the data is also lost whereas if we use bind mounts / volumes the data will be retained even after the container is stopped .

## Difference between bind mounts and volumes

Bind Mounts: Bind mounts can only be used in the host system i.e. the directory you select in the host computer will be assigned as directory for storage for the container .
Volumes: Volumes are more useful as when a volume is created , There is a logical partitioning of the disk and is assigned for the docker container . These volumes can be either locally created or globally mounted .

## Benefits of Volume 

Volumes are beneficial since these are isolated in nature and can enforce security regulations . Also there may be a case where the storage used in host system is not as per requirements (i.e. low I/O , storage , throughput etc) . In such cases it is advisable to use volume storage since volumes doesnt need to be in the local system .Volumes from cloud providers can also be used such as S3 bucket , EBS , NFS etc . By using volumes it is easier to create backup and to change the host provider if required , since the volumes are not binded from the host system .

## command for volume attachment

```
docker -v <parameters>
docker --mount

```
the --mount command is more verbose in nature

# Example

```
docker volume create yogi
docker volume ls
docker build -t example-mount .
docker run -d --mount source=yogi,target=/app nginx:latest

```

# to delete a volume 

```
docker volume rm yogi
```

<img src=https://github.com/yo-404/Docker/blob/main/docker-container-with-volume-mount/images/1.png?raw=true>

