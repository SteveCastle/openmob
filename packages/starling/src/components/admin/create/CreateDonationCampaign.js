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

const CREATE_DONATIONCAMPAIGN = gql`
  mutation createDonationCampaign($donationCampaign: DonationCampaignInput) {
    createDonationCampaign(
      donationCampaign: $donationCampaign
      buildStatic: true
    ) {
      ID
    }
  }
`;

const CreateDonationCampaign = ({ id }) => {
  const createDonationCampaign = useMutation(CREATE_DONATIONCAMPAIGN);

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createDonationCampaign({
          variables: {
            donationCampaign: {
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
                <h1>Create DonationCampaign</h1>
                <Widget>
                  <Label>Title</Label>
                  <Input
                    value={values.Title}
                    type="text"
                    name="Title"
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

CreateDonationCampaign.propTypes = {
  id: PropTypes.string,
};

export default CreateDonationCampaign;
