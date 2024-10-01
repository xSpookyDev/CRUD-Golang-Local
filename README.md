# CRUD de Películas en Golang

Este es un ejemplo básico de una API CRUD (Create, Read, Update, Delete) en Golang que gestiona una lista de películas. Utiliza el paquete `gorilla/mux` para el enrutamiento y `encoding/json` para manejar datos en formato JSON.

## Requisitos

- **Go 1.16+** instalado en tu sistema.
- Conexión a Internet para descargar las dependencias de `gorilla/mux`.
- Postman para revisar endpoints.

## Instalación

1. Clona el repositorio o copia el código fuente en tu máquina local:

   bash
   git clone https://github.com/xSpookyDev/CRUD-Golang-Local
   cd tu-repositorio

2. Con Postman puedes probar cada uno de los endpoint del crud :
   Endpoints
  Obtener todas las películas
  URL: http://localhost:8000/movies
  Devuelve la lista completa de películas en formato JSON.


  Obtener una película por ID
  URL: http://localhost:8000/movies/{id}

  Crear una nueva película
  URL: http://localhost:8000/movies

 

  Descripción: Crea una nueva película.

  Cuerpo (JSON):

  json
  Copiar código
  {
    "isbn": "123123",
    "title": "Batman",
    "director": {
      "firstname": "Jaime",
      "lastname": "Porro"
    }
  }

  Actualizar una película por ID
  URL: http://localhost:8000/movies/{id}
  Cuerpo (JSON):

  json
 
  {
    "isbn": "987654",
    "title": "Batman Begins",
    "director": {
      "firstname": "Chris",
      "lastname": "Nolan"
    }
  }

  Eliminar una película por ID
  URL: http://localhost:8000/movies/{id}

  Método: DELETE





  Método: DELETE
  URL: http://localhost:8000/movies/19021520

  ### Cambios realizados:
- Agregados separadores para mejorar la organización.
- Añadidos títulos y descripciones más claros.
- Formateado el código para que se vea más limpio y fácil de leer.
