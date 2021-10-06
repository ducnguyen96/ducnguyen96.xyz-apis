package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/graph/generated"
	graph "github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/graph/resolver"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/service"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net/http"
	"time"
)

func graphqlHandler() gin.HandlerFunc {
	// Set up a connection to the server.
	kacp := keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 20 seconds if there is no activity
		Timeout:             10 * time.Second, // wait 10 second for ping back
		PermitWithoutStream: true,             // send pings even without active streams
	}

	authConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	authClient := pb.NewAuthServiceClient(authConn)

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		AuthClient: authClient,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Set up a http server.
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("/graphql", graphqlHandler())
	r.GET("/graphql", playgroundHandler())
	r.POST("/api/v1/upload", service.Upload)
	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		panic("Error")
	}
}