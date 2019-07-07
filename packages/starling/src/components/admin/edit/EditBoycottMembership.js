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

const GET_BOYCOTTMEMBERSHIP = gql`
  query getBoycottMembershipById($id: ID!) {
    getBoycottMembership(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      Cause {
        ID
      }
      Boycott {
        ID
      }
    }
  }
`;
const UPDATE_BOYCOTTMEMBERSHIP = gql`
  mutation updateBoycottMembership(
    $id: ID!
    $boycottMembership: BoycottMembershipInput
  ) {
    updateBoycottMembership(
      ID: $id
      boycottMembership: $boycottMembership
      buildStatic: true
    )
  }
`;

function EditBoycottMembership({ id }) {
  const {
    data: { getBoycottMembership: item = {} },
    error,
    loading,
  } = useQuery(GET_BOYCOTTMEMBERSHIP, {
    variables: { id },
  });

  const updateBoycottMembership = useMutation(UPDATE_BOYCOTTMEMBERSHIP);

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
        Cause: parseObject(item.Cause),
        Boycott: parseObject(item.Boycott),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateBoycottMembership({
          variables: {
            id: item.ID,
            boycottMembership: {
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
                  <Label>Cause</Label>
                  <Input
                    value={values.Cause}
                    name="Cause"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Boycott</Label>
                  <Input
                    value={values.Boycott}
                    name="Boycott"
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

EditBoycottMembership.propTypes = {
  id: PropTypes.string,
};

export default EditBoycottMembership;
