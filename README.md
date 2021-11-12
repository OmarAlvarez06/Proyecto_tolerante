# Tolerante_A_Fallas
El proyecto consiste en crear una aplicación donde se puedan registrar a los alumnos con sus respectivas materias y calificaciones. El usuario podrá consultar el promedio de un alumno en especifico y el promedio de una materia en específico. De esta manera el usuario podrá ver, modificar, eliminar y agregar nuevos alumnos con sus respectivas materias y calificaciones. Además, consultando en cualquier momento el promedio de un alumno o de alguna materia en específico.
Pueden observar el funcionamiento en el siguiente video en YouTube: https://www.youtube.com/watch?v=E9K_MWh4_mI
En esta aplicación serán utilizadas las siguientes herramientas:
•	Docker Hub
•	Kubernetes
•	Istio
Para la instalación de cada una de las herramientas anteriores se tienen que realizar lo siguiente: 
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

Para la instalación de “Kubernetes” se tienen que realizar los siguientes pasos:
1.	Para instalar Kubernetes es necesario tener instalado y configurado correctamente Docker. Una vez asegurado de que esto se ha realizado correctamente, procederemos a abrir Docker Desktop.
2.	Una vez abierto e inicializado correctamente, se dirige a la parte de configuración que se encuentra en la parte superior de la aplicación. 
3.	Una vez abierta la configuración entraremos al apartado de Kubernetes que se encuentra de lado izquierdo. 
4.	En esta parte de la configuración estará una opción de habilitar Kubernetes y mostrar los contenedores del sistema. Es necesario aplicar y seleccionar la primera opción de habilitar Kubernetes.
5.	Enseguida una vez seleccionada la opción presionaremos el botón de Aplicar y Restaurar. Docker aplicara los cambios e iniciara nuevamente, y ahora ya podremos visualizar el icono de Kubernetes y si todo funciona correctamente, podremos visualizar los iconos de Docker y Kubernetes en color verde, esto significa que ya funcionan y están corriendo estos dos servicios.

