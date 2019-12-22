# jira-537 (DEVOPS)

## How to 

1. Run the server in K8S

```
kubectl apply -f deployment/server.yaml
```

2. Install WireShark to capture packets,see [how to capture http2](https://github.com/elixir-grpc/grpc/wiki/How-to-capture-HTTP2-packages-using-Wireshark).

3. Turn on OpenVPN and run the client

```
go run helloworld/client/main.go -addr HOSTNAME -wait 351
```
