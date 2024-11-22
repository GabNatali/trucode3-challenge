FROM golang:1.23-alpine

# Instala Air para hot reload
RUN go install github.com/air-verse/air@latest


RUN mkdir app

# Crea el directorio de trabajo
WORKDIR /app

# Copia el archivo go.mod y go.sum
COPY go.mod go.sum ./

# Descarga las dependencias
RUN go mod download

# Copia el código fuente al contenedor
COPY . .

ARG PORT=8080

# Copia los archivos del proyecto
COPY . .

# Expone el puerto
EXPOSE 8080

# Comando para iniciar Air
CMD ["air"]
