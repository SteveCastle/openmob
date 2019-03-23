import React from "react"
import PropTypes from "prop-types"
import { graphql, Link } from "gatsby"
import ThemeProvider from "@openmob/bluebird/src/ThemeProvider"
import skyward from "@openmob/bluebird/src/themes/skyward"
import Layout from "../components/layout"
import SEO from "../components/seo"

const IndexPage = ({
  data: {
    wren: { listCause: causes = [] },
  },
}) => (
  <ThemeProvider theme={skyward}>
  <Layout title="grassroots.dev" id="List view" summary="Debug Mode">
    <SEO title="Home" keywords={[`gatsby`, `application`, `react`]} />
    <h1>List of Causes</h1><ul>
    {causes.map(cause => (
      <li>
      <Link to={`/${cause.Slug}`}>{cause.Title}</Link>
      </li>
    ))}</ul>
    <Link to="/admin">Go to the admin page</Link>
  </Layout>
  </ThemeProvider>
)

export const pageQuery = graphql`
  query IndexQuery {
    wren {
      listCause(limit: 50) {
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
