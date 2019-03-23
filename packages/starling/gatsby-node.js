const path = require(`path`)

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