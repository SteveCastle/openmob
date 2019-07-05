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

const CREATE_PRODUCTMEMBERSHIP = gql`
  mutation createProductMembership($productMembership: ProductMembershipInput) {
    createProductMembership(
      productMembership: $productMembership
      buildStatic: true
    ) {
      ID
    }
  }
`;

const CreateProductMembership = ({ id }) => {
  const createProductMembership = useMutation(CREATE_PRODUCTMEMBERSHIP);

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createProductMembership({
          variables: {
            productMembership: {
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
                <h1>Create ProductMembership</h1>
                <Widget>
                  <Label>Cause</Label>
                  <Input
                    value={values.Cause}
                    type="text"
                    name="Cause"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Product</Label>
                  <Input
                    value={values.Product}
                    type="text"
                    name="Product"
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

CreateProductMembership.propTypes = {
  id: PropTypes.string,
};

export default CreateProductMembership;
