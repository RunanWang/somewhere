FROM ubuntu

# RUN apt-get update  

# RUN apt-get install -y ca-certificates

# ENV TZ=Asia/Shanghai

# RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN mkdir -p /src/bin

RUN mkdir -p /src/conf

COPY main /src/bin

COPY ./conf/config.toml /src/conf/

WORKDIR /src

ENV PATH "$PATH:/src/bin"

