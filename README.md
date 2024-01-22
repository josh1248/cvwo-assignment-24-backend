Planned backend:
- go-gin for routing
- TODO: write into struct with jmoiron/sqlx?
- ~Postgres as the DB, with 3rd part driver at github.com/lib/pq~
**Edit:** I have decided to change the DB implementation from Postgres to SQLite3. A few reasons:
- A previous consideration in using Postgres before SQLite was because SQLite in Go used to rely on a more complicated SQLite implementation with gcc at https://github.com/mattn/go-sqlite3, requiring complicated initialization. However I discovered a pure Go-based SQLite driver at https://github.com/glebarez/go-sqlite that has embedded SQLite and requires no environment or gcc shenanigans. Hence, setup is much easier. A small downside is the high number of dependencies that the latter driver uses (~20 total dependencies) compared to the former (2 dependencies only).
- SQLite stores its databases within files, making for very easy portability. Researching online, I discovered that Postgres-based Go APIs are not very friendly to use. Some require Postgres to be installed before running, others require a containerized Postgres within Docker to work. SQLite offers a more straightforward path to implementing a portable API since it is serverless.
- Lack of access controls in SQLite does not pose an issue for my small web project, since what I am doing is not sensitive.

Self note:
- PascalCase for public access functions is convention, camelCase for private functionality. https://golang.org/doc/effective_go.html#mixed-caps
- cmd/server/main.go as the main starting point of the backend.
- database folder is in charge of talking with the database server.
- models contain the structs we expect to use, and form the M part of the MVC that handles data and interacts with the DB.
- should probably keep a secret file for duplication in database, so that users can type their own password. TODO: Allow user setup to copy some preset configuration file, then fill in their own postgres details.

# How to initialize (self dev log)
Set up working import statements within my repository with the following:
- Commit this template repo into github. I used a dummy package called `repotest` with some trivial public functions to check for linkage later.
- follow https://go.dev/doc/tutorial/create-module, and run `go mod init <link>` in the root repo, where `<link>` was the github link to my root repo without the `http://` at the front. While you technically do not need to have your module name be an actual URL, it is a standard procedure so that others can get your packet with `go get -u ...` remotely without having to download them.
- get your URL link to your test package, like `repotest`, by checking with `go list ./...`
- In some other go file, use this list as your package statement to test this out.
- Run `go mod tidy`. This command will check all your Go files for import statements and download them as needed. It will then adjust the `go.mod` and `go.sum` files accordingly.
- Run your `main.go` file, in this case within `cmd/server/main.go`, to ensure that your packages have been set up nicely

## API testing
We can test our API's CRUD capabilities using the common HTTP requests `GET`, `POST`, `PATCH`, and `DELETE`.

For this project, the DB lies in `localhost:5432` and the API lies in `localhost:8080`. This is combined with the specified routes in `internal/router` (or `internal/routes`) to form the URL we can use.

Note: `localhost`, for our purposes, refers to the self-referential `127.0.0.1`.

### GET
I can test the routes `/users`, which should return a list of all users, by direct access to `localhost:8080/users`. Alternatively, use the following curl command:
```Bash
curl -X GET http://localhost:8080/users
```
curl defaults to `GET`, so `-X GET` is not necessary.

A much easier method involves using Postman. I am using the Postman extension within VSCode, and I will be using Postman from now on. I will revisit `curl` when I am more comfortable.

### POST
Next, we can deliver a payload and check the output.
```Bash

```

## Other Dev Logs

- Renamed data structure folder and package from `models` to `entities`. Instead, the Go package that interacts with the DB files will be called `models`.

SQLite controls: 
- access a SQLite db with `sqlite3 <filename>`.
- find the database layout with `.schema`.
- exit the sqlite3 shell with `;` to exit multiline query mode, then type `.exit`.

# CVWO Assignment Sample Golang App

This sample Golang app is provided to help you experiment and practice web development fundamentals.
It shows how certain functionality can be implemented.
However, do note that this is **far from a model example**.
After all, we want to see how you maximise your learning in web development
and good software development practices.

## Getting Started

### Installing Go

Download and install Go by following the instructions [here](https://go.dev/doc/install).

### Running the app
1. [Fork](https://docs.github.com/en/get-started/quickstart/fork-a-repo#forking-a-repository) this repo.
2. [Clone](https://docs.github.com/en/get-started/quickstart/fork-a-repo#cloning-your-forked-repository) **your** forked repo.
3. Open your terminal and navigate to the directory containing your cloned project.
4. Run `go run cmd/server/main.go` and head over to http://localhost:8000/users to view the response.


### Navigating the code
This is the main file structure. Note that this is simply *one of* various paradigms to organise your code, and is just a bare starting point.
```
.
├── cmd
│   ├── server
├── internal
│   ├── api         # Encapsulates types and utilities related to the API
│   ├── dataacess   # Data Access layer accesses data from the database
│   ├── database    # Encapsulates the types and utilities related to the database
│   ├── handlers    # Handler functions to respond to requests
│   ├── models      # Definitions of objects used in the application
│   ├── router      # Encapsulates types and utilities related to the router
│   ├── routes      # Defines routes that are used in the application
├── README.md
├── go.mod
└── go.sum
```

Main directories/files to note:
* `cmd` contains the main entry point for the application
* `internal` holds most of the functional code for your project that is specific to the core logic of your application
* `README.md` is a form of documentation about the project. It is what you are reading right now.
* `go.mod` contains important metadata, for example, the dependencies in the project. See [here](https://go.dev/ref/mod) for more information
* `go.sum` See [here](https://go.dev/ref/mod) for more information

Try changing some source code and see how the app changes.

## Next Steps

* This project uses [go-chi](https://github.com/go-chi/chi) as a web framework. Feel free to explore other web frameworks such as [gin-gonic](https://github.com/gin-gonic/gin). Compare their pros and cons and use whatever that best justifies the trade-offs.
* Read up more on the [MVC framework](https://developer.mozilla.org/en-US/docs/Glossary/MVC) which this code is designed upon.
* Sometimes, code formatting can get messy and opiniated. Do see how you can incoporate [linters](https://github.com/golangci/golangci-lint) to format your code.
