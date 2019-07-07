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

const CREATE_DONATIONCAMPAIGNMEMBERSHIP = gql`
  mutation createDonationCampaignMembership(
    $donationCampaignMembership: DonationCampaignMembershipInput
  ) {
    createDonationCampaignMembership(
      donationCampaignMembership: $donationCampaignMembership
      buildStatic: true
    ) {
      ID
    }
  }
`;

const CreateDonationCampaignMembership = ({ id }) => {
  const createDonationCampaignMembership = useMutation(
    CREATE_DONATIONCAMPAIGNMEMBERSHIP
  );

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createDonationCampaignMembership({
          variables: {
            donationCampaignMembership: {
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
                <h1>Create DonationCampaignMembership</h1>
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
                  <Label>DonationCampaign</Label>
                  <Input
                    value={values.DonationCampaign}
                    type="text"
                    name="DonationCampaign"
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

CreateDonationCampaignMembership.propTypes = {
  id: PropTypes.string,
};

export default CreateDonationCampaignMembership;
