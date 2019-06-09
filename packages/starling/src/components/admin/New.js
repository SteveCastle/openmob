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

const CREATE_CAUSE = gql`
  mutation newCause(
    $Title: String!
    $Slug: String!
    $Summary: String!
    $FeaturedImage: String
  ) {
    newCause(
      Title: $Title
      Slug: $Slug
      Summary: $Summary
      FeaturedImage: $FeaturedImage
    )
  }
`;

const New = ({ accountId }) => {
  const newCause = useMutation(CREATE_CAUSE);

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        newCause({
          variables: {
            ...values,
          },
        })
      }
    >
      {props => {
        const { values, handleChange, handleBlur, handleSubmit } = props;
        return (
          <Content>
            <Card width={1 / 2}>
              <Form>
                <h1>Create a New Cause</h1>
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
                <Widget>
                  <Label>Slug</Label>
                  <Input
                    value={values.Slug}
                    type="text"
                    name="Slug"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Featured Image</Label>
                  <Input
                    value={values.FeaturedImage}
                    type="text"
                    name="FeaturedImage"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Summary</Label>
                  <TextArea
                    value={values.Summary}
                    name="Summary"
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

New.propTypes = {
  accountId: PropTypes.string,
};

export default New;
