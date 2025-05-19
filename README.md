## Acerca de este proyecto

El objetivo de este repositorio es mostrar como implementaría, a grandes rasgos, un sistema que tenga funcionalidades
semejantes a [Twitter](https://x.com) (ahora X).

### Construido con

* [Go](https://go.dev/) como lenguaje de programación
* [Gin-Gonic](https://github.com/gin-gonic/gin) para implementar la API REST
* [MariaDB](https://mariadb.org/) para persistir los datos
* [JWT](https://jwt.io/) para autenticar a los usuarios

### Estructura de carpetas

``` bash
.
├── cmd
│   └── api
│       └── main.go     -> Punto de entrada. Configura el router y el contenedor.
├── internal            -> Código privado de la aplicación
│   ├── configs         -> Configuraciones de la aplicación
│   ├── domain          -> Estructuras que definen las entidades del sistema
│   ├── helpers         -> Funciones auxiliares que necesitan ser inicializadas
│   ├── infraestructure -> Código relacionado con la comunicación externa saliente
│   │   └── repository  -> Implementación de los repositorios para persistencia de datos
│   ├── interfaces      -> Código relacionado con la comunicación externa entrante
│   │   ├── controller  -> Funciones para manejar comunicación HTTP
│   │   └── dto         -> Objetos de transferencia de datos
│   ├── middleware      -> Middlewares para el framework gin-gonic
│   ├── usecase         -> Funciones que manejan la lógica de negocio
│   ├── test            -> Pruebas de la aplicación
│   └── container.go    -> Definición del contenedor para inyección de dependencias
├── pkg                 -> Código de librería que puede ser usado por aplicaciones externas
├── scripts             -> Scripts de base de datos
├── go.mod              -> Definición de dependencias de golang
└── README.md           -> Este archivo
```

El sistema está diseñado para usar inyección de dependencias basada en `interfaces` de golang. Esto permite realizar
pruebas unitarias fácilmente en los controladores, servicios y repositorios, ya que podemos simular cada una de sus
dependencias. También permite cambiar, por ejemplo, la tecnología de persistencia sin realizar grandes cambios en el
código.

En este modelo, la autenticación se maneja mediante los `middlewares` de gin-gonic. Los `controllers` analizan las
solicitudes HTTP, validando su cuerpo y parámetros, y también convierten errores de negocio en errores de API con el
código de estado y formato correctos. Los `usecases` manejan la lógica de negocio y las validaciones, y orquestan el uso
de las dependencias como clientes HTTP externos o clientes de persistencia. Finalmente, los `repositories` son los
encargados de manejar la persistencia de los diferentes modelos del `domain`.

### Instalación

1. Clonar el repositorio

   ``` sh
    git clone git@github.com:sruta/twitter-demo.git
   ```

2. Utilizar docker para iniciar la base de datos y el sistema

    ``` sh
    docker compose up -d --build
    ```

3. Para el primer uso conectarse a la base de datos con las credenciales presentes en `./docker-compose.yml` y crear el
   schema `twitter_demo`

4. Ejecutar el script presente en `./scripts/setup_db.sql` para crear las tablas necesarias

5. La aplicación ya se encuentra lista para utilizar en `http://localhost:8080`

### Tests

Se han implementado 3 tipos de tests a modo de ejemplo:

1. Tests unitarios para `usecase/user.go`. Prueban los métodos mockeando la base de datos.

    ``` sh
    go test -v ./internal/usecase
    ```

2. Tests de integración para `controller/user.go`. Prueban los métodos utilizando la implementación real del usecase.

    ``` sh
    go test -v ./internal/interfaces/controller
    ```

3. Tests end-to-end para la API. Para ejecutarlo primero se debe iniciar la base de datos y la aplicación utilizando
   docker compose.

    ``` sh
    go test -v ./test/e2e/
    ```

El test end-to-end realiza las siguientes solicitudes mientras va validando los resultados:

1. Creación de un usuario A
2. Creación de un usuario B
3. Login del usuario A
4. Login del usuario B
5. Obtención del usuario A por ID
6. Obtención del usuario B por ID
7. Creación de un tweet para el usuario B
8. Obtención de un timeline vacío para el usuario A
9. Creación de un follower del usuario A al usuario B
10. Obtención de un timeline con un tweet para el usuario A

### Uso

#### Endpoints sin autenticación

* Crear un nuevo usuario:
  ```
  POST /api/v1/user
  ```
  Cuerpo:
  ``` json
  {
    "email": "un_email@un_dominio.com",
    "password": "12345",
    "username": "un_nombre_de_usuario"
  }
  ```
  Respuesta exitosa:
  ``` json
  {
    "id": 1,
    "email": "un_email@un_dominio.com",
    "username": "un_nombre_de_usuario",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
  ```

* Login:
  ```
  POST /api/v1/login
  ```
  Cuerpo:
  ``` json
  {
    "email": "un_email@un_dominio.com",
    "password": "12345"
  }
  ```
  Respuesta exitosa:
  ``` json
  {
    "token": "un_token_JWT"
  }
  ```

#### Endpoints con autenticación

El header `Authorization: Bearer {token_recibido_en_el_login}` debe ser enviado en las solicitudes para poder acceder
a los siguientes recursos.

* Obtener un usuario por ID:
  ```
  GET /api/v1/user/:id
  ```
  Respuesta exitosa:
  ``` json
  {
    "id": 1,
    "email": "un_email@un_dominio.com",
    "username": "un_nombre_de_usuario",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
  ```

* Modificar un usuario por ID:
  ```
  PUT /api/user/:id
  ```
  Cuerpo:
  ``` json
  {
    "id": 1,
    "username": "un_nombre_de_usuario_editado"
  }
  ```
  Respuesta exitosa:
  ``` json
  {
    "id": 1,
    "email": "un_email@un_dominio.com",
    "username": "un_nombre_de_usuario_modificado",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-02-01T00:00:00Z"
  }
  ```

* Crear un tweet:
  ```
  POST /api/v1/tweet
  ```
  Cuerpo:
  ``` json
  {
    "user_id": 1,
    "text": "el_texto_de_un_tweet"
  }
  ```
  Respuesta exitosa:
  ``` json
  {
    "id": 1,
    "user_id": 1,
    "text": "el_texto_de_un_tweet",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
  ```

* Obtener un tweet por ID:
  ```
  GET /api/v1/tweet/:id
  ```
  Respuesta exitosa:
  ``` json
  {
    "id": 1,
    "user_id": 1,
    "text": "el_texto_de_un_tweet",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
  ```

* Modificar un tweet por ID:
  ```
  PUT /api/tweet/:id
  ```
  Cuerpo:
  ``` json
  {
    "id": 1,
    "user_id": 1,
    "text": "el_texto_de_un_tweet_modificado"
  }
  ```
  Respuesta exitosa:
  ``` json
  {
    "id": 1,
    "user_id": 1,
    "text": "el_texto_de_un_tweet_modificado",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-02-01T00:00:00Z"
  }
  ```

* Crear un follower:
  ```
  POST /api/v1/follower
  ```
  Cuerpo:
  ``` json
  {
    "follower_id": 1,
    "followed_id": 2
  }
  ```
  Respuesta exitosa:
  ``` json
  {
    "follower_id": 1,
    "followed_id": 2,
    "created_at": "2025-01-01T00:00:00Z"
  }
  ```

* Obtener el timeline del usuario logueado:
  ```
  GET /api/v1/timeline
  ```
  Respuesta exitosa:
  ``` json
  [
    {
      "id": 1,
      "user_id": 1,
      "text": "el_texto_de_un_tweet",
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": "2025-01-01T00:00:00Z",
      "user": {
        "id": 1,
        "email": "un_email@un_dominio.com",
        "username": "un_nombre_de_usuario",
        "created_at": "2025-01-01T00:00:00Z",
        "updated_at": "2025-01-01T00:00:00Z"  
      } 
    }
  ]
  ```

#### Respuesta errónea

Cuando la solicitud no es exitosa, el servidor responde con el siguiente formato:

``` json
{
  "code": 400,
  "message": "una_descripcion_del_error"
}
```

#### Códigos de respuesta

* `200` para solicitudes `GET`, `PUT` y `DELETE` exitosas
* `201` para solicitudes `POST` exitosas
* `400` para solicitudes fallidas cuando el cliente envía datos incorrectos
* `401` para solicitudes fallidas cuando el cliente debería estar autenticado y no lo está
* `403` para solicitudes fallidas cuando el cliente no está autorizado para realizar la acción
* `404` para solicitudes fallidas cuando la entidad solicitada no se encuentra
* `500` para solicitudes fallidas cuando el sistema falla por sí mismo
