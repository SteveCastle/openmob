import React, { useState } from 'react';
import { useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import Overlay from '@openmob/bluebird/src/components/editor/Overlay';
import Control from '@openmob/bluebird/src/components/editor/Control';
import Widget from '@openmob/bluebird/src/components/editor/Widget';
import TextWidget from '@openmob/bluebird/src/components/editor/TextWidget';
import ConfirmWidget from '@openmob/bluebird/src/components/editor/ConfirmWidget';

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

const UPDATE_FIELD = gql`
  mutation updateField($id: ID!, $field: FieldInput) {
    updateField(ID: $id, field: $field, buildStatic: true)
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
  const updateField = useMutation(UPDATE_FIELD);

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

  const handleUpdateField = field => newValue => () =>
    updateField({
      variables: {
        id: field.ID,
        field: {
          ...field,
          Component: component.ID,
          FieldType: field.FieldType.ID,
          ID: undefined,
          CreatedAt: undefined,
          UpdatedAt: undefined,
          __typename: undefined,
          DataPathValue: undefined,
          StringValue: newValue,
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
      <Overlay
        locked={locked}
        onClick={() => setLock(!locked)}
        nestingLevel={2}
      >
        <Control label="Delete">
          <ConfirmWidget handleSubmit={removeComponent} />
        </Control>
        <Control
          label="Edit Fields"
          options={componentType.ComponentTypeFieldss}
        >
          {Array.isArray(component.Fields) &&
            component.Fields.map(field => (
              <TextWidget
                title={field.FieldType.Title}
                initValue={field.StringValue}
                handleSubmit={handleUpdateField(field)}
                key={field.ID}
              />
            ))}
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
