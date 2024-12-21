## Building base image with Golang and required tool to be used by all devs to runs their local development

Apline image: Lightweight Image for alpine: https://hub.docker.com/_/alpine


### 1) To Build the container:

```bash
cd go-health-tracker/scripts/golang/alpine-image
docker build -f Dockerfile -t alpine-image ./..
```
### 2) To Tag and Push the container
```bash
docker tag alpine-image:latest ajaysc30/go-health-tracker:alpine-3.14
docker push ajaysc30/go-health-tracker:alpine-3.14
```