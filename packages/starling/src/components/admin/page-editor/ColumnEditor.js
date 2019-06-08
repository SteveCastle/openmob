import React from 'react';
import Column from '@openmob/bluebird/src/components/layout/Column';

function ColumnEditor({ children, size }) {
  return (
    <Column size={size} disableSpacing>
      {children}
    </Column>
  );
}

export default ColumnEditor;
