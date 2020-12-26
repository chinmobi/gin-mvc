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


## Authors

* **Zhaoping Yu** - *Initial work* - yuzhaoping1970@gmail.com

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
