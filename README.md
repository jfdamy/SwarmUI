# swarmui

Swarm UI is a toolkit with an Api (REST), a WebUI and a autoscaling service for docker/docker swarm

I spent 3 hours on this project (i don't have a lot a spare time), it's a pet project.
So yes the code is ugly, yes i didn't use redux with react, yes the webui is ugly as hell.
But it's just the beginning so please be gentle and if something is too ugly for you, you can just contribute it ;)

#Run the api and the webui
- run npm run build in webui directory
- export SWARMUI_KV_STORE=etcd (etcd, consul, zk)
- export SWARMUI_KV_HOST="192.168.64.2:4001"
- swarmui need DOCKER env var to connect to the docker daemon or docker swarm
- go build
- ./swarmui serve


#screenshots

![Alt text](/doc/img/1.png?raw=true "List of project")
![Alt text](/doc/img/2.png?raw=true "Create a project (with docker compose file)")
![Alt text](/doc/img/3.png?raw=true "Manage a project")
![Alt text](/doc/img/4.png?raw=true "Edit a project (with docker compose file)")

#TODO

- Produce some Documentation
- Autoscaling service (based on cpu consumption like kubernetes or request per second on an LB like marathon-lb-autoscaling)
- Load Balancer for services
- Bus event for docker event
- A version of the api and the bus event in proto buf, thrift, ... (i don't know yet)
- Cleanup a litle bit the source code