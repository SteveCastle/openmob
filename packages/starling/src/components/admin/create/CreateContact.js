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

const CREATE_CONTACT = gql`
  mutation createContact($contact: ContactInput) {
    createContact(contact: $contact, buildStatic: true) {
      ID
    }
  }
`

const CreateContact = ({ id }) => {
  const createContact = useMutation(CREATE_CONTACT)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createContact({
          variables: {
            contact: {
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
                <h1>Create Contact</h1>
                <Widget>
                  <Label>FirstName</Label>
                  <Input
                    value={values.FirstName}
                    type="text"
                    name="FirstName"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>MiddleName</Label>
                  <Input
                    value={values.MiddleName}
                    type="text"
                    name="MiddleName"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>LastName</Label>
                  <Input
                    value={values.LastName}
                    type="text"
                    name="LastName"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Email</Label>
                  <Input
                    value={values.Email}
                    type="text"
                    name="Email"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>PhoneNumber</Label>
                  <Input
                    value={values.PhoneNumber}
                    type="text"
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

CreateContact.propTypes = {
  id: PropTypes.string,
}

export default CreateContact
