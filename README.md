# ITU MiniTwit WEB application
> This repository is used as a student project for the DevOps, Software Evolution and Software Maintenance course, at ITU. [Course description](https://learnit.itu.dk/local/coursebase/view.php?s=ft&view=public&ciid=642). Original files can be found on [this branch](https://github.com/DevOps-Ben11/minitwit/tree/Orignal-Minitwit)

This project contains a refactored Python Flask WEB application, which has been changed to Go.

**API:**
Expansive features added to the application, include an API which other systems can interact with(user registration, user following, message posting). 

**Public server:**
The system has also been deployed on a publicly accessible server, through Digital Ocean. This deployment has been automated to create a CI/CD pipeline. This pipeline is activated through github actions, when `main` is modified (typically through a pull request). The CI/CD also handles releases, and will publish a minor release every time the pipeline is activated.

**Monitoring:**
Further more, a monitoring system has been added to the application, using prometheus and grafana.  

**Database migration:**
Finally, the database has been migrated from a SQLite database, running on the same droplet as the website, to a PostgreSQL database running on a separate droplet.

## Topics 
- [ITU MiniTwit WEB application](#itu-minitwit-web-application)
  - [Topics](#topics)
  - [Requirements with Docker 🐳](#requirements-with-docker-)
  - [Requirements for Go and React ⚙️](#requirements-for-go-and-react-️)
  - [Environments and Links 🔗](#environments-and-links-)
  - [Starting the project 🛠](#starting-the-project-)
    - [With Docker](#with-docker)
    - [With Go and React](#with-go-and-react)
  - [Deployment](#deployment)

## Requirements with Docker 🐳
- Docker CLI

## Requirements for Go and React ⚙️
- go-lang
- psql
- node.js v20.11.1

## Environments and Links 🔗
- Public timeline:  [http://minitwit.fun/public](http://minitwit.fun/public)
- prometheus: [http://minitwit.fun:9090](http://minitwit.fun:9090/)
- Grafana: [http://minitwit.fun:8080](http://minitwit.fun:8080/)
  - username: admin
  - password: mrt3ukb5yvr@XFU*qgh
- API link: `http://159.89.4.152:5000/sim/<Endpoint>`

## Starting the project 🛠
1. Firstly, pull the repository, and navigate to the minitwit folder:
``` 
git clone git@github.com:DevOps-Ben11/minitwit.git
```
```
cd minitwit
```

2. Here you have 2 options for how to run the application. Run the go-lang file or create an image and run that:

### With Docker
  1. Start the PostgreSQL database:
  ```
  docker run --name my_postgres_db -e POSTGRES_PASSWORD=mysecretpassword -d -p 5431:5432 postgres
  ```

  2. Build the image:
  ```
  docker build -t minitwit-image . 
  ```

  3. Run the application:
  ```
  docker compose -f scripts/docker-compose.local.yml up
  ```

  4. Open [http://localhost:5000/](http://localhost:5000/) to see the application.

### With Go and React
  1. Navigate to the `backend` folder and create `.env` file based on `.env.example`:
  ```
  cd backend
  ```

  2. Start the server by running `go run main.go`
   
  3. To start the client on a different terminal go to the client folder:
  `cd client`
  and run `npm run dev`

  4. Open [http://localhost:3000/](http://localhost:3000/) to see the application.

> If all responses are 500, try clearing the browser cookies to regenerate new ones on the current encryption.

## Deployment
Deployments happen through CI/CD using GitHub Actions when merging to `main`.

The WEB application can be found here [http://minitwit.fun/](http://minitwit.fun/).