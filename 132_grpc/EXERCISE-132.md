# Exercise for gRPC

## Setup

Make sure you have the `protoc` compiler installed. 

Also, make sure you have the two plugins installed, as described here:
https://grpc.io/docs/languages/go/

## Start

This directory contains a schema definition file `pizza.proto`.

Examine the schema, then generate the required files.

You can use the following command to generate the files:

```shell
protoc --go_out=. --go-grpc_out=. pizza.proto
```

Then, run the server and client code:

```shell
go run cmd/server/main.go
```

```shell
go run cmd/client/main.go
```

## Calling from third-party clients

Use a third-party client like `grpcurl`, Insomnia or Yaak to call the server.

## Implement streaming of products

Read the documentation on how to implement streaming in gRPC.

https://grpc.io/docs/languages/go/basics/#server-side-streaming-rpc

Use that knowledge to provide a method body for `func (s *server) ListProducts()`.

## Creating a product

Now extend the `.proto` file as well as client and server to create a new RPC
for creating a new product.

You can actually store the new product in some server-side slice or such if you like, but it is not
mandatory.

Extend the client to call that new RPC.
