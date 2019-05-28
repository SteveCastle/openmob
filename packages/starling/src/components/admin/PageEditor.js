import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import gql from 'graphql-tag'
import ThemeProvider from '@openmob/bluebird/src/ThemeProvider'
import skyward from '@openmob/bluebird/src/themes/skyward'
import Row from '@openmob/bluebird/src/components/layout/Row'
import Column from '@openmob/bluebird/src/components/layout/Column'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Node from '../Node'
import SEO from '../SEO'

const sortByWeight = (a, b) => a.Weight - b.Weight
const GET_PAGE = gql`
  query HomePageQuery($id: ID!) {
    getHomePage(ID: $id) {
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
                IntValue
                FloatValue
                BooleanValue
                DateTimeValue {
                  seconds
                }
                DataPathValue
                DataPath
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
`

function PageEditor({ navigate = () => { }, pageID }) {
    const {
        data: { getHomePage: page = {} },
        error,
        loading,
    } = useQuery(GET_PAGE, {
        variables: {
            id: pageID,
        },
    })
    console.log('page', page)

    if (loading) {
        return <div>Loading...</div>
    }

    if (error) {
        return <div>Error! {error.message}</div>
    }

    return (
        <ThemeProvider theme={skyward}>
            <Content top>
                <div style={{ minHeight: "min-content", display: "flex", flexWrap: "wrap" }}>
                    <SEO
                        title={page.Title}
                        keywords={[`gatsby`, `application`, `react`]}
                    />
                    {(page.Layout.LayoutRows || []).sort(sortByWeight).map(row => (
                        <Row key={row.ID} container={row.Container}>
                            {(row.LayoutColumns || []).sort(sortByWeight).map(column => (
                                <Column key={column.ID} size={column.Width}>
                                    {(column.Components || []).sort(sortByWeight).map(component => (
                                        <Node
                                            id={page.ID}
                                            fields={component.Fields}
                                            path={component.ComponentImplementation.Path}
                                            key={component.ID}
                                        />
                                    ))}
                                </Column>
                            ))}
                        </Row>
                    ))}
                </div>
            </Content>
        </ThemeProvider>
    )
}

PageEditor.propTypes = {
    navigate: PropTypes.func,
}

export default PageEditor
