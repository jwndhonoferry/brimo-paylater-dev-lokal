## CODE DEV
1. router.go -> script untuk route ke url api
2. postgres.go -> script untuk connect ke postgres (connection)
3. models.go -> inisialisasi struct 
4. 


# DOCKER
1. docker build -t brimo-paylater-whitelist-dev-lokal .
2. docker run --name brimo-paylater-whitelist-dev-lokal -d -p 9192:9192 -e POSTGRES_DIALECTOR="host=host.docker.internal port=5432 user=postgres password=postgres dbname=brimo_paylater sslmode=disable" brimo-paylater-whitelist-dev-lokal

# KUBERNETES
1. kubectl create token admin-user -n kubernetes-dashboard
2. kubectl create -f brimo-pl.yaml
3. kubectl get all
4. kubectl describe pod/brimo-pl-dev-lokal
5. kubectl expose pod brimo-pl-dev-lokal --type=NodePort --name=brimo-service --port=9192 --target-port=9192
    Seharusnya setelah expose pod brimo-pl-dev-lokal dengan type NodePort bisa diakses melalui port service (contoh PORT(S) 9192:30676) maka akses localhost:30676/api/service/v1 dll ..
6. kubectl get service brimo-service
7. kubectl describe service brimo-service
8. kubectl port-forward pod/brimo-pl-dev-lokal 9192:9192 -n default
Hit menggunakan http://localhost:9192/api/v1.0/customer/inquiry/{acctno}
9. Create ingress
    - ```kubectl run --rm -it --tty pingkungcurl1 --image=curlimages/curl --restart=Never -- curl --resolve "brimopaylater:9192:10.107.81.7" -i http://brimopaylater```

# JENKINS
pass : 6b86a58270e542b0b2cc4174fb95cf0f
docker pull jenkins/jenkins:latest
docker run -p 8080:8080 -p 50000:50000 -v /home/jenkins jenkins/jenkins

docker run -d -p 8080:8080 -p 50000:50000 -v jenkins_home:/var/jenkins_home -v /var/run/docker.sock:/var/run/docker.sock --name jenkins jenkins/jenkins:latest



# ELASTIC & KIBANA
user : elastic
pass : 5*I93LyJFOZchFoJFb8e
Lihat di single-node-cluster
https://www.elastic.co/guide/en/kibana/current/docker.html#run-kibana-on-docker-for-dev