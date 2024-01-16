# Spotify-api

## Prerequisite

* Docker
* SQL Plus

Few important documents to install the above dependencies. It should be straight forward for windows.

Follow this [documentation](https://collabnix.com/how-to-run-oracle-database-in-a-docker-container-using-docker-compose/) for docker login.

Please follow this [documentation](https://oralytics.com/2022/09/22/running-oracle-database-on-docker-on-apple-m1-chip/) to connect to database after the docker login.

## Start Db

Add your own oracle password in the make file to ``ORACLE_PASSWORD`` variable. Add the same into ``.envrc`` file in ``DB_PASSWORD`` value

```sh
make start_db
```

Also run the content from [sql_scripts/create_table.sql](https://github.com/soham7222/spotify-api/blob/main/sql_scripts/create_table.sql) in the docker with sql plus to generate the table schema. 

## Run test locally with coverage

```sh
make test_local_with_coverage
```

## Run application locally

make sure to add `client_id` and `client_secret` in the `config.json` with your spotify creds.

Please run this from root directory. If you move any where else you have to change the varaibale values of the `constants.go`

```sh
make install_deps
make run_local
```
