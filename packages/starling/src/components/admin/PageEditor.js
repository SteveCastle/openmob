import React from 'react';
import PropTypes from 'prop-types';
import { useQuery, useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import Button from '@openmob/bluebird/src/components/buttons/Button';
import Spinner from '@openmob/bluebird/src/components/loaders/Spinner';

import { Row, Column, Content } from '@openmob/bluebird/src/components/layout';
import Node from '../Node';
import Editor from './page-editor/PageEditor';
import RowEditor from './page-editor/RowEditor';
import ColumnEditor from './page-editor/ColumnEditor';
import ComponentEditor from './page-editor/ComponentEditor';
import SEO from '../SEO';
import { GET_PAGE } from '../../queries/getPage';
const sortByWeight = (a, b) => a.Weight - b.Weight;

const CREATE_ROW = gql`
  mutation createLayoutRow($layoutRow: LayoutRowInput) {
    createLayoutRow(layoutRow: $layoutRow, buildStatic: true) {
      ID
    }
  }
`;

function PageEditor({ pageID, causeID }) {
  const { data, error, loading } = useQuery(GET_PAGE, {
    variables: {
      id: causeID,
    },
  });
  const createLayoutRow = useMutation(CREATE_ROW);

  if (loading) {
    return <Spinner />;
  }

  if (error) {
    return <div>Error! {error.message}</div>;
  }
  const page = data.getCause.HomePage;
  return (
    <Content top>
      <SEO title={page.Title} keywords={[`gatsby`, `application`, `react`]} />
      <Editor>
        {(page.Layout.LayoutRows || []).sort(sortByWeight).map(row => (
          <RowEditor
            pageId={page.ID}
            layoutId={page.Layout.ID}
            causeId={causeID}
            row={row}
            key={row.ID}
          >
            <Row key={row.ID} container={row.Container}>
              {(row.LayoutColumns || []).sort(sortByWeight).map(column => (
                <ColumnEditor
                  key={column.ID}
                  size={column.Width}
                  rowId={row.ID}
                  pageId={page.ID}
                  causeId={causeID}
                  column={column}
                >
                  <Column key={column.ID} size={12}>
                    {(column.Components || [])
                      .sort(sortByWeight)
                      .map(component => (
                        <ComponentEditor
                          key={component.ID}
                          id={component.ID}
                          causeId={causeID}
                          columnId={column.ID}
                          pageId={page.ID}
                          component={component}
                          componentType={component.ComponentType}
                        >
                          <Node
                            id={page.ID}
                            fields={component.Fields}
                            path={component.ComponentImplementation.Path}
                            key={component.ID}
                          />
                        </ComponentEditor>
                      ))}
                  </Column>
                </ColumnEditor>
              ))}
            </Row>
          </RowEditor>
        ))}
        <Button
          block
          label="Add Row"
          onClick={() =>
            createLayoutRow({
              variables: {
                layoutRow: {
                  Container: false,
                  Layout: page.Layout.ID,
                  Weight: 0,
                },
                buildStatic: true,
              },
              refetchQueries: [
                {
                  query: GET_PAGE,
                  variables: { id: page.ID },
                },
              ],
            })
          }
        />
      </Editor>
    </Content>
  );
}

PageEditor.propTypes = {};

export default PageEditor;
