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

const CREATE_POLLMEMBERSHIP = gql`
  mutation createPollMembership($pollMembership: PollMembershipInput) {
    createPollMembership(pollMembership: $pollMembership, buildStatic: true) {
      ID
    }
  }
`;

const CreatePollMembership = ({ id }) => {
  const createPollMembership = useMutation(CREATE_POLLMEMBERSHIP);

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createPollMembership({
          variables: {
            pollMembership: {
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
                <h1>Create PollMembership</h1>
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
                  <Label>Petition</Label>
                  <Input
                    value={values.Petition}
                    type="text"
                    name="Petition"
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

CreatePollMembership.propTypes = {
  id: PropTypes.string,
};

export default CreatePollMembership;
