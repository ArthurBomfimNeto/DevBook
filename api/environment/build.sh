#!/bin/bash
echo "Configurando ambiente de desenvolvimento"
path=$PWD
chmod +x $path/sonarqube/sonar.sh
cd $path/sonarqube
./sonar.sh
