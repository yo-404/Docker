### Docker Networking

Docker networking priciples can be enforced over the containers in order to let them communicate to each other or to stop them from communicating from each other . 

## Types of networking in Docker 

1. Bridge Networking
2. Host Networking
3. Overlay Networking

whenver a container is created , it comes with a by default eth0 network .The same goes for the host system ,it is also present with a eth0 network . The IPs of these can be different from each other and so in order to communicate there is a virtual network v.eth which helps the two entities to communicate with each other . By default when no network is chosed for the created container , it uses the default v.eth bridge to communicate between the host and the created containers . This creates a issue that whoever has access to the host network can access these containers as well .So , In order to get rid of this issue , the use can create his own secure-network which can distribute the containers inside the host and hence stops them from communicating from one another .

Networking in Docker servers both the use cases 
- stopping two containers from communicating from each other 
- letting two containers to communicate from one another 

Host Networking is known to be insecure in business usecase as whoever has access to the host will also have access to the container . If at all a malicious user gains access to your network ,he can affect your host system as well as your containers .


## Example

For EC2 instance
```
sudo apt update
sudo apt install docker.io -y
sudo apt update
systemctl status docker
sudo usemod -aG docker ubuntu
logout

```
login again


### to create a new network in docker

```
docker network create <network-name>
docker network ls
```

### to delete a network in docker

```
docker network rm <network-name>
```

### building a simple network example
```
docker pull nginx:lates
docker run -d --name login nginx:latest

docker inspect login
```
img 1

```
docker run -d --name logout nginx:latest
docker inspect logout
```
img2

```
docker exec -it login /bin/bash
apt-get install iputils-ping -y
//installed the iputils package to ping to the other container
ping 172.17.0.3
```
### creating a third container with custom network

```
docker network create secure-network
docker run -d --name finance --network=secure-network nginx:latest
docker inspect finance
```
img4



```
docker ps
docker exec -it login /bin/bash
ping 172.18.0.2
//pinging to finance container from login container
```
img6

Ping was unsuccessful because both the containers are in different network bridge


### To create a container with host network

```
docker run -d --name host-demo --network=host nginx:latest
docker inspect host-demo

```
img 7 

as we can see the container is not given IP address and is going to use the network configurations same as the host .