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

const CREATE_HOMEPAGE = gql`
  mutation createHomePage($homePage: HomePageInput) {
    createHomePage(homePage: $homePage, buildStatic: true) {
      ID
    }
  }
`

const CreateHomePage = ({ id }) => {
  const createHomePage = useMutation(CREATE_HOMEPAGE)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createHomePage({
          variables: {
            homePage: {
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
                <h1>Create HomePage</h1>
                <Widget>
                  <Label>Title</Label>
                  <Input
                    value={values.Title}
                    name="Title"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Layout</Label>
                  <Input
                    value={values.Layout}
                    name="Layout"
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

CreateHomePage.propTypes = {
  id: PropTypes.string,
}

export default CreateHomePage
