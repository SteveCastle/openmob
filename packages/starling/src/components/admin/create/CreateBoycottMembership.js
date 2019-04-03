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

const CREATE_BOYCOTTMEMBERSHIP = gql`
  mutation createBoycottMembership($boycottMembership: BoycottMembershipInput) {
    createBoycottMembership(
      boycottMembership: $boycottMembership
      buildStatic: true
    ) {
      ID
    }
  }
`

const CreateBoycottMembership = ({ id }) => {
  const createBoycottMembership = useMutation(CREATE_BOYCOTTMEMBERSHIP)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createBoycottMembership({
          variables: {
            boycottMembership: {
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
                <h1>Create BoycottMembership</h1>
                <Widget>
                  <Label>Cause</Label>
                  <Input
                    value={values.Cause}
                    name="Cause"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Boycott</Label>
                  <Input
                    value={values.Boycott}
                    name="Boycott"
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

CreateBoycottMembership.propTypes = {
  id: PropTypes.string,
}

export default CreateBoycottMembership