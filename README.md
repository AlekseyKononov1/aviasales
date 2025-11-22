# aviasales
just pet-project, entry-app in stack {"go","rest", "docker", "postgresql"}

before run you need to ask me to get sql file of database, 500 MB.

run command:
> download docker on desktop
> git clone [https](https://github.com/AlekseyKononov1/aviasales/)
> cd to aviasales project where dockers files
> **docker-compose up --build**

example of client requests:
- http://localhost:8080/flights?start=2026-01-01T00:00:00Z&end=2026-02-01T00:00:00Z&dep_city=Moscow&dep_country=Russia&arr_country=China&arr_country=Japan
- http://localhost:8080/segments/free?flight_id=17109&fare_conditions=Business
