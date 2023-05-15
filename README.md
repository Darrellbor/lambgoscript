# Lambgoscript

Lambda, golang and typescript -- Comparing how typescript and golang fair with compute heavy workloads while showcasing the power of lambda to handle these processes

## How to run
All scripts are handled from within the `makefile`
some of the commands to get you started include
- `make run-sort` to run the golang sorting portion of the application
- `make build-sort` to build the golang sorting portion of the application and run a dry test
- `make run-fibonacci` to run the golang fibonacci portion of the application
- `make build-fibonacci` to build the golang fibonacci portion of the application and run a dry test
- `make test` to run tests on the golang portion of the application, both unit tests and race condition tests
- `make packageInstall` to install all the dependencies required for the typescript portion of the application
- `make startDev` to run the typescript portion of the application in development mode using nodemon
- `make buildTypescript` to build the typescript portion of the application 
- `make start` to first build the typescript application and then run the built app with node