import React from "react"
import PropTypes from "prop-types"
import { graphql, Link } from "gatsby"
import ThemeProvider from "@openmob/bluebird/src/ThemeProvider"
import skyward from "@openmob/bluebird/src/themes/skyward"
import ImageGrid from "@openmob/bluebird/src/components/lists/ImageGrid"
import GridItem from "@openmob/bluebird/src/components/lists/GridItem"

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
      <ImageGrid>
        {causes.map(cause => (
          <Link to={`/${cause.Slug}`}>
            <GridItem title={cause.Title} uri={cause.Photo.URI} />
          </Link>
        ))}
      </ImageGrid>
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
        Photo {
          URI
          Width
          Height
        }
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
