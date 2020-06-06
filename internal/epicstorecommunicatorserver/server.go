package epicstorecommunicatorserver

import "github.com/shurcooL/graphql"

// Server implements the Greeter service
type Server struct {
	GraphQLClient *graphql.Client
}
