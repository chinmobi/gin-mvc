## NOTE

* The app uses the `PORT` env for web server listening, the port's default value is `8080`. <br/> If you want to change the default value, modify the following files: `./.env`, `./Dockerfile`, `./docker/.env`.

* The docker deployment uses the `gin-mvc` as the `APP_NAME`, <br/> If you want to modify the default value, modify the following files: `./.env`, `./docker/.env`.

* The docker deployment uses the `chinmobi` as the `APP_USER` (also as the account name), <br/> If you want to modify the default value, modify the following files: `./Dockerfile`, `./docker/.env`.

* The `./docker/.env` file defines the docker image tag. If you upgrade your project version, modify the value along with that.

* The `./docker/.env` file also defines the docker network names used for docker deployment: `app-back-net`, `app-front-net`.
