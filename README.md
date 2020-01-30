# status_service_bot

бот для алерта по недоступности веб сервисов


docker build -t status_service_bot .
docker run --name status_service_bot --env-file services.env status_service_bot  