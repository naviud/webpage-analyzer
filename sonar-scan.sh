#!/bin/bash

/home/udayanga/work/sonar-scanner-cli-4.6.0.2311-linux/sonar-scanner-4.6.0.2311-linux/bin/sonar-scanner \
  -Dsonar.projectKey=webpage-analyzer \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.login=sqp_62c99f77e002cf3b289dd49c354b9111adb7af98