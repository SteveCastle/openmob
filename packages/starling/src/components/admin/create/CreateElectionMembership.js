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

const CREATE_ELECTIONMEMBERSHIP = gql`
  mutation createElectionMembership(
    $electionMembership: ElectionMembershipInput
  ) {
    createElectionMembership(
      electionMembership: $electionMembership
      buildStatic: true
    ) {
      ID
    }
  }
`

const CreateElectionMembership = ({ id }) => {
  const createElectionMembership = useMutation(CREATE_ELECTIONMEMBERSHIP)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createElectionMembership({
          variables: {
            electionMembership: {
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
                <h1>Create ElectionMembership</h1>
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
                  <Label>Election</Label>
                  <Input
                    value={values.Election}
                    type="text"
                    name="Election"
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

CreateElectionMembership.propTypes = {
  id: PropTypes.string,
}

export default CreateElectionMembership
