import React, { useState } from 'react';
import { useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import Overlay from '@openmob/bluebird/src/components/editor/Overlay';
import Control from '@openmob/bluebird/src/components/editor/Control';
import GET_PAGE from '../../../queries/getPage';

const DELETE_COMPONENT = gql`
  mutation deleteComponent($id: ID!) {
    deleteComponent(ID: $id, buildStatic: true)
  }
`;

function ComponentEditor({ children, component, pageId }) {
  const [locked, setLock] = useState(false);
  const deleteComponent = useMutation(DELETE_COMPONENT);

  return (
    <div style={{ width: '100%' }}>
      <Overlay locked={locked} onClick={() => setLock(!locked)}>
        <Control
          onClick={() =>
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
            })
          }
        >
          Delete
        </Control>
        <Control onClick={() => console.log('change fields')}>
          Change Fields
        </Control>
        <Control onClick={() => console.log('change component implementation')}>
          Change Component Implementation
        </Control>
      </Overlay>
      {children}
    </div>
  );
}

export default ComponentEditor;
