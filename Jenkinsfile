pipeline {

  environment {
    dockerimagename = "ferryjwndhono/brimo-paylater-dev-lokal:2.0"
  }

  agent any

  stages {
    stage('Checkout Source') {
      steps {
        script {
          // If repo is private, provide credentialsId
          git credentialsId: 'github-ferry-user-pass', url: 'https://github.com/jwndhonoferry/brimo-paylater-dev-lokal.git'
        }
      }
    }

    stage('Checkout Docker') {
      steps {
        script {
          sh 'docker --version'
          }
      }
    }
    
    stage('Build Image') {
      steps {
        script {
          // dockerImage = docker.build("${dockerimagename}")
          sh "docker build -t ferryjwndhono/brimo-paylater-dev-lokal ."
        }
      }
    }

    stage('Pushing Image') {
      environment {
        registryCredential = 'dockerhub-ferry'
      }
      steps {
        script {
          docker.withRegistry('https://registry.hub.docker.com', registryCredential) {
            dockerImage.push("latest")
          }
        }
      }
    }

    stage('Deploying Brimo Paylater container to Kubernetes') {
      steps {
        script {
          kubernetesDeploy(
            configs: ["deployment-brimo-pl-local.yaml", "service-brimo-pl-local.yaml"]
          )
        }
      }
    }

  }

}
