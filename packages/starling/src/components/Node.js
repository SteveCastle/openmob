import React, { lazy, Suspense } from 'react'
import ErrorBoundary from './ErrorBoundary'
const getComponent = path =>
  lazy(() => import(`@openmob/bluebird/src/components${path}`))

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
      <Suspense maxDuration={500} fallback={<div />}>
        <Component {...getProps(fields)} key={id} />
      </Suspense>
    </ErrorBoundary>
  )
}

export default Node
