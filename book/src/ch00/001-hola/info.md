Como en todo lenguaje existe la necesidad de empezar con un hola, pues aqui no va a
ser la excepcion, pues espero indicarte de la mejor manera.

# Lo primero decidir

Pues claro toca decidir si queremos solo mostrar un hola mundo en tu compu o crear un modulo dentro de tu repositorio que lo quieres publicar para el publico. Pero para este caso eligo por ti, es mejor iniciar desde local ya luego abordaremos el caso de publicar algo propio tuyo entonces manos a la obra.

## Crear una carpeta para el proyecto, 
Estes en linux o windowso o una mac pues lo primero es crear una carpeta con el nombre de tu proyecto en nuestro caso **hola**


```bash
mkdir hola
cd hola
```
Luego ya dentro de la carpeta **hola** hay que inicializar el modulo hola

```bash
go mod init mycompany.com/hola
```

como puedes ver en el codigo anterior se menciona a hola  como el nombre del paquete de este ejemplo y es hola el mismo nombre con el que va la carpeta, tomemos en cuenta que utilizamos mycompany.com, pero si vamos a crear un modulo que quisieramos publicar, previamente deberemos crear el repositorio y si asi lo hiciesemos tendriamos que cambiar el paso anterior a esto:


```bash
go mod init github.com/user/repositorio/package
```

Entonces digamos que queremos publicar una herramienta llamada apptonta

entonces la publicacion de este modulo seria 

go mod init github.com/nicolasmaduro/basurero/apptonta/hello
