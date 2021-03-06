# Run on Ubuntu
FROM ubuntu:16.04

# Set config args
ARG config_file=./weaviate.conf.json 
ARG action_schema=./schema/test-action-schema.json
ARG thing_schema=./schema/test-thing-schema.json

# Set parametrase build args
ARG release=nightly
ARG platform=linux
ARG architecture=amd64

# Crearing the dir of weaviate
RUN mkdir -p /var/weaviate/config && cd /var/weaviate

# Install needed packages and scripts
RUN echo "BUILDING ${release}_${platform}_${architecture}.zip"

RUN apt-get -qq update && apt-get -qq install -y jq curl zip wget python-pip && \
    wget -q -O /var/weaviate/weaviate.zip "https://storage.googleapis.com/weaviate-dist/nightly/weaviate_${release}_${platform}_${architecture}.zip" && \
    cd /var/weaviate && unzip -o -q -j /var/weaviate/weaviate.zip && \
    rm /var/weaviate/weaviate.zip && \
    chmod +x /var/weaviate/weaviate && \
    pip install cqlsh
    
# Expose dgraph ports
EXPOSE 80

# Copying config files with using args
COPY $config_file /var/weaviate/
COPY $action_schema /var/weaviate/schema/
COPY $thing_schema /var/weaviate/schema/

# Copy script in container
COPY ./weaviate-entrypoint.sh /var/weaviate/weaviate-entrypoint.sh

# Set workdir
WORKDIR /var/weaviate/

# Run!
ENTRYPOINT ["./weaviate-entrypoint.sh"]
