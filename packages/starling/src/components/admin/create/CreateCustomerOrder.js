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

const CREATE_CUSTOMERORDER = gql`
  mutation createCustomerOrder($customerOrder: CustomerOrderInput) {
    createCustomerOrder(customerOrder: $customerOrder, buildStatic: true) {
      ID
    }
  }
`

const CreateCustomerOrder = ({ id }) => {
  const createCustomerOrder = useMutation(CREATE_CUSTOMERORDER)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createCustomerOrder({
          variables: {
            customerOrder: {
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
                <h1>Create CustomerOrder</h1>
                <Widget>
                  <Label>CustomerCart</Label>
                  <Input
                    value={values.CustomerCart}
                    type="text"
                    name="CustomerCart"
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

CreateCustomerOrder.propTypes = {
  id: PropTypes.string,
}

export default CreateCustomerOrder
