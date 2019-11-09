*FLEXKART CATALOG*

* Pre requisites
    * dep
    * golang
    * mysql
* Run steps
    * `dep ensure`
    * `make vendor`
    * to change protobuf definitions `make protobuf`
    * `go run cmd/server/*.go`