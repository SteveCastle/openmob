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

const GET_CONTACT = gql`
  query getContactById($id: ID!) {
    getContact(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      FirstName
      MiddleName
      LastName
      Email
      PhoneNumber
    }
  }
`;
const UPDATE_CONTACT = gql`
  mutation updateContact($id: ID!, $contact: ContactInput) {
    updateContact(ID: $id, contact: $contact, buildStatic: true)
  }
`;

function EditContact({ id }) {
  const {
    data: { getContact: item = {} },
    error,
    loading,
  } = useQuery(GET_CONTACT, {
    variables: { id },
  });

  const updateContact = useMutation(UPDATE_CONTACT);

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
        FirstName: parseObject(item.FirstName),
        MiddleName: parseObject(item.MiddleName),
        LastName: parseObject(item.LastName),
        Email: parseObject(item.Email),
        PhoneNumber: parseObject(item.PhoneNumber),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateContact({
          variables: {
            id: item.ID,
            contact: {
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
                  <Label>FirstName</Label>
                  <Input
                    value={values.FirstName}
                    name="FirstName"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>MiddleName</Label>
                  <Input
                    value={values.MiddleName}
                    name="MiddleName"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>LastName</Label>
                  <Input
                    value={values.LastName}
                    name="LastName"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Email</Label>
                  <Input
                    value={values.Email}
                    name="Email"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>PhoneNumber</Label>
                  <Input
                    value={values.PhoneNumber}
                    name="PhoneNumber"
                    type="text"
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

EditContact.propTypes = {
  id: PropTypes.string,
};

export default EditContact;
