def project = "concrete-plasma-244309"
def appName = "ms-bribrain-brimo-paylater-whitelist"//samakan dengan repository git
def namespace = "digital-bank"
def namespaceuat = "digital-bank-uat"
def proxyType = "http"
def proxyAddress = "172.18.104.20"
def proxyPort = "1707"

pipeline {
  agent {
    kubernetes {
      defaultContainer 'jnlp'
      yaml """
        apiVersion: v1
        kind: Pod
        metadata:
        labels:
          component: ci
        spec:
          tolerations:
          - key: "jenkins"
            operator: "Equal"
            value: "agent"
            effect: "NoSchedule"
          affinity:
            nodeAffinity:
              requiredDuringSchedulingIgnoredDuringExecution:
               nodeSelectorTerms:
                -  matchExpressions:
                   - key: jenkins
                     operator: In
                     values:
                     - agent
          hostAliases:
          - ip: "172.18.46.212"
            hostnames:
            - "sast.bri.co.id"
          - ip: "172.18.224.218"
            hostnames:
            - "devops-k8s-cluster-1-master-1"
          - ip: "172.18.224.56"
            hostnames:
            - "k8s-load-balancer-master-01"
          - ip: "172.18.219.115"
            hostnames:
            - "k8s-loadbalancer-01-sb"
          - ip: "172.18.218.209"
            hostnames:
            - "k8s-loadbalancer-01"
          serviceAccountName: cd-jenkins
          containers:
          - name: gcloud
            image: google/cloud-sdk:263.0.0-alpine
            imagePullPolicy: IfNotPresent
            command:
            - cat
            tty: true
          - name: helm
            image: fathoniadi/helm:2.14.0
            imagePullPolicy: IfNotPresent
            command:
            - cat
            tty: true
          - name: jnlp
            image: mfahry/bri-jnlp-slave:1.7
            imagePullPolicy: IfNotPresent
        """
    }
  }
  
  stages {
    stage('Quality Node') {
      when {
        branch 'development'
      }
      environment {
        scannerHome = tool 'sonarscanner'
      }
      steps {
        withSonarQubeEnv('sonarqube') {
          sh "${scannerHome}/bin/sonar-scanner -X"
        }
      }
    }
    //   stage("Unit Test") {
    //     steps {
    //       container("node") {
    // //        sh "npm install"
    // //        sh "npm test"

    //           sh "echo Skip this processing ..."
    //       }
    //     }
    //   }
    stage("build image") {
      environment {
        IMAGE_REPO = "gcr.io/${project}/${appName}"
        IMAGE_TAG = "${env.GIT_COMMIT.substring(0,7)}"
      }  
      // when {
      //   branch 'sv-branch'
      // }
      steps {
        container("gcloud") {
          // sh "PYTHONUNBUFFERED=1 gcloud builds submit -t ${IMAGE_REPO}:${IMAGE_TAG} ."
        withCredentials([file(credentialsId: "k8s-builder-prod", variable: "JSONKEY")]) {
            sh "gcloud config set proxy/type ${proxyType}"
            sh "gcloud config set proxy/address ${proxyAddress}"
            sh "gcloud config set proxy/port ${proxyPort}"
            sh "gcloud config set builds/timeout 7200"

            sh "cat ${JSONKEY} >> key.json"
            sh "gcloud auth activate-service-account --key-file=key.json"
            sh "gcloud builds submit --project ${project} --tag ${IMAGE_REPO}:${IMAGE_TAG} ."
          }
        }
      }
    }
    stage("Deploy to development") {
      when {
        branch 'development'
      }
      environment {
        IMAGE_REPO = "gcr.io/${project}/${appName}"
        IMAGE_TAG = "${env.GIT_COMMIT.substring(0,7)}"
      }
      steps {
        container("helm") {
          //  sh "helm upgrade --debug --install -f helm/values.dev.yml --set-string image.repository=${IMAGE_REPO},image.tag=${IMAGE_TAG} whitelist-ceria ./helm/whitelist-ceria"
        withCredentials([file(credentialsId: "kubeconfig-gti-dev", variable: "KUBECONFIG")]) {
            // setup kube config
            sh "mkdir -p ~/.kube/"
            sh "cat ${KUBECONFIG} >> ~/.kube/config"
            
            sh """
              helm upgrade ${appName} ./helm/${appName} \
                --set-string image.repository=${IMAGE_REPO},image.tag=${IMAGE_TAG} \
                -f ./helm/values.dev.yaml --debug --install --namespace ${namespace}
            """
          }       
        }
      }
    }
     stage("Deploy to uat") {
      when {
        branch 'uat'
      }
      environment {
        IMAGE_REPO = "gcr.io/${project}/${appName}"
        IMAGE_TAG = "${env.GIT_COMMIT.substring(0,7)}"
      }
      steps {
        container("helm") {
          //  sh "helm upgrade --debug --install -f helm/values.dev.yml --set-string image.repository=${IMAGE_REPO},image.tag=${IMAGE_TAG} whitelist-ceria ./helm/whitelist-ceria"
        withCredentials([file(credentialsId: "kubeconfig-gti-uat", variable: "KUBECONFIG")]) {
            // setup kube config
            sh "mkdir -p ~/.kube/"
            sh "cat ${KUBECONFIG} >> ~/.kube/config"
            
            sh """
              helm upgrade ${appName}-uat ./helm/${appName} \
                --set-string image.repository=${IMAGE_REPO},image.tag=${IMAGE_TAG} \
                -f ./helm/values.uat.yaml --debug --install --namespace ${namespaceuat}
            """
          }       
        }
      }
    } 
    stage("Deploy to production") {
      when {
        branch 'master'
      }
      environment {
        IMAGE_REPO = "gcr.io/${project}/${appName}"
        IMAGE_TAG = "${env.GIT_COMMIT.substring(0,7)}"
      }
      steps {
        container("helm") {
        // sh "helm upgrade --debug --install -f helm/values.prd.yml --set-string image.repository=${IMAGE_REPO},image.tag=${IMAGE_TAG} whitelist-ceria ./helm/whitelist-ceria"
        withCredentials([file(credentialsId: "kubeconfig-gti-prd", variable: "KUBECONFIG")]) {
            // setup kube config
            sh "mkdir -p ~/.kube/"
            sh "cat ${KUBECONFIG} >> ~/.kube/config"
            
            sh """
              helm upgrade ${appName} ./helm/${appName} \
                --set-string image.repository=${IMAGE_REPO},image.tag=${IMAGE_TAG} \
                -f helm/${appName}/values.yaml --debug --install --namespace ${namespace}
            """
            }
          }
        }
      }
     stage('SAST') { 
      when {
         branch 'uat'
       }
    steps {
    script {
    step([$class: 'CxScanBuilder', comment: '', credentialsId: '', excludeFolders: 'helm, vendor', exclusionsSetting: 'job', failBuildOnNewResults: false, failBuildOnNewSeverity: 'HIGH', filterPattern: '!.* ', fullScanCycle: 10, generatePdfReport: true, groupId: '2d95991a-f4d4-43f6-8cb8-4029b9ab4410', incremental: true, password: '{AQAAABAAAAAQ6aiwQviXnYu3oM9JZHpd8ZG+N1gQtlTUD76ogPfFv00=}', preset: '100065', projectName: "$appName", sastEnabled: true, serverUrl: 'https://sast.bri.co.id', sourceEncoding: '1', username: '', vulnerabilityThresholdResult: 'FAILURE', waitForResultsEnabled: false])
        }
      }
    }
    }
  }