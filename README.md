# ITU MiniTwit WEB application

Refactored Python Flask WEB application to Go.

## Requirements ⚙️

- Docker CLI

## Starting the project 🛠

In your prefered terminal do:

1. `git clone git@github.com:DevOps-Ben11/minitwit.git`
2. `cd minitwit`
3. `docker build -t minitwit-image .`
4. `docker run -p 5000:5000 minitwit-image`
5. Open [http://localhost:5000/](http://localhost:5000/) to see the application

## Deploying
Deployments happen through CI/CD using GitHub Actions when merging to `main`.

The WEB application can be found here [http://138.68.126.8:8080](http://138.68.126.8:5000).

## API docs
Our API is documented through Postman and can be found [here](https://documenter.getpostman.com/view/1487273/2sA2rCU2He#intro). API for the **simulations** can be found there as well.