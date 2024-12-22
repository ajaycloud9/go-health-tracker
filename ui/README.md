docker run -it -p 3000:3000 -v /Users/ajaysingh/ws/go-health-tracker/ui:/go-health-frontend/ui/ ui-base

docker run -d --name frontend --network go-health-app-network -p 3000:3000 frontend-image
docker run -d --name backend --network go-health-app-network -p 8000:8000 backend-image

docker run -it -p --network host -v /Users/ajaysingh/ws/go-health-tracker/ui:/go-health-frontend/ui/ ui-base

FROM ajaysc30/go-health-tracker:ui-base-22.11.0 as builder
RUN mkdir -p /go-health-frontend

COPY  ui/ /go-health-frontend/ui/
WORKDIR /go-health-frontend/ui/

# install
RUN npm install
# build for standalone version
RUN npm run build

# For deployment
WORKDIR /go-health-frontend/ui/scripts/server
RUN npm install

FROM ajaysc30/go-health-tracker:ui-base-22.11.0

COPY --from=builder /go-health-frontend/ui/dist /public
COPY --from=builder /go-health-frontend/ui/scripts/server /
# COPY --from=builder /nvo/uinvo/documentation /documentation
EXPOSE 8087 
