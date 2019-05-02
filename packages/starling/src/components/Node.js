import React from 'react'
import Loadable from 'react-loadable'
import { useStaticQuery, graphql } from 'gatsby'
import ErrorBoundary from './ErrorBoundary'
import Spinner from '@openmob/bluebird/src/components/loaders/Spinner'
const getFieldValue = field => {
  switch (field.FieldType.DataType) {
    case 'string':
      return field.StringValue
    case 'graphql':
      return field.DataPath
    default:
      return field.StringValue
  }
}

const getComponent = path =>
  Loadable({
    loader: () => import(`@openmob/bluebird/src/components/elements${path}`),
    loading: () => <Spinner />,
  })

const getProps = fields =>
  (fields || []).reduce(
    (acc, field) =>
      Object.assign(acc, { [field.FieldType.PropName]: getFieldValue(field) }),
    {}
  )

function Node({ path, fields, id }) {
  const data = useStaticQuery(
    graphql`
      query {
        wren {
          getCause(ID: "c75db1d7-17ac-47b1-aac2-879a6761b4cc") {
            ID
            ElectionMemberships {
              Election {
                ID
                Title
              }
            }
          }
        }
      }
    `
  )
  const Component = getComponent(path)
  return (
    <ErrorBoundary>
      <Component {...getProps(fields)} key={id} data={data} />
    </ErrorBoundary>
  )
}

export default Node
