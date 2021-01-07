# Gin-MVC
The MVC framework based on [Gin](https://github.com/gin-gonic/gin).

The framework integrates MVC-developing patterns to build a RESTful web application (especially, the micro-service application).

The framework is lightweight and flexible, easy to extend and configure, and additionally with essential facilities to build an enterprise level project. Especially, the project is issued as out-of-box. You could clone the project, rename the project's module name, and then start your business-logic developing work.

The project acts as a starter to build a real RESTful, MVC-style web application.

The framework's core runtime is based on the [Gin](https://github.com/gin-gonic/gin), [fasthttp](https://github.com/valyala/fasthttp) and [ants](https://github.com/panjf2000/ants) to process millions of concurrent requests.

## Usage

Clone the project:

```sh
mkdir -p ~/work/go & cd ~/work/go
git clone git@github.com:chinmobi/gin-mvc.git <your-project-name>

```

Change the project's module name:

```sh
cd ~/work/go/<your-project-name>
./change-mod.sh <github.com>/<your-account-name>/<your-project-name>
```

Discard the useless `.git` directory:

```sh
cd ~/work/go/<your-project-name>
rm -rf .git
```

Then the project is ready for developing your own RESTful MVC-web application.
<br/>To test the works properly:

```sh
cd ~/work/go/<your-project-name>
make build
./_deploy/entrypoint.sh
```
The web server is listening at `8080` port as default, use `http://localhost:8080/` to access the web server.
<br/>Use `Ctrl+C` to exit the web server.
<br/>Use `make clean` to clean the built objects.
<br/>Use `make` to build more complex targets.

## Project structure

* `app` - The application main context to load and provide the app's components/utilities, it also manages the components/utilities' lifecycle (`Load, Setup/Start, Teardown/Shutdown, Release`).

* `config` - Load and set up the configurations that used to set up the app's components/utilities.

* `controller` - Define controller handers for handling requests.

* `ctx` - Provide request scale context while handling requests, especially for auth security.

* `db` - Manage the data-stores' connections (`redis`, `mongodb`, `prosgres`, etc...), and provide the DAOs(data access objects) that used for `model`.

* `errors` - Define the application scale errors.

* `evt` - Provide the application event utilities to publish / listen / multicast application events.

* `grpool` - Provide the goroutine pool that used for the app's other components.

* `log` - Provide the application logging utilities.

* `middleware` - Define middleware handlers before the `controller` handler while handling requests.

* `model` - Define the business entities and models that used for `service`.

* `mq` - Manage and provide message queueing utilities (`rabbitmq`, `kafka`, etc...).

* `restful` - Provide the RESTful APIs' utilities.

* `routes` - Map the `RESTful API` URL path routes to `controller` handers.

* `security` - Provide the [security](security/ABOUT.md) (authentication and authorization) utilities.

* `service` - Define the business services that used for `controller` and `middleware` handers.

* `tests` - Test cases for integration tests, especially for the `controller` layer about the RESTful APIs.

* `web` - Provide the web server to run the `app`.

## Best practices

1. First of all, analyze the requirements and business rules, and then write the use cases.

2. According to the use cases, design the [RESTful APIs](restful-apis.md) that served for the use cases.

3. Based on the [RESTful APIs](restful-apis.md), design `controller` handers, and DTOs(data transfer objects) used for the handlers.

4. Setup the `routes` that map the [RESTful APIs](restful-apis.md)' paths to `controller` handers.

5. According to the DDD's(domain-driven design) principle, design the business `model` entities and `service` interfaces to implement the `controller` handers.

6. Complete the `db` DAOs to implement the `model` and `service` objects.

**IMPORTANT:**

* Following the agile development. All of the steps can be and must be taking turns to develop, processing from iteration to iteration (step by step).

* Once someone has recognized some use cases, another one can design the APIs and/or design the business models/services simultaneously.

* Once someone has completed some APIs' design, another one can develop the controller handlers and/or develop the models/services and/or develop the DAOs simultaneously, and so on.

## Starting process

The whole starting process of the framework consists of three stages: `app-setup` stage, `web-setup` stage, and `start-run` stage.

* **The `app-setup` stage**

```
1. Load the configure (`config`) that will be used for other components.
2. Setup the logging (`log`).
3. Setup goroutine pool (`grpool`), the pool hasn't been activated.
4. Setup the application event's utility (`evt`).
5-1. Load the databases and cache (`db`).
5-2. Setup the models (`model`) using the database access objects (DAOs).
6. Load the message queuing utilities (`mq`).
7-1. Setup the services (`service`) using models (`model`).
7-2. Register event listeners and message queue consumers.
```

* **The `web-setup` stage**

```
1. Setup the middlewares (`middleware`) with the (`config`) and app services (`sevice`).
2. Setup controllers (`controller`) with app services (`sevice`).
3. Setup the router path routes (`routes`) with middlewares and controllers.
```

* **The `start-run` stage**

```
1. Activate the goroutine pool.
2. Start the http listening.
```

## Architecture

(coming soon...)

## About

Most of the design patterns and key concepts are refined from the [Spring Projects](https://spring.io/projects), the `Spring Framework`, `Spring MVC`, `Spring Security`, `Spring Data`, `Spring Boot`, etc.

As mature technology, the `Spring Projects` provides the best patterns and solutions to build enterprise level applications, especially the web, cloud, and micro-service products.

But the components and concepts about the `Spring` family technologies are huge, and much more flexible and configurable. It is hard to use the `Spring` technologies easily (especially for beginner), and the built projects are so heavy.

I (YuZhaoping) just refine the key thoughts of the `Spring` technologies to ensure the framework is light, but flexible and configurable to use, and more important, using the advanced features of `Golang` and `Gin`.

The project issuing as out-of-box style, is preferred to start the project immediately, not needed to do surrounding works to establish a real project. One aim of the project is to act as a starter using the best practices to build enterprise level applications.

## Authors

* **Zhaoping Yu** - *Initial work* - yuzhaoping1970@gmail.com

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
