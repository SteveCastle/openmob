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

const CREATE_PHOTO = gql`
  mutation createPhoto($photo: PhotoInput) {
    createPhoto(photo: $photo, buildStatic: true) {
      ID
    }
  }
`;

const CreatePhoto = ({ id }) => {
  const createPhoto = useMutation(CREATE_PHOTO);

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createPhoto({
          variables: {
            photo: {
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
                <h1>Create Photo</h1>
                <Widget>
                  <Label>URI</Label>
                  <Input
                    value={values.URI}
                    type="text"
                    name="URI"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Width</Label>
                  <Input
                    value={values.Width}
                    type="number"
                    name="Width"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Height</Label>
                  <Input
                    value={values.Height}
                    type="number"
                    name="Height"
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

CreatePhoto.propTypes = {
  id: PropTypes.string,
};

export default CreatePhoto;
