# Gin-MVC
The MVC framework based on [Gin](https://github.com/gin-gonic/gin).

The framework integrates MVC-developing patterns to build a RESTful web application (especially, the microservice application).

The framework is lightweight and flexible, easy to extend and configure, and additionally with essential facilities to build an enterprise product project. Especially, the project is out of box. You could clone the project, rename the project's module name, and then start your business-logic developing work.

The project acts as a starter to build a real RESTful, MVC-style web application.

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

Then the project is ready to develop your own RESTful MVC-web application.
<br/>To test the works properly:

```sh
cd ~/work/go/<your-project-name>
go build -v
./<your-project-name>
```
The web server is listening at `8080` port as default, use `http://localhost:8080/` to access the web server.
<br/>Use `Ctrl+C` to exit the web server.
<br/>Use `go clean` to clean the built objects.
<br/>Use `make` to build more complex targets.

## Architecture

* `app` - The application main context to load and provide the app's components/utilities, it also manages the components/utilities' lifecycle (`Load, Setup/Start, Teardown/Shutdown, Release`).

* `config` - Load and set up the configurations that used to set up the app's components/utilities.

* `controller` - Define controller handers for handling requests.

* `ctx` - Provide request scale context while handling requests.

* `ctx/security` - Provide authentication and authorization while handling requests.

* `db` - Manage the data-stores' connections (`redis`, `mongodb`, `prosgres`, etc...), and provide the DAOs(data access objects) that used for `model`.

* `errors` - Define the application scale errors.

* `evt` - Provide the application event utilities to publish / listen / multicast application events.

* `grpool` - Provide the Goroutine pool that used for the app's other components.

* `log` - Provide the application logging utilities.

* `middleware` - Define middleware handlers before the `controller` handler while handling requests.

* `model` - Define the business entities and models that used for `service`.

* `mq` - Manage and provide message queueing utilities (`rabbitmq`, `kafka`, etc...).

* `routes` - Map the `RESTful API` URL path routes to `controller` handers.

* `service` - Define the business services that used for `controller` and `middleware` handers.

* `web` - Provide the web server to run the `app`.

## Best Practices

* First of all, analyze the requirements and business rules, and then write the use cases.

* According to the use cases, design the [RESTful APIs](restful-apis.md) that served for the use cases.

* Based on the RESTful APIs, design `controller` handers, and DTOs(data transfer objects) used for the handlers.

* Setup the `routes` that map the RESTful APIs' paths to `controller` handers.

* According to the DDD's(domain-driven design) principle, design the business `model` entities and `service` interfaces to implement the `controller` handers.

* Complete the `db` DAOs to implement the `model` and `service` objects.

## Authors

* **Zhaoping Yu** - *Initial work* - yuzhaoping1970@gmail.com

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
