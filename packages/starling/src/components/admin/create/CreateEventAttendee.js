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

const CREATE_EVENTATTENDEE = gql`
  mutation createEventAttendee($eventAttendee: EventAttendeeInput) {
    createEventAttendee(eventAttendee: $eventAttendee, buildStatic: true) {
      ID
    }
  }
`

const CreateEventAttendee = ({ id }) => {
  const createEventAttendee = useMutation(CREATE_EVENTATTENDEE)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createEventAttendee({
          variables: {
            eventAttendee: {
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
                <h1>Create EventAttendee</h1>
                <Widget>
                  <Label>LiveEvent</Label>
                  <Input
                    value={values.LiveEvent}
                    name="LiveEvent"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Contact</Label>
                  <Input
                    value={values.Contact}
                    name="Contact"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Cause</Label>
                  <Input
                    value={values.Cause}
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

CreateEventAttendee.propTypes = {
  id: PropTypes.string,
}

export default CreateEventAttendee