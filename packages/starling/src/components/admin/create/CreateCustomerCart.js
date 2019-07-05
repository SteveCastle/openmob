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

const CREATE_CUSTOMERCART = gql`
  mutation createCustomerCart($customerCart: CustomerCartInput) {
    createCustomerCart(customerCart: $customerCart, buildStatic: true) {
      ID
    }
  }
`;

const CreateCustomerCart = ({ id }) => {
  const createCustomerCart = useMutation(CREATE_CUSTOMERCART);

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createCustomerCart({
          variables: {
            customerCart: {
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
                <h1>Create CustomerCart</h1>

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

CreateCustomerCart.propTypes = {
  id: PropTypes.string,
};

export default CreateCustomerCart;
