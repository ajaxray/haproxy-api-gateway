FROM haproxy:1.5
# Install prerequisites
RUN apt-get update && apt-get install -y \
curl

COPY haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg