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

const CREATE_COMPONENT = gql`
  mutation createComponent($component: ComponentInput!) {
    createComponent(component: $component, buildStatic: true) {
      ID
    }
  }
`;

function ColumnEditor({ children, size, column, pageId, rowId }) {
  const [locked, setLock] = useState(false);
  const deleteLayoutColumn = useMutation(DELETE_COLUMN);
  const updateLayoutColumn = useMutation(UPDATE_COLUMN);
  const createComponent = useMutation(CREATE_COMPONENT);

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
        <Control
          onClick={() => () =>
            createComponent({
              variables: {
                component: {
                  LayoutColumn: column.ID,
                  ComponentType: '32d88391-4fc4-4a6d-beaf-1d5051da4db5',
                  ComponentImplementation:
                    'd5721029-93e2-49a9-b798-21aff3a11c2c',
                },
                buildStatic: true,
              },
              refetchQueries: [
                {
                  query: GET_PAGE,
                  variables: { id: pageId },
                },
              ],
            })}
        >
          Add Component
        </Control>
        <Control
          onClick={() => () =>
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
            })}
        >
          Change Width
        </Control>
      </Overlay>
      {children}
    </Column>
  );
}

export default ColumnEditor;
