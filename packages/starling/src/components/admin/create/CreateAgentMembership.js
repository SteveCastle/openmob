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

const CREATE_AGENTMEMBERSHIP = gql`
  mutation createAgentMembership($agentMembership: AgentMembershipInput) {
    createAgentMembership(
      agentMembership: $agentMembership
      buildStatic: true
    ) {
      ID
    }
  }
`

const CreateAgentMembership = ({ id }) => {
  const createAgentMembership = useMutation(CREATE_AGENTMEMBERSHIP)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createAgentMembership({
          variables: {
            agentMembership: {
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
                <h1>Create AgentMembership</h1>
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
                  <Label>Agent</Label>
                  <Input
                    value={values.Agent}
                    name="Agent"
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

CreateAgentMembership.propTypes = {
  id: PropTypes.string,
}

export default CreateAgentMembership
