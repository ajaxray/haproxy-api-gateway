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
- The API consumers will use virtual service endpoints which will be pointing to the API Gateway. 
- Then the HAProxy will route them to appropriate service providers. 

How to test
--------------------
First get the repository and start the containers.
```shell script
git clone https://github.com/ajaxray/haproxy-api-gateway
cd haproxy-api-gateway
docker-compose build
docker-compose up -d
```

Now, we'll enter into the shell of gateway container. 
Then will make request using virtual domain names of services.
```shell script
$ docker-compose exec gateway bash
root@abcdef123456:/# curl 'http://sms.services.com/send/sms'
Serving by sms_service for URL "/send/sms".
root@abcdef123456:/# curl 'http://ocr.services.com/extract/content'
Serving by ocr_service for URL "/extract/content".
root@abcdef123456:/# exit
```
Don't forget to `docker-compose down` once testing is done.

That's all at this moment :smile:  
 
What's next?
--------------------
This repo contains only example of domain based **HTTP Routing**. 
But this setup is ready for testing any other features by - 
- Setting configurations in `haproxy/haproxy.cfg`
- Adding more services in `docker-compose.yml`   

Notes
-----
- If you modify `haproxy/haproxy.cfg`, you have to rebuild (`docker-compose build`).
- You DO NOT NEED to have golang installed to use this repository.
- But, if you modify `api_server/server.go`, you have to re-compile for appropriate architecture and OS
```shell script
cd api_server
GOOS=linux GOARCH=386 go build ./server.go
```
- You'll require Go 1.11 or higher to compile the `api_server/server.go`
 