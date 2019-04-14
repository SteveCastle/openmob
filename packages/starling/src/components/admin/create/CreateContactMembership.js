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

const CREATE_CONTACTMEMBERSHIP = gql`
  mutation createContactMembership($contactMembership: ContactMembershipInput) {
    createContactMembership(
      contactMembership: $contactMembership
      buildStatic: true
    ) {
      ID
    }
  }
`

const CreateContactMembership = ({ id }) => {
  const createContactMembership = useMutation(CREATE_CONTACTMEMBERSHIP)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createContactMembership({
          variables: {
            contactMembership: {
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
                <h1>Create ContactMembership</h1>
                <Widget>
                  <Label>Cause</Label>
                  <Input
                    value={values.Cause}
                    type="text"
                    name="Cause"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Contact</Label>
                  <Input
                    value={values.Contact}
                    type="text"
                    name="Contact"
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

CreateContactMembership.propTypes = {
  id: PropTypes.string,
}

export default CreateContactMembership
