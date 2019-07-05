/* eslint-disable */

import React from 'react';
import PropTypes from 'prop-types';
import { useMutation } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import { Formik } from 'formik';
import Content from '@openmob/bluebird/src/components/layout/Content';
import Card from '@openmob/bluebird/src/components/cards/Card';
import Form from '@openmob/bluebird/src/components/forms/Form';
import Widget from '@openmob/bluebird/src/components/forms/Widget';
import Label from '@openmob/bluebird/src/components/forms/Label';
import Input from '@openmob/bluebird/src/components/forms/Input';
import TextArea from '@openmob/bluebird/src/components/forms/TextArea';
import Button from '@openmob/bluebird/src/components/buttons/Button';

const CREATE_LAYOUT = gql`
  mutation createLayout($layout: LayoutInput) {
    createLayout(layout: $layout, buildStatic: true) {
      ID
    }
  }
`;

const CreateLayout = ({ id }) => {
  const createLayout = useMutation(CREATE_LAYOUT);

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createLayout({
          variables: {
            layout: {
              ...values,
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
                <h1>Create Layout</h1>
                <Widget>
                  <Label>LayoutType</Label>
                  <Input
                    value={values.LayoutType}
                    type="text"
                    name="LayoutType"
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
};

CreateLayout.propTypes = {
  id: PropTypes.string,
};

export default CreateLayout;
