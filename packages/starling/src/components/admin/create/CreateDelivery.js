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

const CREATE_DELIVERY = gql`
  mutation createDelivery($delivery: DeliveryInput) {
    createDelivery(delivery: $delivery, buildStatic: true) {
      ID
    }
  }
`;

const CreateDelivery = ({ id }) => {
  const createDelivery = useMutation(CREATE_DELIVERY);

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createDelivery({
          variables: {
            delivery: {
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
                <h1>Create Delivery</h1>

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

CreateDelivery.propTypes = {
  id: PropTypes.string,
};

export default CreateDelivery;
