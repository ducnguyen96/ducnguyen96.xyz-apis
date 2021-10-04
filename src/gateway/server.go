package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/graph/generated"
	graph "github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/graph/resolver"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

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
	r.POST("/graphql", graphqlHandler())
	r.GET("/graphql", playgroundHandler())
	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		panic("Error")
	}
}