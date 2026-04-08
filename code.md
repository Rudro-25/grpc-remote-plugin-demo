# code

db.proto


proto/db.pb.go
Contains: DBRequest, DBResponse

proto/db_grpc.pb.go
connector

```
client.go
   ↓
calls GetCredentials()
   ↓
gRPC sends request
   ↓
server.go receives it
   ↓
server logic runs
   ↓
response returned
   ↓
client prints result
```