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

const CREATE_AGENT = gql`
  mutation createAgent($agent: AgentInput) {
    createAgent(agent: $agent, buildStatic: true) {
      ID
    }
  }
`

const CreateAgent = ({ id }) => {
  const createAgent = useMutation(CREATE_AGENT)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createAgent({
          variables: {
            agent: {
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
                <h1>Create Agent</h1>
                <Widget>
                  <Label>Account</Label>
                  <Input
                    value={values.Account}
                    type="text"
                    name="Account"
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

CreateAgent.propTypes = {
  id: PropTypes.string,
}

export default CreateAgent
