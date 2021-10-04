Sau khi có [protos](https://github.com/ducnguyen96/ducnguyen96.xyz-protos) ta sẽ thiết lập gRPC server và client. 

Tham khảo thêm [ở đây](https://github.com/gin-gonic/examples/tree/master/grpc)
## 1. Get proto
```shell
go get github.com/ducnguyen96/ducnguyen96.xyz-protos
```

## 2. Setup HTTP Server và gRPC client cho gateway
```go
package main

import (
	"fmt"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewGreeterClient(conn)

	// Set up a http server.
	r := gin.Default()
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		// Contact the server and print out its response.
		req := &pb.HelloRequest{Name: name}

		res, err := client.SayHello(c, req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.Message),
		})
	})

	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		panic("Error")
	}
}
```
## 3. Setup gRPC server cho authservice 
```go
package main

import (
	"context"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// server is used to implement utils.GreeterServer
type server struct {
	// Embed the unimplemented server
	pb.UnimplementedGreeterServer
}

// SayHello implements utils.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// listening to tcp
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```
## 4. Run server trên cả gateway và authservice
```shell
go run gateway/server.go
go run authservice/server.go
```
## 5. Test
```shell
curl -v 'http://localhost:8080/hello/ducnguyen96'
```

## 6. Thêm graphql cho gateway
### 6.1. Write Schema
### 6.2. Config gqlgen.yaml
Tham khảo thêm ở [đây](https://gqlgen.com/config/)
```yaml
# Where are all the schema files located? globs are supported eg  src/**/*.graphqls
schema:
  - schema/**/*.graphqls
  - schema/**/*.graphql
# Where should the generated server code go?
exec:
  filename: graph/generated/generated.go
  package: generated

#exec:
#  package: resolver
#  filename: resolver/
# Uncomment to enable federation
# federation:
#   filename: graph/generated/federation.go
#   package: generated

# Where should any generated models go?
model:
  filename: graph/model/models_gen.go
  package: model

# Where should the resolver implementations go?
resolver:
  layout: follow-schema
  dir: graph/resolver
  package: graph
  filename_template: '{name}.resolvers.go'

# Optional: turn on use `gqlgen:"fieldName"` tags in your models
# struct_tag: json

# Optional: turn on to use []Thing instead of []*Thing
# omit_slice_element_pointers: false

# Optional: set to speed up generation time by not performing a final validation pass.
# skip_validation: true

# gqlgen will search for any type names in the schema in these go packages
# if they match it will use them, otherwise it will generate them.
autobind:
  - 'github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway'

# This section declares type mapping between the GraphQL and go type systems
#
# The first line in each type will be used as defaults for resolver arguments and
# modelgen, the others will be allowed when binding to fields. Configure them to
# your liking
models:
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.ID
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
  Int:
    model:
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
```
### 6.3. Get gqlgen
```shell
go get -d github.com/99designs/gqlgen
```
### 6.3. Gen
```shell
go run github.com/99designs/gqlgen generate
```
### 6.4. graphqlHandler
1. Add Resolver for graph
```go
// graph/resolver/resolver.go
import (
pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
)
type Resolver struct{
	AuthClient pb.GreeterClient
}
```
2. Map client in resolver
```go
func graphqlHandler() gin.HandlerFunc {
	// Set up a connection to the server.
	authConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(authConn)

	authClient := pb.NewGreeterClient(authConn)

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		AuthClient: authClient,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
```
3. Use graphqlHandler
```go
func main() {
	// Set up a http server.
	r := gin.Default()
	r.POST("/graphql", graphqlHandler())
	r.GET("/graphql", playgroundHandler())
	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		panic("Error")
	}
}
```
4. GraphqlPlayground
```go
// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
```