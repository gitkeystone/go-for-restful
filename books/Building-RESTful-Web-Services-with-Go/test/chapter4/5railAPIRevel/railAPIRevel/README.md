# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

   revel run myapp

### Go to http://localhost:9000/ and you'll see:

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file 允许我们设置主机、端口、开发模式/生产模式
        routes        Routes definition file 定义了端点、REST 动词和函数处理器的三元组（在这里，是控制器的函数）。这是组合路由、动词和函数处理器所必需的

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here 执行 API 逻辑的逻辑容器
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

