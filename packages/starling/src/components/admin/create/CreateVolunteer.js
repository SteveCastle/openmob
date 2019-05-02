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

const CREATE_VOLUNTEER = gql`
  mutation createVolunteer($volunteer: VolunteerInput) {
    createVolunteer(volunteer: $volunteer, buildStatic: true) {
      ID
    }
  }
`

const CreateVolunteer = ({ id }) => {
  const createVolunteer = useMutation(CREATE_VOLUNTEER)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createVolunteer({
          variables: {
            volunteer: {
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
                <h1>Create Volunteer</h1>
                <Widget>
                  <Label>VolunteerOpportunity</Label>
                  <Input
                    value={values.VolunteerOpportunity}
                    type="text"
                    name="VolunteerOpportunity"
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

CreateVolunteer.propTypes = {
  id: PropTypes.string,
}

export default CreateVolunteer
