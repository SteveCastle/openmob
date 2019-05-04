const generatedResolvers = require('./generated/resolvers.js');
const { mergeResolvers } = require('merge-graphql-schemas');


// Resolvers object is passed grpc client to use for data fetching.
const resolvers = client => ({
    Field: {
        // If there is a DataPath value in the field.
        // Fetch the Membership entries from the CMS and return as json on DataPathValue field.
        DataPathValue: async ({ DataPath }, _, ctx) => {
            if (DataPath && ctx.cause) {
                const query = DataPath.split('.')
                const memberData = await client[query[0]]()
                    .sendMessage({ api: 'v1', filters: [{ Cause: ctx.cause }], limit: 10 })
                const memberItems = memberData.items
                const fetchItems = async item => await client[`Get${query[1]}`]()
                    .sendMessage({ api: 'v1', ID: item[query[1]] });
                const data = async () => {
                    return await Promise.all(memberItems.map(item => fetchItems(item)))
                }
                return JSON.stringify(data.item)
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
    }
})


// Merge the generated resolvers with our custom resolvers and export, passing gRPC client along.
const resolversList = client => [generatedResolvers(client), resolvers(client)]
module.exports = client => (mergeResolvers(resolversList(client)));
