

# Contexto:

Gracias a un dataset creado por Barry Becker basado en un censo de 1994 que contiene datos demográficos, tenemos los datos necesarios para comprender qué tipo de personas ganaban más o menos de $50,000 USD.
El problema es que los datos recopilados sólo están disponibles en un archivo de texto plano y actualmente carecemos de una herramienta que nos permita visualizarlos y derivar hipótesis y conclusiones a partir de esa información.

## Antes de comenzar.

Renombra este archivo como instructions.md para que puedas agregar tu propio README.md

La fuente de datos se encuentra en la carpeta data. Es un archivo de texto plano con más de 30,000 registros por lo que como primer recomendación pasarla a una base de datos estructurada sería lo adecuado. Otras recomendaciones:

1. Utiliza una base de datos Postgres.

## Historia de Usuario 1: Como usuario, quiero filtrar y ordenar los datos para obtener información específica.

- Contexto:
  - necesitamos una herramienta que nos permita ver, filtrar y ordenar los datos para obtener información específica sobre ingresos, educación, estado civil, ocupación, edad y nivel de ingresos.
- Descripción:
  - Quisiera poder visualizar los datos en una tabla.
  - Desearía poder filtrar los datos por educación, estado civil, ocupación, rango de edad y nivel de ingresos.
  - Necesito la capacidad de ordenar los datos de manera ascendente o descendente según cualquiera de los atributos de la tabla.
  - Necesito poder ver los elementos con paginación para tener una mejor presentación de los mismos.
- Criterios de Aceptación:
  - Para optimizar la velocidad de los resultados, ejecutar los filtros y la paginación en el backend.
  - Debe haber opciones de filtrado visible en la interfaz.
  - Los datos deben reaccionar dinámicamente a los filtros seleccionados.

## Historia de Usuario 2: Como usuario registrado, quiero poder autenticarme para persistir mis configuraciones de visualización.

- Descripción:
  - Quiero tener la opción de registrarme e iniciar sesión para guardar mis preferencias de visualización -filtros y ordenamiento-.
  - La autenticación debería permitirme acceder a mis configuraciones personalizadas desde cualquier dispositivo.
- Criterios de Aceptación:
  - La interfaz debe tener opciones de registro e inicio de sesión.
  - Después de iniciar sesión, mis configuraciones personalizadas deben ser cargadas automáticamente.

## Historia de Usuario 3: Como usuario, quiero que la aplicación sea accesible desde diferentes dispositivos.

- Contexto
  - necesitamos una aplicación que sea accesible y funcional en diversos dispositivos, incluyendo escritorios, tablets y teléfonos móviles.
- Descripción:
  - La aplicación debe ser responsive y funcionar correctamente en escritorios, tablets y teléfonos móviles.
  - Quiero una experiencia de usuario consistente independientemente del dispositivo que esté utilizando.
- Criterios de Aceptación:
  - La aplicación debe ser accesible y usable en diferentes tamaños de pantalla.
  - Todas las funcionalidades deben ser fácilmente utilizables en dispositivos móviles.
  - 
¡Éxito!



##go run main.go --env-path .env
