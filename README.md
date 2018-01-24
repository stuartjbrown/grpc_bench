# grpc_bench
Example code demonstrating a simple benchmark comparison of gRPC vs REST.  Includes test of simple gRPC (Lookup + Set) and REST (GET + POST) endpoints

Quickstart:
* `make get`
* `make build`
* Run `bin/grpc_bench`
* `make bench`

Results:

```
BenchmarkGRPCLookup-4      	   10000	    202760 ns/op
BenchmarkGRPCSetEntity-4   	   10000	    207029 ns/op
BenchmarkRESTLookup-4      	    2000	   1090675 ns/op
BenchmarkRESTSet-4         	    1000	   1709004 ns/op
```
