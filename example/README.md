```shell
git submodule add https://github.com/getstackhead/pluginlib

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pluginlib/stackhead.proto
```
