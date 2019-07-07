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

const CREATE_FIELD = gql`
  mutation createField($field: FieldInput) {
    createField(field: $field, buildStatic: true) {
      ID
    }
  }
`;

const CreateField = ({ id }) => {
  const createField = useMutation(CREATE_FIELD);

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createField({
          variables: {
            field: {
              ...values,
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
                <h1>Create Field</h1>
                <Widget>
                  <Label>FieldType</Label>
                  <Input
                    value={values.FieldType}
                    type="text"
                    name="FieldType"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>StringValue</Label>
                  <TextArea
                    value={values.StringValue}
                    name="StringValue"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>IntValue</Label>
                  <Input
                    value={values.IntValue}
                    type="number"
                    name="IntValue"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>FloatValue</Label>
                  <Input
                    value={values.FloatValue}
                    type="number"
                    name="FloatValue"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>BooleanValue</Label>
                  <Input
                    value={values.BooleanValue}
                    type="checkbox"
                    name="BooleanValue"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>DateTimeValue</Label>
                  <Input
                    value={values.DateTimeValue}
                    type="text"
                    name="DateTimeValue"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>DataPath</Label>
                  <TextArea
                    value={values.DataPath}
                    name="DataPath"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Component</Label>
                  <Input
                    value={values.Component}
                    type="text"
                    name="Component"
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

CreateField.propTypes = {
  id: PropTypes.string,
};

export default CreateField;
