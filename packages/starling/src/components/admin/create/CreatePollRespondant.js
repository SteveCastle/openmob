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

const CREATE_POLLRESPONDANT = gql`
  mutation createPollRespondant($pollRespondant: PollRespondantInput) {
    createPollRespondant(pollRespondant: $pollRespondant, buildStatic: true) {
      ID
    }
  }
`

const CreatePollRespondant = ({ id }) => {
  const createPollRespondant = useMutation(CREATE_POLLRESPONDANT)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createPollRespondant({
          variables: {
            pollRespondant: {
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
                <h1>Create PollRespondant</h1>
                <Widget>
                  <Label>Poll</Label>
                  <Input
                    value={values.Poll}
                    type="text"
                    name="Poll"
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

CreatePollRespondant.propTypes = {
  id: PropTypes.string,
}

export default CreatePollRespondant
