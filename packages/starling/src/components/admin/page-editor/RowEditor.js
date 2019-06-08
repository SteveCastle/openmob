import React, { useState } from 'react';
import { useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';

import Row from '@openmob/bluebird/src/components/layout/Row';
import Overlay from '@openmob/bluebird/src/components/editor/Overlay';
import Control from '@openmob/bluebird/src/components/editor/Control';

const UPDATE_ROW = gql`
  mutation updateLayoutRow($id: ID!, $layoutRow: LayoutRowInput) {
    updateLayoutRow(ID: $id, layoutRow: $layoutRow, buildStatic: true)
  }
`;

const DELETE_ROW = gql`
  mutation deleteLayoutRow($id: ID!) {
    deleteLayoutRow(ID: $id, buildStatic: true)
  }
`;

const CREATE_COLUMN = gql`
  mutation createLayoutColumn($layoutColumn: LayoutColumnInput) {
    createLayoutColumn(layoutColumn: $layoutColumn, buildStatic: true) {
      ID
    }
  }
`;

function RowEditor({ children, row, layoutId }) {
  const [locked, setLock] = useState(false);
  const updateLayoutRow = useMutation(UPDATE_ROW);
  const deleteLayoutRow = useMutation(DELETE_ROW);
  const createColumn = useMutation(CREATE_COLUMN);

  return (
    <Row disableSpacing>
      <Overlay locked={locked} onClick={() => setLock(!locked)}>
        <Control
          onClick={() =>
            deleteLayoutRow({
              variables: {
                id: row.ID,
                buildStatic: true,
              },
            })
          }
        >
          Delete
        </Control>
        <Control
          onClick={() =>
            updateLayoutRow({
              variables: {
                id: row.ID,
                layoutRow: {
                  Weight: row.Weight,
                  Layout: layoutId,
                  Container: !row.Container,
                },
                buildStatic: true,
              },
            })
          }
        >
          Toggle Container
        </Control>
        <Control onClick={createColumn}>Add Column</Control>
      </Overlay>

      {children}
    </Row>
  );
}

export default RowEditor;
