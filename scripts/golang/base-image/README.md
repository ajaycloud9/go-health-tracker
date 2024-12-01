## Building base image with Golang and required tool to be used by all devs to runs their local development

Apline linux with Golang: hence the base is from here: https://hub.docker.com/_/golang


### 1) To Build the container:

```bash
cd go-health-tracker/scripts/golang/base-image
docker build -f Dockerfile -t go-health-tracker-base-image ./..
```
### 2) To Tag and Push the container
```bash
docker tag go-health-tracker-base-image:latest ajaysc30/go-health-tracker:go-base-image-1.22
docker push ajaysc30/go-health-tracker:go-base-image-1.22
```