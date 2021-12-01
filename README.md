# Tolerante_A_Fallas
El proyecto consiste en crear una aplicación donde se puedan registrar a los alumnos con sus respectivas materias y calificaciones. El usuario podrá consultar el promedio de un alumno en especifico y el promedio de una materia en específico. De esta manera el usuario podrá ver, modificar, eliminar y agregar nuevos alumnos con sus respectivas materias y calificaciones. Además, consultando en cualquier momento el promedio de un alumno o de alguna materia en específico.
Pueden observar el funcionamiento en el siguiente video en YouTube: https://www.youtube.com/watch?v=E9K_MWh4_mI
En esta aplicación serán utilizadas las siguientes herramientas:
•	Docker Hub
•	Kubernetes
•	Istio
Para la instalación de cada una de las herramientas anteriores se tienen que realizar lo siguiente:

| INSTALACIÓN DOCKER |.
Para la instalación de “Docker” se realizan los siguientes pasos:
1.	Dirigirte a la página oficial de Docker el cual se encuentra en el siguiente enlace: https://www.docker.com/
2.	En la esquina superior derecha se encuentra el botón “Empezar” el cual te llevara a la siguiente página donde podrás seleccionar la aplicación de Docker que deseas instalar o utilizar y para el sistema operativo que desees. Docker es una aplicación disponible para Windows, Linux y Mac.
3.	Seleccionar el sistema operativo que deseas utilizar con Docker y automáticamente la página descargará un archivo ejecutable el cual será el ejecutable de Docker para su instalación en el sistema operativo que seleccionaste. 
4.	Una vez finalice la descarga o inclusive antes, si has seleccionado Windows como sistema operativo asegurarse de tener activada la opción de Hyper-V en Windows y la opción Windows Subsystem for Linux. 
Una vez activadas, el sistema pedirá reiniciar el ordenador, es por eso por lo que es recomendable hacerlo cuando la descarga ya haya finalizado. Algo importante es también conocer y saber si el ordenador cuenta con las especificaciones mínimas que requiere Docker para funcionar correctamente. De ser así, continuaremos.
5.	Enseguida necesitaremos descargar WSL 2 para Windows. Para realizar la descargar y la correcta configuración pueden visitar alguna página web o algún tutorial, descargar el instalador/actualizador y configurar como la versión por defecto de la virtualización de Windows. Enlace de ayuda: https://www.windowscentral.com/how-install-wsl2-windows-10
6.	Una vez instalado y configurado el sistema con lo anterior, es necesario de reiniciar el ordenador.
7.	Una vez reiniciado el ordenador ejecutaremos el archivo que descargamos de Docker. Este archivo nos ejecutara el instalador el cual nos guiara para la correcta instalación de Docker. Es muy sencillo y su instalación puede tardar un poco o no por distintos factores. Esperar a que se instale correctamente.
8.	Una vez instalado correctamente, si la aplicación no se ha ejecutado automáticamente tendrás que buscarla en tu menú de aplicaciones y ejecutarla. Su ejecución puede tardar un poco, pero si todo salió bien, no debería mostrar algún error y la aplicación se abrirá automáticamente mostrando un mensaje de que no se han encontrado contenedores ejecutándose y les proporcionara una opción de comando para comenzar a ejecutar un contenedor en Docker. Si todo salió correctamente, Docker ya podría utilizarse sin ningún inconveniente. Si proporciona algún error, la aplicación le proporcionara el nombre del error y algunas alternativas para su solución.

| INSTALACIÓN KUBERNETES |.
Para la instalación de “Kubernetes” se tienen que realizar los siguientes pasos:
1.	Para instalar Kubernetes es necesario tener instalado y configurado correctamente Docker. Una vez asegurado de que esto se ha realizado correctamente, procederemos a abrir Docker Desktop.
2.	Una vez abierto e inicializado correctamente, se dirige a la parte de configuración que se encuentra en la parte superior de la aplicación. 
3.	Una vez abierta la configuración entraremos al apartado de Kubernetes que se encuentra de lado izquierdo. 
4.	En esta parte de la configuración estará una opción de habilitar Kubernetes y mostrar los contenedores del sistema. Es necesario aplicar y seleccionar la primera opción de habilitar Kubernetes.
5.	Enseguida una vez seleccionada la opción presionaremos el botón de Aplicar y Restaurar. Docker aplicara los cambios e iniciara nuevamente, y ahora ya podremos visualizar el icono de Kubernetes y si todo funciona correctamente, podremos visualizar los iconos de Docker y Kubernetes en color verde, esto significa que ya funcionan y están corriendo estos dos servicios.

