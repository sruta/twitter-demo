## About The Project

The goal of this project is to show how I would implement a real world system.

The "SoccerManagerBELite" system is described [here](https://drive.google.com/file/d/1-xZjVGmqS1PlHyB9Z2WBvJxrUFpuii_f/view).

### Built With

* [Go](https://go.dev/) Language
* [Gin-Gonic](https://github.com/gin-gonic/gin) to run the REST API
* [MariaDB](https://mariadb.org/) to persist the data
* [JWT](https://jwt.io/) to authenticate users

### Folder Structure

``` bash
.
├── cmd
│   └── api
│       └── main.go   -> Entrypoint. Setups the router and container
├── internal          -> Private application and library code
│   ├── configs
│   ├── controller    -> Functions to handle all things related to http communication
│   ├── domain        -> Structs that define the system entities
│   ├── helpers       -> Helper functions that need to be initialized by the application
│   ├── middleware    -> Middlewares to be used with gin-gonic framework
│   ├── repository    -> Functions to communicate with the persistance layer
│   ├── service       -> Functions that handle the application logic
│   ├── test          -> Application tests, not done :(
│   └── container.go  -> Container definition used to inject dependencies
├── pkg               -> Library code that can be used by external applications
├── scripts           -> Database scripts
├── go.mod            -> golang dependencies definition
└── README.md
```

The system is coded to use dependency injection based in golang `interfaces`. This allows to unit test controllers,
services and repositories easily because we can mock each one of their dependencies. Also allows to
change, for example, the persistence technology without doing big changes in the codebase. 


In this model the authentication is handled by gin-gonic's `middlewares`. The `controllers` analyze the http requests and validate
their body and parameters, also turn business errors into api errors that had the correct status
code and format. The `services` handle business logic and validations and orchestrate the usage of the dependencies
like external http clients or persistence clients. Finally, the `repositories` are the ones who handle the persistence 
of the different `domain` models.

### Getting Started

1. Clone the repo

   ``` sh
    git clone git@git.toptal.com:screening/Santiago-Ruta.git
   ```
   
2. Set database configurations to connect to a running MariaDB instance. The file is present at `internal/configs/mysql.go`.
   
    Default values are:

      ``` golang
       var MySQLProd = MySQL{
           Host:         "localhost",
           Database:     "soccer_manager",
           User:         "root",
           Pass:         "root",
           Driver:       "mysql",
           MaxOpenConns: 10,
           MaxIdleConns: 10,
       }
    ```
   
3. For the first use, run the scripts present at `scripts/`.
4. For the first use, run the `go mod tidy` command to install the dependencies. `go1.18` must be installed in the 
current machine to be able to run the system.
5. Run `go mod cmd/api/main.go` to get the system up and running at `localhost:8080`.

### Usage

#### Unauthenticated Endpoints

* Create a new user:
    ```
    POST /api/user
    ```
    Body:
    ``` json
    {
        "email": "an_email@a_domain.com",
        "password": "12345"
    }
    ```
    Successful response:
    ``` json
    {
        "id": 1,
        "email": "an_email@a_domain.com"
    }
    ```
* Login: 
    ```
    POST /api/login
    ```
    Body:
    ``` json
    {
        "email": "an_email@a_domain.com",
        "password": "12345"
    }
    ```
    Successful response:
    ``` json
    {
        "token": "a_JWT_token_to_be_sent_in_further_requests"
    }
    ```

#### Authenticated Endpoints

`Authorization: Bearer {token_received_at_login}` header must be sent in order to access the following endpoints.

* Get all users:
    ```
    GET /api/user
    ```
    Successful response:
    ``` json
    [
        {
            "id": 1,
            "email": "an_email@a_domain.com"
        }
    ]
    ```

* Get a specific user:
    ```
    GET /api/user/:id
    ```
    Successful response:
    ``` json
    {
        "id": 1,
        "email": "an_email@a_domain.com"
    }
    ```

* Get all teams:
    ```
    GET /api/team
    ```
    Successful response:
    ``` json
    [
        {
            "id": 1,
            "user_id": 1,
            "country_id": 1,
            "name": "a_team_name",
            "value": 20000000,
            "budget": 5000000
        }
    ]
    ```

* Get logged-in user's team:
    ```
    GET /api/my-team
    ```
  Successful response:
    ``` json
    {
        "id": 1,
        "user_id": 1,
        "country_id": 1,
        "name": "a_team_name",
        "value": 20000000,
        "budget": 5000000
    }
    ```

* Get a specific team:
    ```
    GET /api/team/:id
    ```
  Successful response:
    ``` json
    {
        "id": 1,
        "user_id": 1,
        "country_id": 1,
        "name": "a_team_name",
        "value": 20000000,
        "budget": 5000000
    }
    ```

* Get team's players:
    ```
    GET /api/team/:id/players
    ```
  Successful response:
    ``` json
    [
        {
            "id": 1,
            "team_id": 1,
            "country_id": 1,
            "first_name": "a_player_first_name",
            "last_name": "a_player_last_name",
            "age": 20,
            "position": "DEFENDER",
            "value": 1000000
        }
    ]
    ```

* Modify a team: (only the owner can do this)
    ```
    PUT /api/team/:id
    ```
    Body:
    ``` json
    {
        "country_id": 2,
        "name": "another_team_name"
    }
    ```
  Successful response:
    ``` json
    {
        "id": 1,
        "user_id": 1,
        "country_id": 2,
        "name": "another_team_name",
        "value": 20000000,
        "budget": 5000000
    }
    ```

* Get all players:
    ```
    GET /api/player
    ```
  Successful response:
    ``` json
    [
        {
            "id": 1,
            "team_id": 1,
            "country_id": 1,
            "first_name": "a_player_first_name",
            "last_name": "a_player_last_name",
            "age": 20,
            "position": "DEFENDER",
            "value": 1000000
        }
    ]
    ```

* Get a specific player:
    ```
    GET /api/player/:id
    ```
  Successful response:
    ``` json
    {
        "id": 1,
        "team_id": 1,
        "country_id": 1,
        "first_name": "a_player_first_name",
        "last_name": "a_player_last_name",
        "age": 20,
        "position": "DEFENDER",
        "value": 1000000
    }
    ```

* Modify a player: (only the owner can do this)
    ```
    PUT /api/player/:id
    ```
    Body:
    ``` json
    {
        "country_id": 2,
        "first_name": "another_player_first_name",
        "last_name": "another_player_last_name"
    }
    ```
  Successful response:
    ``` json
    {
        "id": 1,
        "team_id": 1,
        "country_id": 2,
        "first_name": "another_player_first_name",
        "last_name": "another_player_last_name",
        "age": 20,
        "position": "DEFENDER",
        "value": 1000000
    }
    ```

* Get all countries:
    ```
    GET /api/country
    ```
  Successful response:
    ``` json
    [
        {
            "id": 1,
            "name": "ARGENTINA"
        }
    ]
    ```

* Get a specific country:
    ```
    GET /api/country/:id
    ```
  Successful response:
    ``` json
    {
        "id": 1,
        "name": "ARGENTINA"
    }
    ```

* Get all available transfers:
    ```
    GET /api/transfer
    ```
  Successful response:
    ``` json
    [
        {
            "player_id": 1,
            "value": 2500000
        }
    ]
    ```

* Get a specific available transfer:
    ```
    GET /api/transfer/:playerID
    ```
  Successful response:
    ``` json
    {
        "player_id": 1,
        "value": 2500000
    }
    ```

* Create a transfer: (only the player owner can do this)
    ```
    POST /api/transfer
    ```
    Body:
    ``` json
    {
        "player_id": 1,
        "value": 2500000
    }
    ```
  Successful response:
    ``` json
    {
        "player_id": 1,
        "value": 2500000
    }
    ```

* Delete a transfer: (only the player owner can do this)
    ```
    DELETE /api/transfer/:playerID
    ```
  Successful response:
    ``` json
    {}
    ```

* Buy a transferable player:
    ```
    POST /api/transfer/:playerID/buy
    ```
  Successful response:
    ``` json
    {}
    ```

#### Response Codes

* `200` for successful `GET`, `PUT` and `DELETE` requests
* `201` for successful `POST` request
* `400` for unsuccessful requests when the client sends incorrect data
* `401` for unsuccessful requests when the client should be authenticated and is not
* `403` for unsuccessful requests when the client is not authorized to do what is trying to do
* `404` for unsuccessful requests when the requested entity is not found
* `500` for unsuccessful requests when the system fails because of himself 

#### Error Response

When then request is unsuccessful, the server answers with the following format: 

``` json
    {
        "code": 400,
        "message": "a_description_of_what_happened"
    }
```

## Further Improvements

* Add unit tests for the controllers, services and repositories
* Add docker configuration for simpler set up of the develop enviroment
* Get the configs present at `internal/configs` from environment variables
* Add query params to the `GET` endpoints to allow better searches, pagination and related entities serialization options