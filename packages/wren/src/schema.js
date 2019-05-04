const { mergeTypes } = require('merge-graphql-schemas');
const generatedSchema = require('./generated/schema.js');

const schema = `
type Field {
    DataPathValue: String
}
type Query {
    getHomePageLayout(ID: ID!): Layout
}
`

const types = [
    generatedSchema,
    schema,
];

module.exports = mergeTypes(types, { all: true });
