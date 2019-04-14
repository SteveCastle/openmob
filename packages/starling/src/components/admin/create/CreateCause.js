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

const CREATE_CAUSE = gql`
  mutation createCause($cause: CauseInput) {
    createCause(cause: $cause, buildStatic: true) {
      ID
    }
  }
`

const CreateCause = ({ id }) => {
  const createCause = useMutation(CREATE_CAUSE)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createCause({
          variables: {
            cause: {
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
                <h1>Create Cause</h1>
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
                  <Label>Slug</Label>
                  <Input
                    value={values.Slug}
                    type="text"
                    name="Slug"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Summary</Label>
                  <Input
                    value={values.Summary}
                    type="text"
                    name="Summary"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>HomePage</Label>
                  <Input
                    value={values.HomePage}
                    type="text"
                    name="HomePage"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Photo</Label>
                  <Input
                    value={values.Photo}
                    type="text"
                    name="Photo"
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

CreateCause.propTypes = {
  id: PropTypes.string,
}

export default CreateCause
