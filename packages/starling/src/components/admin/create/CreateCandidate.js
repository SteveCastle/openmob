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

const CREATE_CANDIDATE = gql`
  mutation createCandidate($candidate: CandidateInput) {
    createCandidate(candidate: $candidate, buildStatic: true) {
      ID
    }
  }
`;

const CreateCandidate = ({ id }) => {
  const createCandidate = useMutation(CREATE_CANDIDATE);

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createCandidate({
          variables: {
            candidate: {
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
                <h1>Create Candidate</h1>
                <Widget>
                  <Label>Election</Label>
                  <Input
                    value={values.Election}
                    type="text"
                    name="Election"
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

CreateCandidate.propTypes = {
  id: PropTypes.string,
};

export default CreateCandidate;
