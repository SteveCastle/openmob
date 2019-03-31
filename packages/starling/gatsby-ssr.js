import React from "react"
import ApolloClient from "apollo-boost"
import fetch from "isomorphic-fetch"
import { ApolloProvider } from "react-apollo-hooks"

const client = new ApolloClient({
  uri: "http://localhost:4000",
  fetch,
})

export const wrapRootElement = ({ element }) => (
  <ApolloProvider client={client}>{element}</ApolloProvider>
)
