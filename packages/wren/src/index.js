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

// Import generated API
const generatedSchema = require('./schema.js');

const generatedResolvers = require('./resolvers.js');
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

const typeDefs = generatedSchema;

const resolvers = generatedResolvers(client);

// In the most basic sense, the ApolloServer can be started
// by passing type definitions (typeDefs) and the resolvers
// responsible for fetching the data for those types.
const server = new ApolloServer({ typeDefs, resolvers });

// This `listen` method launches a web-server.  Existing apps
// can utilize middleware options, which we'll discuss later.
server.listen().then(({ url }) => {
  console.log(`🚀  Server ready at ${url}`);
});
