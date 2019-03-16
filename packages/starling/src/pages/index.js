import React from "react"
import PropTypes from "prop-types"
import { graphql, Link } from "gatsby"

import Layout from "../components/layout"
import SEO from "../components/seo"

const IndexPage = ({
  data: {
    wren: { getCause: cause = {} },
  },
}) => (
  <Layout title={cause.Title}>
    <SEO title="Home" keywords={[`gatsby`, `application`, `react`]} />
    <h1>{cause.ID}</h1>
    <p>{cause.Summary}</p>
    {cause.HomePage.Layout.LayoutRows.map(row => (
      <div />
    ))}
    <Link to="/admin">Go to the admin page</Link>
  </Layout>
)

export const pageQuery = graphql`
  query IndexQuery {
    wren {
      getCause(ID: "503c9ea5-6fc7-4954-b0ed-9aea35877c44") {
        ID
        Title
        Slug
        Summary
        HomePage {
          ID
          Title
          Layout {
            ID
            LayoutRows {
              ID
              LayoutColumns {
                ID
                Components {
                  ID
                  ComponentType {
                    Title
                  }
                  ComponentImplementation {
                    Title
                  }
                }
              }
            }
            LayoutType {
              ID
            }
          }
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
