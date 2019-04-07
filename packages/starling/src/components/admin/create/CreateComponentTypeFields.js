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
import Button from '@openmob/bluebird/src/components/buttons/Button'

const CREATE_COMPONENTTYPEFIELDS = gql`
  mutation createComponentTypeFields(
    $componentTypeFields: ComponentTypeFieldsInput
  ) {
    createComponentTypeFields(
      componentTypeFields: $componentTypeFields
      buildStatic: true
    ) {
      ID
    }
  }
`

const CreateComponentTypeFields = ({ id }) => {
  const createComponentTypeFields = useMutation(CREATE_COMPONENTTYPEFIELDS)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createComponentTypeFields({
          variables: {
            componentTypeFields: {
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
                <h1>Create ComponentTypeFields</h1>
                <Widget>
                  <Label>ComponentType</Label>
                  <Input
                    value={values.ComponentType}
                    name="ComponentType"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>FieldType</Label>
                  <Input
                    value={values.FieldType}
                    name="FieldType"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Weight</Label>
                  <Input
                    value={values.Weight}
                    name="Weight"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Required</Label>
                  <Input
                    value={values.Required}
                    name="Required"
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

CreateComponentTypeFields.propTypes = {
  id: PropTypes.string,
}

export default CreateComponentTypeFields