| DESPLIEGUE EN KUBERNETES |.
Para montar la aplicación en un contenedor de kubernetes:
1. Después de la instalación de kubernetes debemmos verificar que la variable de entorno kubectl se haya establecido correctamente, en caso de que no, podremos hacerlo manualmente desde el editor de variables de entorno, en nuestro caso, de Windows.
2. El siguiente paso es iniciar el servicio de minikube utilizando el comando "minikube start".
3. Una vez iniciado el servicio nos aparecerá que kubectl ahora esta configurado para usar minikube y aparecera el contenedor minikube dentro de los contenedores activos de Docker.
4. Ahora utilizaremos un comando para verificar que el cluster de kubernetes se inició correctamente, el comando es "kubectl ger svc", en el listado deberá aparecer el cluster de kubernetes con información adicional.
5. El suihuiente paso es utilizar dos comandos para verificar el estado de nuestros pods y deployments dentro de la alicación. para esto escribimos "kubectl get pods" y "kubectl get deployments", con esto nos apareceran los listados de los pods y deployments activos.
6. Para poder montar el rpoyecto dentro del cluster es necesario tener la imagen del proyecto en u repositorio de Docker Hub, para ello utilizaremos el comando "docker build -t username/proyecto-tolerante .", en la parte que dice userneame debemos poner el nombre de usuario de nuestra cuenta de Docker Hub.
7. Con el paso anterior se crear la imagen que se puede subir al repositorio. Para esto se utiliza el comando "docker push username/proyecto-tolerante:latest". Con este comando comenzará a ahcerse el push dentro del repositorio, una vez termine nos avisará que el push se hizo correctamente.
8. Ahora lo que queda es hacer el deploy en kubernetes, para esto nececitamos un archivo de configuración .yaml (llamado deployment.yaml) donde especificaremos la app que queremos desplegar, la imagen desde el repositorio, los limites de memoria y cpu y el puerto donde será montada. El archivo de configuración es el siguiente:
9.  apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: my-proyecto-tolerante
    spec:
      replicas: 1
      selector:
        matchLabels:
          app: proyecto-tolerante
      template:
        metadata:
          labels:
            app: proyecto-tolerante
        spec:
          containers:
          - name: proyecto-tolerante-container
            image: ccisaias/proyecto-tolerante
            resources:
              limits:
                memory: "128Mi"
                cpu: "500m"
            ports:
            - containerPort: 3000

10. Previamente debemos tener además, el contenedor de esta aplicación activo dentro de docker, para ello construimos la imagen y montamos el contenedor con los siguientes comandos respectivamente: "docker build -t proyecto-tolerante-img ." y "docker run -d -p 3333:3000 --name proyecto-tolerante-container proyecto-tolerante-img".
11. Ahora haremos el deploy utilizando "kubectl create -f deployment.yaml".
12. Después revisamos que el pod y el deployment se hayan configurado correctamente y esten funcionando, esto con los comandos: "kubectl create -f deployment.yaml" y "kubectl get pods".
13. El ultimo paso será hacer un expose a la aplicación con el comando "kubectl expose deployment my-proyecto-tolerante --type=NodePort --name=proyecto-tolerante-svc --target-port=3000".
14. verificamos que se haya creado utilizando "kubectl get svc" y deberá aparecer en el listado al parecido a: 
            NAME                     TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
            proyecto-tolerante-svc   NodePort    10.111.127.140   <none>        3000:30652/TCP   81s
15. Para finalizar se utiliza el comando "minikube ip" para obtener la ip a la cual vamos a acceder deade el navegador.
16. Ahora solo queda probar el cluster accediendo a minikube-ip:nodePort, que en este caso es "192.168.99.100:30652", si nos muestra la aplicación, se ah configurado correctamente el cluster de kubernetes.
  

