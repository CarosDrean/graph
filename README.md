# Implementación de un Grafo en Go

Este proyecto implementa un grafo en arbol binario en Go, permte realizar operaciones básicas de agregar y eliminar nodos, 
así como una función para balancear el grafo.

## Pruebas

En el archivo `main.go` se encuentran la inicializaión del grafo como las llamadas a los metodos de agregar, eliminar e imprimir el grafo.

### Puede ejecutar el proyecto con el comando
```shell
go run main.go
```

## Estructura

Se crearon 2 paquetes, graph y ordered:

### graph

En este paquete va la implementacion del grafo como arbol binario, exportando solo los metodos necesarios para sus consumidores: Add, Delete y Print*;
la función de balanceo no se exporta, ya que es un proceso interno que se ejecuta al insertar o eliminar nodos.

### ordered

Este paquete nos sirve para poder reusar nuestro grafo con distintos tipos de datos o estructuras, en este caso solo hacemos la implementacion para que el grafo trabaje con enteros.

## Desafío

Considero que la parte del balanceo es la parte mas complicada, por lo abstracta que es.