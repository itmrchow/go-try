# gRPC (Remote Procedure Calls)
- base on http/2
- Protocol Buffers作為介面描述語言
- 有
  - authentication
  - 雙向串流
  - 流量控制
  - 阻塞
  - 非阻塞繫結
  - 取消和逾時

## 常見場景
- mico service 之間互動
- 將手機服務、瀏覽器連接至後台
- 產生高效的客戶端庫

## Protocol Buffers
- 是一種開源跨平台的序列化資料結構的協定

```
syntax = "proto3";

package poker;
option go_package = "itmrchow/go-project/try";

message GetNutsRequest{
    repeated string hand = 1;
    repeated string river = 2;
}

message GetNutsResponse{
    repeated string card = 1;
}

service Poker {
    rpc GetNuts(GetNutsRequest) returns (GetNutsResponse);
}
```

產生代碼指令
```
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
poker.proto
```

## init
```
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go get -u google.golang.org/grpc
```

## gRPC 四種生命週期
1. 一對一 (Unary)
2. 多對一 (Client-side streaming): Client 送很多請求(streaming) , Server只回一個 
3. 一對多 (Server-side streaming): Client 送一個請求 , Server回覆多個(streaming)
4. 多對多 (Bidirectional streaming): Client 和 Server 都以 Streaming 的形式交互

## 開發順序
1. 寫.proto
2. create pb file
  - .pb.go : 序列反序列
  - gRPC 服務器和客戶端的程式 
3. 

# 參考
https://yulinchou.medium.com/2022-%E7%94%A8-golang-%E5%AF%A6%E4%BD%9C%E7%B0%A1%E6%98%93-grpc-server-88dadeac8a3d