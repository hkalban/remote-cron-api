# Docker containers to use as test remote machines 

FROM ubuntu

USER root

RUN apt-get update && apt-get install sudo -y && apt-get install -y ssh
RUN sudo service ssh start

EXPOSE 22
