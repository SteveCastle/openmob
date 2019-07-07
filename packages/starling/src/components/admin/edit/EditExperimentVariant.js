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

const GET_EXPERIMENTVARIANT = gql`
  query getExperimentVariantById($id: ID!) {
    getExperimentVariant(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      Title
      VariantType
      Experiment {
        ID
      }
      LandingPage {
        ID
      }
      Field {
        ID
      }
      Component {
        ID
      }
    }
  }
`;
const UPDATE_EXPERIMENTVARIANT = gql`
  mutation updateExperimentVariant(
    $id: ID!
    $experimentVariant: ExperimentVariantInput
  ) {
    updateExperimentVariant(
      ID: $id
      experimentVariant: $experimentVariant
      buildStatic: true
    )
  }
`;

function EditExperimentVariant({ id }) {
  const {
    data: { getExperimentVariant: item = {} },
    error,
    loading,
  } = useQuery(GET_EXPERIMENTVARIANT, {
    variables: { id },
  });

  const updateExperimentVariant = useMutation(UPDATE_EXPERIMENTVARIANT);

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
        Title: parseObject(item.Title),
        VariantType: parseObject(item.VariantType),
        Experiment: parseObject(item.Experiment),
        LandingPage: parseObject(item.LandingPage),
        Field: parseObject(item.Field),
        Component: parseObject(item.Component),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateExperimentVariant({
          variables: {
            id: item.ID,
            experimentVariant: {
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
                  <Label>Title</Label>
                  <Input
                    value={values.Title}
                    name="Title"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>VariantType</Label>
                  <Input
                    value={values.VariantType}
                    name="VariantType"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Experiment</Label>
                  <Input
                    value={values.Experiment}
                    name="Experiment"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>LandingPage</Label>
                  <Input
                    value={values.LandingPage}
                    name="LandingPage"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Field</Label>
                  <Input
                    value={values.Field}
                    name="Field"
                    type="text"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Component</Label>
                  <Input
                    value={values.Component}
                    name="Component"
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

EditExperimentVariant.propTypes = {
  id: PropTypes.string,
};

export default EditExperimentVariant;
