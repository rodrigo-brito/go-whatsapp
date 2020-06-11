import React from "react";
import ReactDOM from "react-dom";
import Chat from "./chat";
import { ApolloProvider } from "@apollo/react-hooks";
import { split, InMemoryCache } from "apollo-boost";
import { WebSocketLink } from "apollo-link-ws";
import { HttpLink } from "apollo-link-http";
import { getMainDefinition } from "apollo-utilities";
import { ApolloClient } from "apollo-client";

// API URL
let apiURL = "chat.brito.com.br";
if (process.env.NODE_ENV !== "production") {
  apiURL = "localhost:8080";
}

const httpLink = new HttpLink({
  uri: `http://${apiURL}/graphql`,
});

const wsLink = new WebSocketLink({
  uri: `ws://${apiURL}/graphql`,
  options: {
    reconnect: true,
  },
});

const link = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === "OperationDefinition" &&
      definition.operation === "subscription"
    );
  },
  wsLink,
  httpLink
);

const client = new ApolloClient({
  link,
  cache: new InMemoryCache(),
});

ReactDOM.render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <Chat />
    </ApolloProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
