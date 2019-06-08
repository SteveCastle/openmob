import React, { useState } from 'react';
import Column from '@openmob/bluebird/src/components/layout/Column';
import Overlay from '@openmob/bluebird/src/components/editor/Overlay';
import Control from '@openmob/bluebird/src/components/editor/Control';

function ColumnEditor({
  children,
  size,
  deleteColumn,
  addComponent,
  changeWidth,
}) {
  const [locked, setLock] = useState(false);
  return (
    <Column size={size} disableSpacing>
      <Overlay locked={locked} onClick={() => setLock(!locked)}>
        <Control onClick={deleteColumn}>Delete</Control>
        <Control onClick={addComponent}>Add Component</Control>
        <Control onClick={changeWidth}>Change Width</Control>
      </Overlay>
      {children}
    </Column>
  );
}

export default ColumnEditor;
