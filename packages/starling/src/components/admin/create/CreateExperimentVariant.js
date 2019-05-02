/* eslint-disable */

import React from 'react'
import PropTypes from 'prop-types'
import { useMutation } from 'react-apollo-hooks'
import gql from 'graphql-tag'
import { Formik } from 'formik'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import TextArea from '@openmob/bluebird/src/components/forms/TextArea'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CREATE_EXPERIMENTVARIANT = gql`
  mutation createExperimentVariant($experimentVariant: ExperimentVariantInput) {
    createExperimentVariant(
      experimentVariant: $experimentVariant
      buildStatic: true
    ) {
      ID
    }
  }
`

const CreateExperimentVariant = ({ id }) => {
  const createExperimentVariant = useMutation(CREATE_EXPERIMENTVARIANT)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createExperimentVariant({
          variables: {
            experimentVariant: {
              ...values,
            },
          },
        })
      }
    >
      {props => {
        const { values, handleChange, handleBlur, handleSubmit } = props
        return (
          <Content>
            <Card>
              <Form>
                <h1>Create ExperimentVariant</h1>
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
                  <Label>VariantType</Label>
                  <Input
                    value={values.VariantType}
                    type="text"
                    name="VariantType"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Experiment</Label>
                  <Input
                    value={values.Experiment}
                    type="text"
                    name="Experiment"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>LandingPage</Label>
                  <Input
                    value={values.LandingPage}
                    type="text"
                    name="LandingPage"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Field</Label>
                  <Input
                    value={values.Field}
                    type="text"
                    name="Field"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Component</Label>
                  <Input
                    value={values.Component}
                    type="text"
                    name="Component"
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
        )
      }}
    </Formik>
  )
}

CreateExperimentVariant.propTypes = {
  id: PropTypes.string,
}

export default CreateExperimentVariant
