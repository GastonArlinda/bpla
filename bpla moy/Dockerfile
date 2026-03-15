FROM ubuntu:24.04

RUN apt-get update && apt-get install -y g++ cmake libssl-dev librdkafka-dev

WORKDIR /app

COPY . /app

RUN g++ -std=c++17 *.cpp -o producer -lrdkafka++ -lrdkafka -pthread

CMD ["./producer"]