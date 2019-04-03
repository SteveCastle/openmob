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

const CREATE_EMAILADDRESS = gql`
  mutation createEmailAddress($emailAddress: EmailAddressInput) {
    createEmailAddress(emailAddress: $emailAddress, buildStatic: true) {
      ID
    }
  }
`

const CreateEmailAddress = ({ id }) => {
  const createEmailAddress = useMutation(CREATE_EMAILADDRESS)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createEmailAddress({
          variables: {
            emailAddress: {
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
                <h1>Create EmailAddress</h1>
                <Widget>
                  <Label>Address</Label>
                  <Input
                    value={values.Address}
                    name="Address"
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

CreateEmailAddress.propTypes = {
  id: PropTypes.string,
}

export default CreateEmailAddress
