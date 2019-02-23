const path = require('path');
const { ApolloServer, gql } = require('apollo-server');
const protoLoader = require('@grpc/proto-loader');
const grpc = require('grpc');
const grpc_promise = require('grpc-promise');
const PROTO_PATH = path.resolve(
  __dirname,
  '../../shrike/src/api/proto/v1/shrike.proto'
);
const THIRD_PARTY = path.resolve(__dirname, '../../shrike/third_party');

// Connect to Shrike gRPC server.
var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: false,
  oneofs: true,
  includeDirs: [THIRD_PARTY]
});
var shrike = grpc.loadPackageDefinition(packageDefinition).shrike.v1;
var client = new shrike.ShrikeService(
  'localhost:9090',
  grpc.credentials.createInsecure()
);

grpc_promise.promisifyAll(client);

client
  .CreateCause()
  .sendMessage({
    api: 'v1',
    item: { Title: 'Cool Cause', Summary: 'Cool Summary' }
  })
  .then(resp => console.log(resp));

const typeDefs = gql`
  type Time {
    seconds: Int!
    nanos: Int!
  }

  input CauseInput {
    Title: String
    description: String
  }

  type Cause {
    ID: Int!
    CreatedAt: Time
    UpdatedAt: Time
    Title: String
    description: String
  }

  type Query {
    getCause(ID: Int): Cause
    listCause: [Cause]
  }

  type Mutation {
    createCause(cause: CauseInput): Cause
    updateCause(ID: Int, cause: CauseInput): Int
    deleteCause(ID: Int): Int
  }
`;

const resolvers = {
  Query: {
    listCause: (_, { Limit, Cursor, Order, Filter }) =>
      client
        .ListCause()
        .sendMessage({ api: 'v1' })
        .then(res => res.items),
    getCause: (_, { ID }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item)
  },
  Mutation: {
    createCause: (_, { cause }) =>
      client
        .CreateCause()
        .sendMessage({ api: 'v1', item: { ...cause } })
        .then(res => ({ ID: res.ID, ...cause })),
    updateCause: (_, { ID, cause }) =>
      client
        .UpdateCause()
        .sendMessage({ api: 'v1', item: { ID, ...cause } })
        .then(res => res.updated),
    deleteCause: (_, { ID }) =>
      client
        .DeleteCause()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted)
  }
};

// In the most basic sense, the ApolloServer can be started
// by passing type definitions (typeDefs) and the resolvers
// responsible for fetching the data for those types.
const server = new ApolloServer({ typeDefs, resolvers });

// This `listen` method launches a web-server.  Existing apps
// can utilize middleware options, which we'll discuss later.
server.listen().then(({ url }) => {
  console.log(`🚀  Server ready at ${url}`);
});
