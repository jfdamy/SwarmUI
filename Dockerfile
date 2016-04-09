FROM alpine:latest

ADD SwarmUI /swarmui
ADD webui/build /webui/build

CMD ["/swarmui s &", "/swarmui a"]