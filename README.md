# ITU MiniTwit WEB application
http://64.226.68.241:8080/public

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