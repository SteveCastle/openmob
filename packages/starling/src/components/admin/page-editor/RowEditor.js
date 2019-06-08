import React from 'react';
import Row from '@openmob/bluebird/src/components/layout/Row';

function RowEditor({ children }) {
  return <Row disableSpacing>{children}</Row>;
}

export default RowEditor;
