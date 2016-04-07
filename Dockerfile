FROM alpine:latest

ADD swarmui /swarmui
ADD webui/build /webui/build

CMD ["/swarmui s &", "/swarmui a"]