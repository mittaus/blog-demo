# Construir imagen con nuestra aplicación
docker build -t mittaus/go-app:7.0 .

# Poner en ejecución un contenedor

docker run -it --rm -port 8080:5000 --name my-go-app mittaus/go-app:1.0

docker run -d -p 8000:8000 -e 'server.port=8000' --name myapp3 mittaus/go-app:6.0

docker run -d --network="proyecto_default" -p 8000:8000 -e 'server.port=8000' -e 'database.server=db.dev'  --name web3 proyecto_web:latest