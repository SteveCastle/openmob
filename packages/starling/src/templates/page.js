import React from "react"
import PropTypes from "prop-types"
import { graphql, Link } from "gatsby"
import ThemeProvider from "@openmob/bluebird/src/ThemeProvider"
import skyward from "@openmob/bluebird/src/themes/skyward"

import Row from "@openmob/bluebird/src/components/layout/Row"
import Column from "@openmob/bluebird/src/components/layout/Column"

import Layout from "../components/Layout"
import Node from "../components/Node"

import SEO from "../components/SEO"

const sortByWeight = (a, b) => a.Weight - b.Weight

const IndexPage = ({
  data: {
    wren: { getCause: cause = {} },
  },
}) => (
  <ThemeProvider theme={skyward}>
    <Layout title={cause.Title} id={cause.ID} summary={cause.Summary}>
      <SEO title={cause.Title} keywords={[`gatsby`, `application`, `react`]} />
      {(cause.HomePage.Layout.LayoutRows || []).sort(sortByWeight).map(row => (
        <Row key={row.ID} container={row.Container}>
          {(row.LayoutColumns || []).sort(sortByWeight).map(column => (
            <Column key={column.ID} size={column.Width}>
              {(column.Components || []).sort(sortByWeight).map(component => (
                <Node
                  fields={component.Fields}
                  path={component.ComponentImplementation.Path}
                  key={component.ID}
                />
              ))}
            </Column>
          ))}
        </Row>
      ))}
      <Link to="/app">Go to the admin page</Link>
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
              Container
              Weight
              LayoutColumns {
                ID
                Width
                Weight
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
