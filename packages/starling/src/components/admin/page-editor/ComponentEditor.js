import React, { useState } from 'react';
import Overlay from '@openmob/bluebird/src/components/editor/Overlay';
import Control from '@openmob/bluebird/src/components/editor/Control';

function ComponentEditor({
  children,
  deleteComponent,
  changeFields,
  changeImplementation,
}) {
  const [locked, setLock] = useState(false);
  return (
    <div style={{ width: '100%' }}>
      <Overlay locked={locked} onClick={() => setLock(!locked)}>
        <Control onClick={deleteComponent}>Delete</Control>
        <Control onClick={changeFields}>Change Fields</Control>
        <Control onClick={changeImplementation}>
          Change Component Implementation
        </Control>
      </Overlay>
      {children}
    </div>
  );
}

export default ComponentEditor;
