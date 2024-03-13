# Migration plan

- Transfer the new docker-compose to the server

- Make sure the connection string is in the environment of the server (it should be)

- Make a pull request so we are ready to merge to decrease downtime

- Transfer data from sqlite to postgres on the server

  - This will also test that the connection string works and can connect to the postgres database

- Merge the pull request

- Pray 🙏
