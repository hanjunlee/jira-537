# jira-537 (DEVOPS)

## How to 

1.Run the server in K8S

```
kubectl apply -f deployment/server.yaml
```

2.Install WireShark to capture packets

See [how to capture http2](https://github.com/elixir-grpc/grpc/wiki/How-to-capture-HTTP2-packages-using-Wireshark).

3.Turn on OpenVPN and run the client

```
go run helloworld/client/main.go -addr helloworld.buzzvil-internal.com:50051 -wait 351
```
