const { mergeResolvers } = require('merge-graphql-schemas');
const generatedResolvers = require('./generated/resolvers.js');
const modules = require('./modules')

// Resolvers object is passed grpc client to use for data fetching.
const resolvers = client => ({
    Field: {
        // If there is a DataPath value in the field.
        // Fetch the Membership entries from the CMS and return as json on DataPathValue field.
        DataPathValue: async ({ DataPath }, _, ctx) => {
            if (DataPath && ctx.cause) {
                const data = await modules[DataPath].getFieldValue(client, ctx.cause)
                return data
            }
            return null
        },
    },
    Query: {
        getCause: async (_, { ID }, ctx) => {
            ctx.cause = ID;
            const data = await client
                .GetCause()
                .sendMessage({ api: 'v1', ID })
            return data.item
        },
    },
    Mutation: {
        newCause: async (_, { Title }, ctx) => {
            const { item: layoutId = ID } = await client
                .CreateLayout()
                .sendMessage({ api: 'v1', ID })
            const { item: homePageID = ID } = await client
                .CreateHomePage()
                .sendMessage({ api: 'v1', ID })
            const { item: photoId = ID } = await client
                .CreatePhoto()
                .sendMessage({ api: 'v1', ID })
            const { item: causeID = ID } = await client
                .CreatePhoto()
                .sendMessage({ api: 'v1', ID })
            return causeID
        },
    }
})


// Merge the generated resolvers with our custom resolvers and export, passing gRPC client along.
const resolversList = client => [generatedResolvers(client), resolvers(client)]
module.exports = client => (mergeResolvers(resolversList(client)));
