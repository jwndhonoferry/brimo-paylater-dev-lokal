pipeline {

  environment {
    dockerimagename = "ferryjwndhono/brimo-paylater-dev-lokal:1.0"
    dockerImage = ""
  }

  agent any

  stages {

    stage('Checkout Source') {
      steps {
        git 'https://github.com/jwndhonoferry/brimo-paylater-dev-lokal.git'
      }
    }

    stage('Build image') {
      steps{
        script {
          dockerImage = docker.build dockerimagename
        }
      }
    }

    stage('Pushing Image') {
      environment {
               registryCredential = 'docker-ferry'
           }
      steps{
        script {
          docker.withRegistry( 'https://registry.hub.docker.com', registryCredential ) {
            dockerImage.push("latest")
          }
        }
      }
    }

    stage('Deploying Brimo Paylater container to Kubernetes') {
      steps {
        script {
          kubernetesDeploy(configs: "deployment-brimo-pl-local.yaml", "service-brimo-pl-local.yaml")
        }
      }
    }

  }

}