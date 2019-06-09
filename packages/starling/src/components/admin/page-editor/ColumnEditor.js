import React, { useState } from 'react';
import { useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import Column from '@openmob/bluebird/src/components/layout/Column';
import Overlay from '@openmob/bluebird/src/components/editor/Overlay';
import Control from '@openmob/bluebird/src/components/editor/Control';
import GET_PAGE from '../../../queries/getPage';

const UPDATE_COLUMN = gql`
  mutation updateLayoutColumn($id: ID!, $layoutColumn: LayoutColumnInput) {
    updateLayoutColumn(ID: $id, layoutRow: $layoutRow, buildStatic: true)
  }
`;

const DELETE_COLUMN = gql`
  mutation deleteLayoutColumn($id: ID!) {
    deleteLayoutColumn(ID: $id, buildStatic: true)
  }
`;

function ColumnEditor({ children, size, column, pageId, rowId }) {
  const [locked, setLock] = useState(false);
  const deleteLayoutColumn = useMutation(DELETE_COLUMN);
  const updateLayoutColumn = useMutation(UPDATE_COLUMN);

  return (
    <Column size={size} disableSpacing>
      <Overlay locked={locked} onClick={() => setLock(!locked)}>
        <Control
          onClick={() =>
            deleteLayoutColumn({
              variables: {
                id: column.ID,
                buildStatic: true,
              },
              refetchQueries: [
                {
                  query: GET_PAGE,
                  variables: { id: pageId },
                },
              ],
            })
          }
        >
          Delete
        </Control>
        <Control onClick={() => console.log('add component')}>
          Add Component
        </Control>
        <Control
          onClick={() =>
            updateLayoutColumn({
              variables: {
                id: column.ID,
                layoutColumn: {
                  Width: column.Weight,
                  LayoutRow: rowId,
                },
                buildStatic: true,
              },
              refetchQueries: [
                {
                  query: GET_PAGE,
                  variables: { id: pageId },
                },
              ],
            })
          }
        >
          Change Width
        </Control>
      </Overlay>
      {children}
    </Column>
  );
}

export default ColumnEditor;
