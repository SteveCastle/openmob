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

const CREATE_VOLUNTEEROPPORTUNITY = gql`
  mutation createVolunteerOpportunity(
    $volunteerOpportunity: VolunteerOpportunityInput
  ) {
    createVolunteerOpportunity(
      volunteerOpportunity: $volunteerOpportunity
      buildStatic: true
    ) {
      ID
    }
  }
`

const CreateVolunteerOpportunity = ({ id }) => {
  const createVolunteerOpportunity = useMutation(CREATE_VOLUNTEEROPPORTUNITY)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createVolunteerOpportunity({
          variables: {
            volunteerOpportunity: {
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
                <h1>Create VolunteerOpportunity</h1>
                <Widget>
                  <Label>Title</Label>
                  <Input
                    value={values.Title}
                    type="text"
                    name="Title"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>VolunteerOpportunityType</Label>
                  <Input
                    value={values.VolunteerOpportunityType}
                    type="text"
                    name="VolunteerOpportunityType"
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

CreateVolunteerOpportunity.propTypes = {
  id: PropTypes.string,
}

export default CreateVolunteerOpportunity
