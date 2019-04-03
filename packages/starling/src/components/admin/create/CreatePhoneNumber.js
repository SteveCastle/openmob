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

const CREATE_PHONENUMBER = gql`
  mutation createPhoneNumber($phoneNumber: PhoneNumberInput) {
    createPhoneNumber(phoneNumber: $phoneNumber, buildStatic: true) {
      ID
    }
  }
`

const CreatePhoneNumber = ({ id }) => {
  const createPhoneNumber = useMutation(CREATE_PHONENUMBER)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createPhoneNumber({
          variables: {
            phoneNumber: {
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
                <h1>Create PhoneNumber</h1>
                <Widget>
                  <Label>PhoneNumber</Label>
                  <Input
                    value={values.PhoneNumber}
                    name="PhoneNumber"
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

CreatePhoneNumber.propTypes = {
  id: PropTypes.string,
}

export default CreatePhoneNumber
