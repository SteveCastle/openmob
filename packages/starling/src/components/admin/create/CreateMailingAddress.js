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

const CREATE_MAILINGADDRESS = gql`
  mutation createMailingAddress($mailingAddress: MailingAddressInput) {
    createMailingAddress(mailingAddress: $mailingAddress, buildStatic: true) {
      ID
    }
  }
`

const CreateMailingAddress = ({ id }) => {
  const createMailingAddress = useMutation(CREATE_MAILINGADDRESS)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createMailingAddress({
          variables: {
            mailingAddress: {
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
                <h1>Create MailingAddress</h1>
                <Widget>
                  <Label>StreetAddress</Label>
                  <Input
                    value={values.StreetAddress}
                    type="text"
                    name="StreetAddress"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>City</Label>
                  <Input
                    value={values.City}
                    type="text"
                    name="City"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>State</Label>
                  <Input
                    value={values.State}
                    type="text"
                    name="State"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>ZipCode</Label>
                  <Input
                    value={values.ZipCode}
                    type="text"
                    name="ZipCode"
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

CreateMailingAddress.propTypes = {
  id: PropTypes.string,
}

export default CreateMailingAddress
