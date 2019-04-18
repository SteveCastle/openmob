import React from "react"
import PropTypes from "prop-types"
import { graphql, Link } from "gatsby"
import ThemeProvider from "@openmob/bluebird/src/ThemeProvider"
import skyward from "@openmob/bluebird/src/themes/skyward"
import ImageGrid from "@openmob/bluebird/src/components/lists/ImageGrid"
import GridItem from "@openmob/bluebird/src/components/lists/GridItem"
import Container from "@openmob/bluebird/src/components/layout/Container"

import Layout from "../components/Layout"
import SEO from "../components/SEO"

const IndexPage = ({
  data: {
    wren: { listCause: causes = [] },
  },
}) => (
  <ThemeProvider theme={skyward}>
    <Layout title="grassroots.dev" id="List view" summary="Debug Mode">
      <SEO title="Home" keywords={[`gatsby`, `application`, `react`]} />
      <Container>
        <ImageGrid>
          {(causes || []).map(cause => (
            <Link to={`/${cause.Slug}`} key={cause.ID}>
              <GridItem
                title={cause.Title}
                caption={cause.Summary}
                uri={cause.Photo.URI}
              />
            </Link>
          ))}
        </ImageGrid>
      </Container>
    </Layout>
  </ThemeProvider>
)

export const pageQuery = graphql`
  query IndexQuery {
    wren {
      listCause(limit: 50) {
        ID
        Title
        Summary
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
