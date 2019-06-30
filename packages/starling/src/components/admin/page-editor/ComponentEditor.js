import React, { useState } from 'react';
import { useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import Overlay from '@openmob/bluebird/src/components/editor/Overlay';
import Control from '@openmob/bluebird/src/components/editor/Control';
import Widget from '@openmob/bluebird/src/components/editor/Widget';

import GET_PAGE from '../../../queries/getPage';

// GraphQL Queries to perform the actions of this editor.
const DELETE_COMPONENT = gql`
  mutation deleteComponent($id: ID!) {
    deleteComponent(ID: $id, buildStatic: true)
  }
`;

const UPDATE_COMPONENT = gql`
  mutation updateComponent($id: ID!, $component: ComponentInput) {
    updateComponent(ID: $id, component: $component, buildStatic: true)
  }
`;

// UI for editing a component node in the layout tree.
function ComponentEditor({
  children,
  component,
  pageId,
  columnId,
  componentType,
}) {
  const [locked, setLock] = useState(false);
  const deleteComponent = useMutation(DELETE_COMPONENT);
  const updateComponent = useMutation(UPDATE_COMPONENT);
  // Function to delete the component from the row.
  const removeComponent = () => () =>
    deleteComponent({
      variables: {
        id: component.ID,
        buildStatic: true,
      },
      refetchQueries: [
        {
          query: GET_PAGE,
          variables: { id: pageId },
        },
      ],
    });
  // Callback function to change the component implementation to
  // a new ID. Passed to widget that calls function with result value of widget.
  const changeImplementation = newID => () =>
    updateComponent({
      variables: {
        id: component.ID,
        component: {
          ComponentType: component.ComponentType.ID,
          ComponentImplementation: newID,
          LayoutColumn: columnId,
          Wieght: component.Weight,
        },
      },
      refetchQueries: [
        {
          query: GET_PAGE,
          variables: { id: pageId },
        },
      ],
    });
  return (
    <div style={{ width: '100%', position: 'relative' }}>
      <Overlay locked={locked} onClick={() => setLock(!locked)}>
        <Control label="Delete">
          <Widget handleSubmit={removeComponent} />
        </Control>
        <Control
          label="Edit Fields"
          options={componentType.ComponentTypeFieldss}
        >
          <Widget
            handleSubmit={() => console.log(componentType, component.Fields)}
          />
        </Control>
        <Control
          label="Change Style"
          options={componentType.ComponentImplementations}
        >
          <Widget
            handleSubmit={changeImplementation}
            options={componentType.ComponentImplementations}
          />
        </Control>
      </Overlay>
      {children}
    </div>
  );
}

export default ComponentEditor;
