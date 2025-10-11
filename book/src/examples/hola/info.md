Este folder informa de la creacion de un software hecho con go.

Lo primero es crear un folder como encontraras aqui para este ejemplo.

```bash
mkdir hola
cd hola
```

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
