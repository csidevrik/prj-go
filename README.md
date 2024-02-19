go projects 
------------------------------------------------

This repo shares examples of projects created using go language programming.

## Install Cobra framework
Para hacer una cli ya existe un framework llamado cobra, este framework permite obtener una aplicacion cli organizada y elegante, usando go como lenguaje de programacion.

- [Cobra framework](https://cobra.dev//)

Ademas tenemos una herramienta cli de cobra que puedes utilizar para generar mas aplicaciones.
- [Cobra cli ](https://github.com/spf13/cobra-cli/blob/main/README.MD)

## Pasos para la creacion del primer proyecto
En nuestro caso tenemos un proyecto al que llamamos netmacfilter, esta herramienta deberia tener algunas funciones:

Commands:   
    - recognizeFormat tomar una mac y presentar el formato en el cual esta presentado.
    - 


```shell
    mkdir netmacfilter
```
```shell
    cd netmacfilter
```
```shell
    go mod init github.com/carlossiguam/prj-go/netmacfilter/getformatMAC
```

### Nombrando al author y estableciendo la licencia
A continuacion ponemos el autor  del proyecto de la cli y ademas la licencia de distribucion de la cli,
para esto tenemos el siguiente comando

Ok not yet get the information  about to write on this part -l license,.

```shell
cobra-cli init -a "carlos sigua" -l "MIT" 
```

Now the next secuencia, first you neeed to create a function to do an action specific for then call it into tha cli-app, ok let's get to work.

Now we open the folder called cmd and into this we create a function, that by other hands for this cli app this function print if a data adding as argument is or not is format MAC ADDRESS aruba, cisco or huawei or linux or windows.

Use this tutorial as example 
- [Cobra cli example ](https://www.digitalocean.com/community/tutorials/how-to-use-the-cobra-package-in-go)


Aprendiendo a crear una aplicacion de saludo en red.

## Step 1 Install cobra-cli
Instalar cobra-cli
Habiendo primero declarado la variable de entorno GOBIN, la cual por defecto suele instalarse en ~/go/bin, ahora el siguiente paso es utilizar el framework cobra-cli

## Step 2 Create a project o aplicacion
Para esto tenemos que crear una carpeta con el nombre de la aplicacion o del proyecto en este caso netgreeting
```shell
    mkdir netgreeting
    cd netgreeting
```
Donde netgreeting es el nombre de la aplicacion, mas no el nombre de un modulo, ahora vamos a crear un modulo

## Step 3 Inicializar el modulo uno
Esto es decirle 

solo necesitas go mod init

solo necesitas cobra-cli init


