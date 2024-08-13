# Guía de Instalación de PostgreSQL usando Docker

Este documento proporciona instrucciones para ejecutar una instancia de PostgreSQL en un contenedor Docker.

## Prerrequisitos

Antes de comenzar, asegúrate de tener instalado Docker en tu sistema. Si no lo tienes, puedes seguir las instrucciones de instalación en la [documentación oficial de Docker](https://docs.docker.com/get-docker/).

## Instalación

Para desplegar PostgreSQL en un contenedor Docker, sigue los siguientes pasos:

### 1. Ejecutar el contenedor de PostgreSQL

Para iniciar un contenedor de PostgreSQL, utiliza el siguiente comando:

```bash
docker run --name some-postgres -e POSTGRES_PASSWORD=1124 -p 5432:5432 -d postgres
```

 ### 2. Testear contenedor
```bash
docker ps
```

 