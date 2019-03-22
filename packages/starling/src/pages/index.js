import React from "react"
import PropTypes from "prop-types"
import { graphql, Link } from "gatsby"
import Row from "@openmob/bluebird/src/components/layout/Row"
import Column from "@openmob/bluebird/src/components/layout/Column"

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
      <Row tracing={5} key={row.ID}>
        {row.LayoutColumns.map(column => (
          <Column tracing={5} key={column.ID} width={column.Width}>
            {column.Components.map(component => (
              <div>{component.ID}</div>
            ))}
          </Column>
        ))}
      </Row>
    ))}
    <Link to="/admin">Go to the admin page</Link>
  </Layout>
)

export const pageQuery = graphql`
  query IndexQuery {
    wren {
      getCause(ID: "6d48aea0-6c8e-44ce-bd2d-01db5cc444e2") {
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
                  Fields{
                    ID
                    FieldType{
                      DataType
                    }
                    StringValue
                  }
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
