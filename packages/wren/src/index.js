const path = require('path');
const { ApolloServer, gql } = require('apollo-server');
const protoLoader = require('@grpc/proto-loader');
const grpc = require('grpc');
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

client.CreateCause(
  { api: 'v1', item: { title: 'Cool Cause', summary: 'Cool Summary' } },
  function(err, response) {
    if (err) {
      console.error(err);
      return;
    }
    console.log('Greeting:', response);
  }
);

// This is a (sample) collection of books we'll be able to query
// the GraphQL server for.  A more complete example might fetch
// from an existing data source like a REST API or database.
const books = [
  {
    title: 'Harry Potter and the Chamber of Secrets',
    author: 'J.K. Rowling'
  },
  {
    title: 'Jurassic Park',
    author: 'Michael Crichton'
  }
];

// Type definitions define the "shape" of your data and specify
// which ways the data can be fetched from the GraphQL server.
const typeDefs = gql`
  # Comments in GraphQL are defined with the hash (#) symbol.

  # This "Book" type can be used in other type declarations.
  type Book {
    title: String
    author: String
  }

  # The "Query" type is the root of all GraphQL queries.
  # (A "Mutation" type will be covered later on.)
  type Query {
    books: [Book]
  }
`;

// Resolvers define the technique for fetching the types in the
// schema.  We'll retrieve books from the "books" array above.
const resolvers = {
  Query: {
    books: () => books
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
