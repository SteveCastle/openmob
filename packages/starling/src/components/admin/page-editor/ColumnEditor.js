import React, { useState } from 'react';
import { useMutation, useQuery } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import Column from '@openmob/bluebird/src/components/layout/Column';
import Overlay from '@openmob/bluebird/src/components/editor/Overlay';
import Control from '@openmob/bluebird/src/components/editor/Control';
import Widget from '@openmob/bluebird/src/components/editor/Widget';
import { GET_PAGE } from '../../../queries/getPage';
import { LIST_COMPONENT_TYPES } from '../../../queries/listComponentTypes';

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
  mutation newComponent(
    $LayoutColumn: ID!
    $ComponentType: ID!
    $ComponentImplementation: ID!
  ) {
    newComponent(
      LayoutColumn: $LayoutColumn
      ComponentType: $ComponentType
      ComponentImplementation: $ComponentImplementation
    )
  }
`;

function ColumnEditor({ children, size, column, pageId, rowId }) {
  const [locked, setLock] = useState(false);
  const deleteLayoutColumn = useMutation(DELETE_COLUMN);
  const updateLayoutColumn = useMutation(UPDATE_COLUMN);
  const newComponent = useMutation(CREATE_COMPONENT);
  const {
    data: { listComponentType: componentTypes = [] },
    error,
    loading,
  } = useQuery(LIST_COMPONENT_TYPES);

  if (loading) {
    return <div />;
  }

  if (error) {
    return <div>Error! {error.message}</div>;
  }

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

  const handleCreateComponent = newID => () =>
    newComponent({
      variables: {
        LayoutColumn: column.ID,
        ComponentType: componentTypes.find(component => {
          return component.ID === newID;
        }).ID,
        ComponentImplementation: componentTypes.find(component => {
          return component.ID === newID;
        }).ComponentImplementation.ID,
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
          <Widget
            handleSubmit={handleCreateComponent}
            options={componentTypes}
          />
        </Control>
        <Control label="Change Width">
          <Widget
            handleSubmit={handleChangeWidth}
            initValue={size}
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
