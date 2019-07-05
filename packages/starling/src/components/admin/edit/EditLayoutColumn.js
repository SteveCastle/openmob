/* eslint-disable */

import React from 'react';
import { useQuery, useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import { Formik } from 'formik';
import PropTypes from 'prop-types';
import Content from '@openmob/bluebird/src/components/layout/Content';
import Card from '@openmob/bluebird/src/components/cards/Card';
import Form from '@openmob/bluebird/src/components/forms/Form';
import Widget from '@openmob/bluebird/src/components/forms/Widget';
import Label from '@openmob/bluebird/src/components/forms/Label';
import Input from '@openmob/bluebird/src/components/forms/Input';
import TextArea from '@openmob/bluebird/src/components/forms/TextArea';
import Button from '@openmob/bluebird/src/components/buttons/Button';
import parseObject from '../../../common/helpers';

const GET_LAYOUTCOLUMN = gql`
  query getLayoutColumnById($id: ID!) {
    getLayoutColumn(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      LayoutRow {
        ID
      }
      Width
      Weight
    }
  }
`;
const UPDATE_LAYOUTCOLUMN = gql`
  mutation updateLayoutColumn($id: ID!, $layoutColumn: LayoutColumnInput) {
    updateLayoutColumn(ID: $id, layoutColumn: $layoutColumn, buildStatic: true)
  }
`;

function EditLayoutColumn({ id }) {
  const {
    data: { getLayoutColumn: item = {} },
    error,
    loading,
  } = useQuery(GET_LAYOUTCOLUMN, {
    variables: { id },
  });

  const updateLayoutColumn = useMutation(UPDATE_LAYOUTCOLUMN);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error! {error.message}</div>;
  }

  return (
    <Formik
      initialValues={{
        ID: parseObject(item.ID),
        CreatedAt: parseObject(item.CreatedAt),
        UpdatedAt: parseObject(item.UpdatedAt),
        LayoutRow: parseObject(item.LayoutRow),
        Width: parseObject(item.Width),
        Weight: parseObject(item.Weight),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateLayoutColumn({
          variables: {
            id: item.ID,
            layoutColumn: {
              ...values,
              ID: undefined,
              CreatedAt: undefined,
              UpdatedAt: undefined,
            },
          },
        })
      }
    >
      {props => {
        const { values, handleChange, handleBlur, handleSubmit } = props;
        return (
          <Content>
            <Card>
              <Form>
                <h1>Edit {item.ID}</h1>
                <Widget>
                  <Label>ID</Label>
                  <Input
                    value={values.ID}
                    disabled
                    name="ID"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>CreatedAt</Label>
                  <Input
                    value={values.CreatedAt}
                    disabled
                    name="CreatedAt"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>UpdatedAt</Label>
                  <Input
                    value={values.UpdatedAt}
                    disabled
                    name="UpdatedAt"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>LayoutRow</Label>
                  <Input
                    value={values.LayoutRow}
                    name="LayoutRow"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Width</Label>
                  <Input
                    value={values.Width}
                    name="Width"
                    type="number"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Weight</Label>
                  <Input
                    value={values.Weight}
                    name="Weight"
                    type="number"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>

                <Button
                  label="Save"
                  block
                  variant="primary"
                  onClick={handleSubmit}
                />
              </Form>
            </Card>
          </Content>
        );
      }}
    </Formik>
  );
}

EditLayoutColumn.propTypes = {
  id: PropTypes.string,
};

export default EditLayoutColumn;
