# Intro
Hi there! This is a Go library that interacts with form3 api accounts resource that you can use in your project.

* Please keep in mind that this is my first interaction with Go lang.
* I started the project in nodejs, then switched to Go lang, so actually there are both libraries present in the repository. 

## Approach
Accountapi library is designed as a light passthrough layer that mediates between a client and a server. Specific types were created to serve as input and output, they are converted later to json. Any errors on server side are passed to the user to be handled, default http client settings are used (except for timeout and max conns per host). Input validation is left for the server side to be performed in order not to duplicate this functionality in the library. There are unit tests and integration test that cover library functionality and communication (focus here was not on a server logic).

## How should I run Go lib unittests and integration tests
To run all unit tests and integration tests please call:
```
docker-compose up --build --abort-on-container-exit
```
* unit tests are executed on docker build level
* integration tests are executed in a test container using this library, communicating with the fake api container

### How should I run Nodejs lib unittests and integration tests
```
docker-compose -f docker-compose.js.yml up --build --abort-on-container-exit
```

## Examples
There are few examples for lib usage in both Go lang and Nodejs in 'example' directory.
You can play with them locally after:
- running docker-compose command above without '--abort-on-container-exit' flag
- exporting ACCOUNT_API_LIB_HOST=http://localhost:8080

## Further steps
The library could be enhanced with configurable retry functionality, configurable timeout and other http params. Also interface abstration could be done, to allow client mocking the library easily.

## Task description and fakeapi container
https://github.com/form3tech-oss/interview-accountapi

## Resources
- https://gobyexample.com/
- https://golang.org/doc/tutorial/
- https://www.soberkoder.com/consume-rest-api-go/
- https://tutorialedge.net/golang/consuming-restful-api-with-go/
- https://github.com/googleapis/google-api-go-client
- https://github.com/jarcoal/httpmock
- https://pkg.go.dev/google.golang.org/api/blogger/v3
- https://blog.alexellis.io/golang-writing-unit-tests/
- https://olegcodes.medium.com/mocking-http-services-in-go-6b76215a81c9
- https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
- https://www.loginradius.com/blog/async/tune-the-go-http-client-for-high-performance/
- many others...
