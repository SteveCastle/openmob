import React from 'react';
import Loadable from 'react-loadable';
import ErrorBoundary from './ErrorBoundary';
import Spinner from '@openmob/bluebird/src/components/loaders/Spinner';

// Return the field data to map to props given the field types data type.
const getFieldValue = field => {
  switch (field.FieldType.DataType) {
    case 'string':
      return field.StringValue;
    case 'module':
      return JSON.parse(field.DataPathValue);
    default:
      return field.StringValue;
  }
};

const getComponent = path =>
  Loadable({
    loader: () => import(`@openmob/bluebird/src/components/elements${path}`),
    loading: () => <Spinner />,
  });

const getProps = fields =>
  (fields || []).reduce(
    (acc, field) =>
      Object.assign(acc, { [field.FieldType.PropName]: getFieldValue(field) }),
    {}
  );

function Node({ path, fields, id }) {
  const Component = getComponent(path);
  return (
    <ErrorBoundary>
      <Component {...getProps(fields)} key={id} />
    </ErrorBoundary>
  );
}

export default Node;
