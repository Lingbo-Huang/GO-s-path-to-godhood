syntax = "proto3";
package service;

// 用于指定golang package的名字，新版本要求必须包含 /
option go_package = "service/";

message HelloRequest {
int64 id = 1;
}

message HelloResponse {
string name = 1;
}

service HelloService {
rpc SayHello(HelloRequest) returns (HelloResponse);
}