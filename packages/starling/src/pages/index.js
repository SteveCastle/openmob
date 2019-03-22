import React from "react"
import PropTypes from "prop-types"
import { graphql, Link } from "gatsby"

import Layout from "../components/layout"
import SEO from "../components/seo"

const IndexPage = ({
  data: {
    wren: { listCause: causes = [] },
  },
}) => (
  <Layout title={"grassroots.dev"}>
    <SEO title="Home" keywords={[`gatsby`, `application`, `react`]} />
    <h1>List of Causes</h1><ul>
    {causes.map(cause => (
      <li>
      <Link to={`/${cause.Slug}`}>{cause.Title}</Link>
      </li>
    ))}</ul>
    <Link to="/admin">Go to the admin page</Link>
  </Layout>
)

export const pageQuery = graphql`
  query IndexQuery {
    wren {
      listCause(limit: 10) {
        ID
        Title
        Slug
    }
  }
}
`
IndexPage.propTypes = {
  data: PropTypes.shape({
    wren: PropTypes.shape({
      getCause: PropTypes.shape({}),
    }),
  }),
}

export default IndexPage
