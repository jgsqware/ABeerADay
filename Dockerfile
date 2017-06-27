FROM ubuntu:16.10
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*
RUN curl -O https://ftp.openssl.org/source/old/1.0.1/openssl-1.0.1.tar.gz
RUN tar xvf openssl-1.0.1.tar.gz
RUN ls
COPY a-beer-a-day /
EXPOSE 8080
ENTRYPOINT ["/a-beer-a-day"]

