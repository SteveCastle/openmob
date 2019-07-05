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

const CREATE_VOLUNTEEROPPORTUNITYMEMBERSHIP = gql`
  mutation createVolunteerOpportunityMembership(
    $volunteerOpportunityMembership: VolunteerOpportunityMembershipInput
  ) {
    createVolunteerOpportunityMembership(
      volunteerOpportunityMembership: $volunteerOpportunityMembership
      buildStatic: true
    ) {
      ID
    }
  }
`;

const CreateVolunteerOpportunityMembership = ({ id }) => {
  const createVolunteerOpportunityMembership = useMutation(
    CREATE_VOLUNTEEROPPORTUNITYMEMBERSHIP
  );

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createVolunteerOpportunityMembership({
          variables: {
            volunteerOpportunityMembership: {
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
                <h1>Create VolunteerOpportunityMembership</h1>
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
                  <Label>VolunteerOpportunity</Label>
                  <Input
                    value={values.VolunteerOpportunity}
                    type="text"
                    name="VolunteerOpportunity"
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

CreateVolunteerOpportunityMembership.propTypes = {
  id: PropTypes.string,
};

export default CreateVolunteerOpportunityMembership;
