import React from 'react';
import PropTypes from 'prop-types';
import { useQuery, useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import Button from '@openmob/bluebird/src/components/buttons/Button';

import Row from '@openmob/bluebird/src/components/layout/Row';
import Column from '@openmob/bluebird/src/components/layout/Column';
import Content from '@openmob/bluebird/src/components/layout/Content';
import Node from '../Node';
import Editor from './page-editor/PageEditor';
import RowEditor from './page-editor/RowEditor';
import ColumnEditor from './page-editor/ColumnEditor';
import ComponentEditor from './page-editor/ComponentEditor';
import SEO from '../SEO';

const sortByWeight = (a, b) => a.Weight - b.Weight;
const GET_PAGE = gql`
  query HomePageQuery($id: ID!) {
    getHomePage(ID: $id) {
      ID
      Title
      Layout {
        ID
        LayoutRows {
          ID
          Container
          Weight
          LayoutColumns {
            ID
            Width
            Weight
            Components {
              ID
              Fields {
                ID
                FieldType {
                  Title
                  DataType
                  PropName
                }
                StringValue
                IntValue
                FloatValue
                BooleanValue
                DateTimeValue {
                  seconds
                }
                DataPathValue
                DataPath
              }
              ComponentImplementation {
                Path
              }
            }
          }
        }
        LayoutType {
          ID
        }
      }
    }
  }
`;

const CREATE_ROW = gql`
  mutation createLayoutRow($layoutRow: LayoutRowInput) {
    createLayoutRow(layoutRow: $layoutRow, buildStatic: true) {
      ID
    }
  }
`;

const CREATE_COLUMN = gql`
  mutation createLayoutColumn($layoutColumn: LayoutColumnInput) {
    createColumn(layoutColumn: $cause, buildStatic: true)
  }
`;

const CREATE_COMPONENT = gql`
  mutation createComponent($component: ComponentInput) {
    createComponent(component: $component, buildStatic: true)
  }
`;

function PageEditor({ navigate = () => {}, pageID }) {
  const {
    data: { getHomePage: page = {} },
    error,
    loading,
  } = useQuery(GET_PAGE, {
    variables: {
      id: pageID,
    },
  });

  const createLayoutRow = useMutation(CREATE_ROW);

  console.log('page', page);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error! {error.message}</div>;
  }

  return (
    <Content top>
      <SEO title={page.Title} keywords={[`gatsby`, `application`, `react`]} />
      <Editor>
        {(page.Layout.LayoutRows || []).sort(sortByWeight).map(row => (
          <RowEditor>
            <Row key={row.ID} container={row.Container}>
              {(row.LayoutColumns || []).sort(sortByWeight).map(column => (
                <ColumnEditor size={column.Width}>
                  <Column key={column.ID} size={12}>
                    {(column.Components || [])
                      .sort(sortByWeight)
                      .map(component => (
                        <ComponentEditor>
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
            })
          }
        />
      </Editor>
    </Content>
  );
}

PageEditor.propTypes = {
  navigate: PropTypes.func,
};

export default PageEditor;
