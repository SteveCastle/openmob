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

const CREATE_COMPONENTTYPE = gql`
  mutation createComponentType($componentType: ComponentTypeInput) {
    createComponentType(componentType: $componentType, buildStatic: true) {
      ID
    }
  }
`

const CreateComponentType = ({ id }) => {
  const createComponentType = useMutation(CREATE_COMPONENTTYPE)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createComponentType({
          variables: {
            componentType: {
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
                <h1>Create ComponentType</h1>
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
        )
      }}
    </Formik>
  )
}

CreateComponentType.propTypes = {
  id: PropTypes.string,
}

export default CreateComponentType
