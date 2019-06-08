import React, { useState } from 'react';
import Row from '@openmob/bluebird/src/components/layout/Row';
import Overlay from '@openmob/bluebird/src/components/editor/Overlay';
import Control from '@openmob/bluebird/src/components/editor/Control';

function RowEditor({ children, deleteRow, addColumn, makeContainer }) {
  const [locked, setLock] = useState(false);
  return (
    <Row disableSpacing>
      <Overlay locked={locked} onClick={() => setLock(!locked)}>
        <Control onClick={deleteRow}>Delete</Control>
        <Control onClick={addColumn}>Add Column</Control>
        <Control onClick={makeContainer}>Make Container</Control>
      </Overlay>

      {children}
    </Row>
  );
}

export default RowEditor;
