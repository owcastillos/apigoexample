# apigoexample

[![go](https://img.shields.io/badge/go-v1.19.X-cyan.svg)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-v20.10.X-red.svg)](https://www.docker.com/)

>A simple Golang project to show how to implement a deployment on a K3S Server
>

## Prerequisites

You need the following things properly installed on your linux computer.

* [Git](http://git-scm.com/)
* [Go](https://golang.org/)
* [Docker](https://www.docker.com/)

Also, you need to configure the Docker repository and credentials and you need to have a K3S server

## Installation

Following you can find the instructions:

* `git clone https://github.com/owcastillos/apigoexample` this repository
* Change into the new directory `cd apigoexample`
* If you want to run locally, execute `go run main.go`
* Open `Dockerfile` and set your DB env variables
* Open `go-api-test.yml` and change image value `owcastillos/go-api-test:latest` by `{your_docker_user}/{your_docker_repo}:latest`
* Execute `sudo docker build --tag {your_docker_user}/{your_docker_repo} .` Don't forget last dot character
* If you want test last created image, you can run:
`sudo docker run --publish 8080:8080 {your_docker_user}/{your_docker_repo}` 
* Then, you must to execute `sudo docker push {your_docker_user}/{your_docker_repo}`
* When your image is published, you must go to your K3S server and create a `{your_file}.yml` file with `go-api-test.yml` content
* In your K3S server, you must execute `sudo k3s kubectl apply -f {your_file}.yml`
* If you execute `sudo k3s kubectl get ingress,svc,pods -n go-namespace`, you would find something like this:
```
NAME                  TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
service/go-api-test   ClusterIP   XX.XX.XXX.XX   <none>        80/TCP    2ms

NAME                               READY   STATUS    RESTARTS   AGE
pod/go-api-test-6d6b77fc46-zfqv9   1/1     Running   0          2m58s
```
* And finally, you can test your deployment
`curl -X GET http://XX.XX.XXX.XX`

## WARNING
You must not to generate an image from Mac, because this will not running on the K3S server