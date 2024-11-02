To run test for User entity you should use:
- go test ./api/models/User -v

To run test for NginxServer entity:
 - go test ./api/models/NginxServer -v

To see the coverage you need use:
For user:
-go test -cover ./api/models/User
For NginxServer:
-go test -cover ./api/models/NginxServer