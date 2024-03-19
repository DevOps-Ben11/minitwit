# ITU MiniTwit WEB application
> This repository is used as a student project for the DevOps, Software Evolution and Software Maintenance course, at ITU. [Course description](https://learnit.itu.dk/local/coursebase/view.php?s=ft&view=public&ciid=642). Original files can be found on [this branch](https://github.com/DevOps-Ben11/minitwit/tree/Orignal-Minitwit)

This project contains a refactored Python Flask WEB application, which has been changed to Go.

**API:**
Expansive features added to the application, include an API which other systems can interact with(user registration, user following, message posting). 

**Public server:**
The system has also been deployed on a publicly accessible server, through digital ocean. This deployment has been automated to create a CI/CD pipeline. This pipeline is activated through github actions, when `main` is modified (typically through a pull request). The CI/CD also handles releases, and will publish a minor release every time the pipeline is activated. 

**Monitoring:**
Further more, a monitoring system has been added to the application, using prometheus and grafana.  

**Database migration:**
Finally, the database has been migrated from a SQLite database, running on the same droplet as the website, to a postgres PSQL database running on a seperate droplet.

## Topics 
- [Requirements](#requirements-⚙️)
- [Environments and Links](#environments-and-links-🔗)
- [Starting the project](#starting-the-project-🛠)
- [Deployment](#deployment)
- [API](#api-docs)
- [Vagrant](#vagrant)

## Requirements ⚙️
- go-lang
- psql
- Docker CLI
- vagrant (Optional, for restarting server)

## Environments and Links 🔗
- Public timeline:  [http://138.68.126.8/public](http://138.68.126.8/public)
- prometheus: [http://138.68.126.8:9090](http://138.68.126.8:9090)
- Grafana: [http://138.68.126.8:8080](http://138.68.126.8:8080)
- API link: `http://138.68.126.8:5000/sim/<Endpoint>`

## Starting the project 🛠
1. Firstly, pull the repository, and navigate to the minitwit folder:
``` 
git clone git@github.com:DevOps-Ben11/minitwit.git
```
```
cd minitwit
```
2. Start the PSQL database for local set-up: 
```
docker run --name my_postgres_db -e POSTGRES_PASSWORD=mysecretpassword -d -p 5431:5432 postgres
```
3. Navigate to the `backend` folder.
```
cd backend
``` 
4. Define the PSQL connection string in local environement: 
```
export PSQL_CON_STR="postgresql://postgres:mysecretpassword@localhost:5431/postgres"
```

5. Here you have 2 options for how to run the application. Run the go-lang file or create an image and run that. 
    1. **Go-lang** This can cause dependency issues however, to run it through go-lang, run: 
    ```
    go run main.go
    ```
    2. **Docker image** Alternativly, to avoid dependacy issues, you can build a docker image. and run, with:   
    ```
    docker build -t minitwit-image .
    ```
    Next, run the newly created image.
    ```
    docker run -v ./tmp:/app/tmp -p 5000:5000 minitwit-image
    ```
5. Open [http://localhost:5000/](http://localhost:5000/) to see the application

## Deployment
Deployments happen through CI/CD using GitHub Actions when merging to `main`.

The WEB application can be found here [http://138.68.126.8:5000](http://138.68.126.8:5000).

## API docs
Our API is documented through Postman and can be found [here](https://documenter.getpostman.com/view/1487273/2sA2rCU2He#intro). API for the **simulations** can be found there as well.


## Vagrant
To generalize the publication of the droplet on digitalOcean, Vagrant files have been used. Here, the procedure to start the server is documented. 
To stop the server. 
```
vagrant down
```

To start the server 
```
vagrant up
```

These two commands read the `.env` file (not published to git) and based on the variables, will publish and make a droplet on digital ocean. 
