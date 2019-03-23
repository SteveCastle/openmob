import React from "react"
import PropTypes from "prop-types"
import { graphql, Link } from "gatsby"
import ThemeProvider from "@openmob/bluebird/src/ThemeProvider"
import skyward from "@openmob/bluebird/src/themes/skyward"

import Row from "@openmob/bluebird/src/components/layout/Row"
import Column from "@openmob/bluebird/src/components/layout/Column"

import Layout from "../components/layout"
import Node from "../components/Node"

import SEO from "../components/seo"

const IndexPage = ({
  data: {
    wren: { getCause: cause = {} },
  },
}) => (
  <ThemeProvider theme={skyward}>
    <Layout title={cause.Title} id={cause.ID} summary={cause.Summary}>
      <SEO title={cause.Title} keywords={[`gatsby`, `application`, `react`]} />
      {cause.HomePage.Layout.LayoutRows.map(row => (
        <Row key={row.ID}>
          {row.LayoutColumns.map(column => (
            <Column key={column.ID} width={column.Width}>
              {column.Components.map(component => (
                <Node
                  fields={component.Fields}
                  path={component.ComponentImplementation.Path}
                  id={component.ID}
                />
              ))}
            </Column>
          ))}
        </Row>
      ))}
      <Link to="/admin">Go to the admin page</Link>
    </Layout>
  </ThemeProvider>
)

export const pageQuery = graphql`
  query LandingPageQuery($id: ID!) {
    wren {
      getCause(ID: $id) {
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
                  Fields {
                    ID
                    FieldType {
                      Title
                      DataType
                      PropName
                    }
                    StringValue
                  }
                  ComponentImplementation {
                    Path
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
