## Run gRPC server and Proxy Server

```go
$ go run server/main.go

2018/07/17 16:52:53 gRPC Server is running at port 50051
2018/07/17 16:52:53 Proxy Server is running at port 8088

```

## Hit from browser

```js
http://localhost:8088/apis/demo/intro/json?name=Sajib%20Khan

{
  message: "Hello Sajib Khan"
}
```

## Run go clilent

```go
$ go run client/main.go

2018/07/17 17:06:20 Greeting: Hello Mr. Alice
```

## Deploy server in Kubernetes cluster (e.g. minikube)

1. Build docker image for ubuntu (thoug your local machine is MAC)

```
$ bash
$ GOOS=linux GOARCH=amd64 go build -o myapp server/

# 'myapp' binary will be created. location: server/myapp
```

2. Build a docker image with the `myapp` binary

```
$ docker build -t sajib/grpc . 
```

2. Login docker hub & push the image 'sajib/grpc'

```
$ docker login
# provide email/password

$ docker push <username>/<image>
# e.g. docker push sajib/grpc
```

3. Run `minikube` cluster

```
$ minikube start
```

4.Create a `deployment` with the image

grpc-deployment.yaml sample:

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: grpc-server
  labels:
    app: grpc
spec:
  replicas: 2
  selector:
    matchLabels:
      app: grpc
  template:
    metadata:
      labels:
        app: grpc
    spec:
      containers:
      - name: grpc-server
        image: sajib/grpc2
        ports:
        - containerPort: 8088
```

Create the deployment

```
$ kubectl apply -f grpc-deployment.yaml
```

wait until the pods are ready!

```
$ kubectl get pods -w
```

5. Create a service to access from browser

grpc-service.yaml sample:

```
kind: Service
apiVersion: v1
metadata:
  name: grpc-service
spec:
  type: NodePort
  selector:
    app: grpc
  ports:
  - protocol: TCP
    port: 1234
    targetPort: 8088
```

Create the service:

```
$ kubectl create -f grpc-service.yaml
```

See the service:

```
$ kubectl get services

NAME           TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
grpc-service   NodePort    10.111.102.157   <none>        1234:32249/TCP   26m
```

- Now hit from browser with url: `http://192.168.99.100:32249/apis/demo/intro/json?name=Sajib`

**N.B.** ip is `192.168.99.100` ($ minikube ip) and port is `32249` not `1234` cause port `1234` is binded
 with `32249` port!