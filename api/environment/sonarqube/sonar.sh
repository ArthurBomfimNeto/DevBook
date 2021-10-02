#!/bin/bash
sudo docker stop sonar
sudo docker rm sonar
workspace="${PWD/"/environment/sonarqube"/}"
cd $workspace/environment/sonarqube
echo "Iniciando Sonar"
sudo docker-compose up -d
sleep 60
cd $workspace
echo "Configurando Sonar Scanner"
wget https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.5.0.2216-linux.zip
unzip sonar-scanner-cli-4.5.0.2216-linux.zip
binary="$workspace/sonar-scanner-4.5.0.2216-linux/bin/sonar-scanner"
go test -coverprofile=cover.out -covermode=count ./...
$binary -Dproject.settings=sonar-project.properties
rm $workspace/sonar-scanner-cli-4.5.0.2216-linux.zip
rm -rf $workspace/sonar-scanner-4.5.0.2216-linux
echo "Métricas disponíveis pela url: http://localhost:9000/"
