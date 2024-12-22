# go-health-tracker

Hi Team, 

Please review the Go Health Design Document for more details link [https://mysidehussle.notion.site/Go-Health-Technical-Design-Document-14ddfe6efb398057a0c2f4a09c86284b]



## Steps for local backend development

# This will build the app and, create the go app backend docker image
cd /go-health-tracker/scripts/deploy/helm
./build_system.sh

# At this point if you are using docker desktop run the container with
# port enabled on 8332 

## Steps for local front development

cd go-health-tracker/ui/go-health-frontend/Dockerfile
docker build -f Dockerfile -t go-health-frontend-development .

# At this point if you are using docker desktop run the container with
# port enabled on 3000


Notes for future:

1) Need to update the docker building part to create the UI image 
2) Enhance the code on how to now introduce database to this
3) May a docker compose file to have Backend, database and frontend
4) Once that is set we can start the Backend and, database developement
5) Once the APIs are development for basic for example
    5.1) /createuser
    5.2) /addweightentry
    5.3) /weightjourney
    5.4) /weighttrend
    5.5) /weightcomparison (B/w user-1 and user-2)

