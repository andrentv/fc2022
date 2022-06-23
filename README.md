kafka git clone https://github.com/andrentv/fc2022_kafka.git
sudo apt-get remove docker docker-engine docker.io containerd runc
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
systemctl start docker
systemctl enable docker
sudo systemctl enable docker.service
sudo systemctl enable containerd.service
docker-compose up -d
docker-compose ps
docker exec -it simulator bash
docker exec -it kafka_kafka_1 bash
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal
kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
// {"clientId":"1","routeId":"1"} 
// {"clientId":"2","routeId":"2"}
// {"clientId":"3","routeId":"3"}
