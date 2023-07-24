# Docker Frequently asked Questions

### What is Docker

 Docker is an open source containerization platform . It enables developers to package applications into containers .Docker provides the ability to package and run an application in a loosely environment called container.Containers are lightweight and contain everything needed to run the application, so you do not need to rely on what is currently installed on the host .


### How containers are different from Virtual Machines

Containers run on a hosting platform / Host operating System which has the docker or any other containerization application installed . The applications are then installed on the container engine . Here each application is installed as a seperate container . The Container uses a lightweight software package that contains the required dependencies that are required to execture the contained application inside it. 

Pros
- they are very fast to modify and iterate on
- Robust ecosystem - Most runtime system offer a hosted public repo for pre made containers

Cons
- Shared host exploits - Containers usually share the same underlying hardware on the host operating system . Due to this, there may be a chance where exploit in one container could break out affect the other shared hardware and resources .


Virtual Machines 

Software called a hypervisor separates the machine’s resources from the hardware and provisions them appropriately so they can be used by the VM. These virtual machines emulate low levl hardware devices like CPU, disk and networking devices and may use OS to run seperate OS . Here each create virtual machine is isolated from one another . 

Pros
- Full isolation - All the VM environments are isolated from one another 

Cons
- Speed - can be time consuming to build and regenerate as it encompasses of a full OS . 
- Storage - Takes a lot of space .


### Docker Life-Cycle

A Docker lifecycle consists of the following sets 

- First of all the user would create a Dockerfile which encompasses the set of commands that is used to create image .
- Docker images are then created by running the docker the dockerfile
- Docker image is then executed in order to run the docker container
- User can also push /pull containers from the created container images 


### Different Components of Docker

- Docker client - Docker client is used to communicate between docker daemon and user via the docker CLI

- Docker daemon - Docker daemon is responisble for executing all the clients actions

- Docker Registeries - A docker registery is used to store docker images .It is usually a public registry that anyone can use to push/pull container images

- images - An image is a template with instruvtions for creating docker container .Usually  an image is based on another iamge with some additional customization .user can also create their won images and can publish it onto the docker registries.

Containers - Containers are light weight instances of an image .User can create ,start , stop ,move and delete a container using docker API or CLI .

### What is the difference between docker COPY and docker ADD

Docker ADD is used to copy the files from URL whereas Docker COPY is used to copy the files from the host system into the container


### What is the difference between CMD and Entrypoin in Docker

- CMD - Sets default parameters that can be overridden from the Docker command line interface (CLI) while running a docker container.

- ENTRYPOINT -  Sets default parameters that cannot be overridden while executing Docker containers with CLI parameters.


### What are the networking type in Docker and what is the default networking type

The default networking in Docker is Bridge 
Types of networking in docker are 
1. Bridge 
2. Overlay
3. Host
4. MacVlan
5. IPvlan

Bridge - In terms of Docker, a bridge network uses a software bridge which allows containers connected to the same bridge network to communicate, while providing isolation from containers which are not connected to that bridge network. The Docker bridge driver automatically installs rules in the host machine so that containers on different bridge networks cannot communicate directly with each other.

Overlay - The overlay network driver creates a distributed network among multiple Docker daemon hosts. This network sits on top of (overlays) the host-specific networks, allowing containers connected to it (including swarm service containers) to communicate securely when encryption is enabled. Docker transparently handles routing of each packet to and from the correct Docker daemon host and the correct destination container.

Host -If you use the host network mode for a container, that container’s network stack is not isolated from the Docker host (the container shares the host’s networking namespace), and the container does not get its own IP-address allocated. For instance, if you run a container which binds to port 80 and you use host networking, the container’s application is available on port 80 on the host’s IP address.

macVlan - Some applications, especially legacy applications or applications which monitor network traffic, expect to be directly connected to the physical network. In this type of situation, you can use the macvlan network driver to assign a MAC address to each container’s virtual network interface, making it appear to be a physical network interface directly connected to the physical network. In this case, you need to designate a physical interface on your Docker host to use for the Macvlan, as well as the subnet and gateway of the network. You can even isolate your Macvlan networks using different physical network interfaces.

IPvlan - The IPvlan driver gives users total control over both IPv4 and IPv6 addressing. The VLAN driver builds on top of that in giving operators complete control of layer 2 VLAN tagging and even IPvlan L3 routing for users interested in underlay network integration. For overlay deployments that abstract away physical constraints see the multi-host overlay driver.

