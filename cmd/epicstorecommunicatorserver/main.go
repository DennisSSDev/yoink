package main

import (
	"net/http"

	"github.com/dennisssdev/yoink/internal/epicstorecommunicatorserver"
	"github.com/dennisssdev/yoink/rpc/epicstorecommunicator"
	"github.com/shurcooL/graphql"
)

// EpicStoreGQLEndpoint is utilized for the gql client
const EpicStoreGQLEndpoint = "https://graphql.epicgames.com/graphql"

func main() {
	server := &epicstorecommunicatorserver.Server{GraphQLClient: graphql.NewClient(EpicStoreGQLEndpoint, nil)}
	twirpEpicStoreHandler := epicstorecommunicator.NewEpicStoreCommunicatorServer(server, nil)

	http.ListenAndServe(":8034", twirpEpicStoreHandler)
}
