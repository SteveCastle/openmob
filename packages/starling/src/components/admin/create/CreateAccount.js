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

const CREATE_ACCOUNT = gql`
  mutation createAccount($account: AccountInput) {
    createAccount(account: $account, buildStatic: true) {
      ID
    }
  }
`

const CreateAccount = ({ id }) => {
  const createAccount = useMutation(CREATE_ACCOUNT)

  return (
    <Formik
      onSubmit={(values, { setSubmitting }) =>
        createAccount({
          variables: {
            account: {
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
                <h1>Create Account</h1>
                <Widget>
                  <Label>Username</Label>
                  <Input
                    value={values.Username}
                    type="text"
                    name="Username"
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

CreateAccount.propTypes = {
  id: PropTypes.string,
}

export default CreateAccount
