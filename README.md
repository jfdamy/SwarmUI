[![GoDoc](https://godoc.org/github.com/jfdamy/SwarmUI?status.svg)](https://godoc.org/github.com/jfdamy/SwarmUI)
[![travis](https://travis-ci.org/jfdamy/SwarmUI.svg?branch=master)](https://travis-ci.org/jfdamy/SwarmUI.svg?branch=master)

# swarmui

Swarm UI is a toolkit with an Api (REST), a WebUI and a autoscaling service for Docker/Docker Swarm

The api use docker compose file to describe services for a project

Two autoscaling type : 
    - By node number of contianer of your service = number of node (nice for logging, monitoring, etc)
    - By CPU consumption with the limit to scale up and down (eg: 15% - 70%) and the min and max containers

I spent a few hours on this project (i don't have a lot a spare time), it's a pet project.
So yes the code is ugly, yes i didn't use redux with react, yes the webui is ugly as hell.
But it's just the beginning so please be gentle and if something is too ugly for you, you can just contribute it ;)


#screenshots

![Alt text](/doc/img/1.png?raw=true "List of project")
![Alt text](/doc/img/2.png?raw=true "Create a project (with docker compose file)")
![Alt text](/doc/img/3.png?raw=true "Manage a project")
![Alt text](/doc/img/4.png?raw=true "Edit a project (with docker compose file)")

#Run with docker
- launch etcd https://coreos.com/etcd/docs/latest/docker_guide.html 
- docker run -it -p 8080:8080 -e SWARMUI_KV_HOST="IP_ETCD:4001" -e SWARMUI_KV_STORE=etcd -v /var/run/docker.sock:/var/run/docker.sock jfdamy/swarmui /swarmui s
- docker run -it -e SWARMUI_KV_HOST="IP_ETCD:4001" -e SWARMUI_KV_STORE=etcd -v /var/run/docker.sock:/var/run/docker.sock jfdamy/swarmui /swarmui a

#Run the api and the webui
- run npm run build in webui directory
- export SWARMUI_KV_STORE=etcd (etcd, consul, zk)
- export SWARMUI_KV_HOST="192.168.64.2:4001"
- swarmui need DOCKER env var to connect to the docker daemon or docker swarm
- go build
- ./swarmui s for the REST API and WebUI
- ./swarmui a for the autoscaling service

#Dependencies
- Use libcompose to handle docker compose files of your projects
- Use etcd or consul or zk (use libkv of docker) to store docker compose files of your projects
- Docker or Docker Swarm obviously

#TODO

- Produce some Documentation
- Load Balancer for services
- Job scheduler
- A version of the api and the bus event in proto buf, thrift, ... (i don't know yet)
- Cleanup a litle bit the source code
