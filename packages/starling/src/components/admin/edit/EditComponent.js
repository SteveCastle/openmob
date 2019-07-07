/* eslint-disable */

import React from 'react';
import { useQuery, useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import { Formik } from 'formik';
import PropTypes from 'prop-types';
import Spinner from '@openmob/bluebird/src/components/loaders/Spinner';
import Content from '@openmob/bluebird/src/components/layout/Content';
import Card from '@openmob/bluebird/src/components/cards/Card';
import Form from '@openmob/bluebird/src/components/forms/Form';
import Widget from '@openmob/bluebird/src/components/forms/Widget';
import Label from '@openmob/bluebird/src/components/forms/Label';
import Input from '@openmob/bluebird/src/components/forms/Input';
import TextArea from '@openmob/bluebird/src/components/forms/TextArea';
import Button from '@openmob/bluebird/src/components/buttons/Button';
import parseObject from '../../../common/helpers';

const GET_COMPONENT = gql`
  query getComponentById($id: ID!) {
    getComponent(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      ComponentType {
        ID
      }
      ComponentImplementation {
        ID
      }
      LayoutColumn {
        ID
      }
      Weight
    }
  }
`;
const UPDATE_COMPONENT = gql`
  mutation updateComponent($id: ID!, $component: ComponentInput) {
    updateComponent(ID: $id, component: $component, buildStatic: true)
  }
`;

function EditComponent({ id }) {
  const {
    data: { getComponent: item = {} },
    error,
    loading,
  } = useQuery(GET_COMPONENT, {
    variables: { id },
  });

  const updateComponent = useMutation(UPDATE_COMPONENT);

  if (loading) {
    return <Spinner />;
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
        ComponentType: parseObject(item.ComponentType),
        ComponentImplementation: parseObject(item.ComponentImplementation),
        LayoutColumn: parseObject(item.LayoutColumn),
        Weight: parseObject(item.Weight),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateComponent({
          variables: {
            id: item.ID,
            component: {
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
          <Content top>
            <Card width={9 / 10}>
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
                  <Label>ComponentType</Label>
                  <Input
                    value={values.ComponentType}
                    name="ComponentType"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>ComponentImplementation</Label>
                  <Input
                    value={values.ComponentImplementation}
                    name="ComponentImplementation"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>LayoutColumn</Label>
                  <Input
                    value={values.LayoutColumn}
                    name="LayoutColumn"
                    type="text"
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

EditComponent.propTypes = {
  id: PropTypes.string,
};

export default EditComponent;
