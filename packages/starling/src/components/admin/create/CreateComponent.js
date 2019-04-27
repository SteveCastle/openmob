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

const CREATE_COMPONENT = gql`
  mutation createComponent($component: ComponentInput) {
    createComponent(component: $component, buildStatic: true) {
      ID
    }
  }
`

const CreateComponent = ({ id }) => {
  const createComponent = useMutation(CREATE_COMPONENT)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createComponent({
          variables: {
            component: {
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
                <h1>Create Component</h1>
                <Widget>
                  <Label>ComponentType</Label>
                  <Input
                    value={values.ComponentType}
                    type="text"
                    name="ComponentType"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>ComponentImplementation</Label>
                  <Input
                    value={values.ComponentImplementation}
                    type="text"
                    name="ComponentImplementation"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>LayoutColumn</Label>
                  <Input
                    value={values.LayoutColumn}
                    type="text"
                    name="LayoutColumn"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Weight</Label>
                  <Input
                    value={values.Weight}
                    type="number"
                    name="Weight"
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

CreateComponent.propTypes = {
  id: PropTypes.string,
}

export default CreateComponent
