### About

This is a simple two-endpoint echo/golang project that lets you request an IPv4 or Ipv6 and returns the location data.

#### Installation

To install the repo, please clone it, and install the necessary packages.
```sh
$ git clone https://github.com/domgolonka/kashoo-geo.git 
$ go mod download
```

**You will need to register for an API key. Please see [ipgeolocation.io](https://ipgeolocation.io/)**

Once you have your API Key from [ipgeolocation.io](https://ipgeolocation.io/), you will need to create a environment file in the root project directory.
```sh
$ APIKEY=YOUR_API_KEY
$ echo 'TOKEN='$APIKEY >> .env
```

Note: You will need to install Swagger-codegen. See [Swagger-codegen](https://swagger.io/tools/swagger-codegen/)

#### Building & Running locally

To build the app, swagger-codegen, & swagger-doc. Use the shell script. This will execute everything except for running the app.

```shell script
make
```

To Run the application:

```shell script
make run
```

##### Other useful commands
To build the application:

```shell script
make build
```
To build the Swagger-docs:

```shell script
swag-init
```
To build the Swagger-codegen:

```shell script
swag-code
```