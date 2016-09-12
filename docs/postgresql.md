## Setup PostgreSQL

```
sudo apt-get install postgresql postgresql-contrib
sudo su - postgres
psql
createuser josh -s
\password josh (enter password: password123)
DROP DATABASE testdb; CREATE DATABASE testdb;
```
