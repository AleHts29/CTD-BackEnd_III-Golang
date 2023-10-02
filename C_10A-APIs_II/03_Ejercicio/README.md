### Especialización en Back End III - Ejercitación

## Ejercitación Individual

**Nivel de complejidad:** Medio 🔥🔥

### Objetivo

El objetivo de esta guía práctica es afianzar los conceptos sobre arquitectura web, JSON y el framework Gin, que hemos visto en este módulo.

### Problema

Una empresa necesita un sistema para gestionar sus empleados. Para lograrlo, requiere un servidor que ejecute una API que les permita manipular a los empleados.

#### Ejercicio 1: Iniciando el proyecto

1. Crea un repositorio en GitHub para almacenar tus avances. Este repositorio será utilizado para mantener todo lo que realices durante las distintas prácticas de Go Web.

2. Clona el repositorio creado.

3. Inicia tu proyecto Go con el comando `go mod init`.

4. A continuación, crea un archivo `main.go` donde cargarás un slice de empleados desde una función que devuelva una lista de empleados. Este slice se debe cargar cada vez que se inicie la API para realizar las distintas consultas.

   La estructura de los empleados es la siguiente:

   ```go
   type Empleado struct {
       Id     int
       Nombre string
       Activo bool
   }
   ```

#### Ejercicio 2: Empleados

Vamos a levantar un servidor utilizando el paquete Gin en el puerto 8080. Para probar nuestros endpoints, utilizaremos Postman.

1. Crea una ruta `/` que devuelva un mensaje de bienvenida al sistema. Por ejemplo: "¡Bienvenido a la empresa Gophers!".

2. Crea una ruta `/employees` que devuelva la lista de todos los empleados en formato JSON.

3. Crea una ruta `/employees/:id` que devuelva un empleado por su ID. Asegúrate de manejar el caso en el que no se encuentre un empleado con ese ID.

4. Crea una ruta `/employeesparams` que permita crear un empleado mediante parámetros y lo devuelva en formato JSON.

5. Crea una ruta `/employeesactive` que devuelva la lista de empleados activos. También debería ser capaz de devolver la lista de los empleados no activos.

¡Espero que esta guía práctica te ayude a afianzar tus conocimientos sobre desarrollo web con Go y Gin!