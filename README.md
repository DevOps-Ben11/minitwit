# ITU MiniTwit WEB application

Refactored Python Flask WEB application to Go.

## Requirements ⚙️

- Docker CLI

## Starting the project 🛠

In your prefered terminal do:

1. `git clone git@github.com:DevOps-Ben11/minitwit.git`
2. `cd minitwit`
3. `docker build -t minitwit .`
4. `docker run -p 8080:8080 minitwit`
5. Open [http://localhost:8080/](http://localhost:8080/) to see the application

## Deploying
The WEB application can be found here [http://64.226.68.241:8080/public](http://64.226.68.241:8080/public).

1. Create `.env` file based on `.env.example`.
2. Running `vagrant up` will create a new droplet on DigitalOcean which will install docker and build the GO application
3. To redeploy, run `vagrant provision` which will re-sync the local folder with VPS and rerun docker commands.
