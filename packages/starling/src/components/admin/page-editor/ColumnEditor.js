import React, { useState } from 'react';
import { useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import Column from '@openmob/bluebird/src/components/layout/Column';
import Overlay from '@openmob/bluebird/src/components/editor/Overlay';
import Control from '@openmob/bluebird/src/components/editor/Control';
import Widget from '@openmob/bluebird/src/components/editor/Widget';

import GET_PAGE from '../../../queries/getPage';

const UPDATE_COLUMN = gql`
  mutation updateLayoutColumn($id: ID!, $layoutColumn: LayoutColumnInput) {
    updateLayoutColumn(ID: $id, layoutColumn: $layoutColumn, buildStatic: true)
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

  const handleDelete = () => () =>
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
    });

  const createCreateComponent = () => () =>
    createComponent({
      variables: {
        component: {
          LayoutColumn: column.ID,
          ComponentType: '32d88391-4fc4-4a6d-beaf-1d5051da4db5',
          ComponentImplementation: 'd5721029-93e2-49a9-b798-21aff3a11c2c',
        },
        buildStatic: true,
      },
      refetchQueries: [
        {
          query: GET_PAGE,
          variables: { id: pageId },
        },
      ],
    });

  const handleChangeWidth = newWidth => () =>
    updateLayoutColumn({
      variables: {
        id: column.ID,
        layoutColumn: {
          Width: parseInt(newWidth),
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
    });

  return (
    <Column size={size} disableSpacing>
      <Overlay
        locked={locked}
        onClick={() => setLock(!locked)}
        nestingLevel={1}
      >
        <Control label="Delete Column">
          <Widget handleSubmit={handleDelete} />
        </Control>
        <Control label="Create Component">
          <Widget handleSubmit={createCreateComponent} />
        </Control>
        <Control label="Change Width">
          <Widget
            handleSubmit={handleChangeWidth}
            options={[
              { ID: 12, Title: '12' },
              { ID: 6, Title: '6' },
              { ID: 3, Title: '3' },
            ]}
          />
        </Control>
      </Overlay>
      {children}
    </Column>
  );
}

export default ColumnEditor;
