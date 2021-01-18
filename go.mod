module github.com/chinmobi/gin-mvc

go 1.15

require (
	github.com/chinmobi/ginmod v0.1.1
	github.com/chinmobi/modlib/evt v0.1.1
	github.com/chinmobi/modlib/grpool v0.1.0
	github.com/chinmobi/modlib/log v0.1.0
	github.com/gin-gonic/gin v1.6.3
	github.com/golobby/config v1.1.2
	github.com/stretchr/testify v1.7.0
	github.com/valyala/fasthttp v1.19.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
)

replace github.com/golobby/config v1.1.2 => github.com/chinmobi/config v1.2.1-alpha
