import React from 'react'
import Loadable from 'react-loadable'
import ErrorBoundary from './ErrorBoundary'

const getComponent = path =>
  Loadable({
    loader: () => import(`@openmob/bluebird/src/components${path}`),
    loading: () => <div>loading</div>,
  })

const getProps = fields =>
  fields.reduce(
    (acc, field) =>
      Object.assign(acc, { [field.FieldType.PropName]: { ...field } }),
    {}
  )

function Node({ path, fields, id }) {
  const Component = getComponent(path)
  return (
    <ErrorBoundary>
      <Component {...getProps(fields)} key={id} />
    </ErrorBoundary>
  )
}

export default Node
