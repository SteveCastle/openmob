const path = require(`path`)
const { GraphQLString } = require("gatsby/graphql")

exports.createPages = ({ graphql, actions }) => {
  const { createPage } = actions
  const pageTemplate = path.resolve(`src/templates/Page.js`)
  // Query for markdown nodes to use in creating pages.
  // You can query for whatever data you want to create pages for e.g.
  // products, portfolio items, landing pages, etc.
  return graphql(`
  {wren{
    listCause(limit: 50){
      ID
      Title
      Slug
    }
  }}
  `).then(result => {
    if (result.errors) {
      throw result.errors
    }

    // Create blog post pages.
    (result.data.wren.listCause || []).forEach(cause => {
      console.log(cause.Slug);
      createPage({
        // Path for this page — required
        path: `${cause.Slug}`,
        component: pageTemplate,
        context: {
          id: cause.ID
        },
      })
    })
  })
}
exports.setFieldsOnGraphQLNodeType = ({ type }) => {
  if (type.name === `WREN`) {
    return {
      newField: {
        type: GraphQLString,
        args: {
          myArgument: {
            type: GraphQLString,
          }
        },
        resolve: (source, fieldArgs) => {
          return `Id of this node is ${source.id}.
                  Field was called with argument: ${fieldArgs.myArgument}`
        }
      }
    }
  }

  // by default return empty object
  return {}
}