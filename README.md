## Steps for Local Development

Build Docker image
```bash
cd /go-health-tracker/scripts/deploy/helm
./build_system.sh
```

Deploy Backend, database and frontend
```bash
cd go-health-tracker
docker-compose up -d
```
### Port list
```bash
BACKEND SERVER  - 8332
FRONTEND SERVER - 3000
```

```
Accessing Server URLs

Go Backend: http://localhost:8332/hello
Go Frontend: http://localhost:3000/
```
| Server    | URL    |
|--------------|--------------|
| Go Backend  | http://localhost:8332/hello  | 
| Go Frontend  | http://localhost:3000/  |


### Notes for future:

1) Need to update the docker building part to create the UI image (Done)
    
    1.1) Docker compose for UI and Backend (Done)
2) Enhance the code on how to now introduce SQL database
3) May a docker compose file to have Backend, database and frontend
4) Once that is set we can start the Backend and, database developement
5) Once the APIs are development for basic for example
    5.1) /createuser
    5.2) /addweightentry
    5.3) /weightjourney
    5.4) /weighttrend
    5.5) /weightcomparison (B/w user-1 and user-2)

Design for GO HEALTH App: [Functional Spec link](https://mysidehussle.notion.site/Go-Health-Technical-Design-Document-14ddfe6efb398057a0c2f4a09c86284b)
