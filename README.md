Configuring API Gateway using HAProxy 
=================================

This repo was built just to test various features of HAProxy 
related to configuring API Gateway using docker.

What it does?
-----------------
After `docker-compose up` it will start 3 containers.
One for API Gateway (HAProxy) and two others for different imaginary services - OCR and SMS.

#### Service Containers
The service containers are running a tiny server (written in `api_server/server.go`).
It's listening on 8083 port and always responds with `SERVICE_NAME` and requesting URL.
This `SERVICE_NAME` is an env variable, set from `docker-compose.yml`.     

#### API Gateway Container 
It's based on official `haproxy:1.5`. 
It will use (COPY) configuration file at `haproxy/haproxy.cfg`.
We've also added some extra hosts entry for accessing services with meaningful domain names.  

In this repo, we are going to test -
> The API consumers will use virtual service endpoints which will be pointing to the API Gateway. 
Then the HAProxy will route them to appropriate service providers. 

How to test
--------------------
We will enter into the shell of gateway container. 
Then will make request using virtual domain names of services.
```sh
$ docker-compose exec gateway bash
root@abcdef123456:/# curl 'http://sms.services.com/send/sms'
Serving by sms_service for URL "/send/sms".
root@abcdef123456:/# curl 'http://ocr.services.com/extract/content'
Serving by ocr_service for URL "/extract/content".
```

That's all at this moment :smile: 

What's next?
--------------------
This repo contains only example of domain based **HTTP Routing**. 
But this setup is ready for testing any other features by - 
- Setting configurations in `haproxy/haproxy.cfg`
- Adding more services in `docker-compose.yml`   
 