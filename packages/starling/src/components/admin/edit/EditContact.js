import React from 'react'
import { useQuery, useMutation } from 'react-apollo-hooks'
import gql from 'graphql-tag'
import { Formik } from 'formik'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const MILLISECONDS = 1000
const isObject = a => !!a && a.constructor === Object
const getValue = obj =>
  Object.entries(obj).reduce((acc, entry) => {
    if (entry[0] === 'seconds') {
      return new Date(entry[1] * MILLISECONDS).toString()
    }
    if (entry[0] === 'ID') {
      return entry[1]
    }
    return acc
  }, '')
const parseObject = obj => (isObject(obj) ? getValue(obj) : obj)

const GET_CONTACT = gql`
  query getContactById($id: ID!) {
    getContact(ID: $id) {
      ID
      CreatedAt {
        seconds
        nanos
      }
      UpdatedAt {
        seconds
        nanos
      }
      FirstName
      MiddleName
      LastName
      Email
      PhoneNumber
    }
  }
`
const UPDATE_CONTACT = gql`
  mutation updateContact($id: ID!, $contact: ContactInput) {
    updateContact(ID: $id, contact: $contact, buildStatic: true)
  }
`

function EditContact({ id }) {
  const {
    data: { getContact: item = {} },
    error,
    loading,
  } = useQuery(GET_CONTACT, {
    variables: { id },
  })

  const updateContact = useMutation(UPDATE_CONTACT)

  if (loading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <div>Error! {error.message}</div>
  }

  return (
    <Formik
      initialValues={{
        ID: parseObject(item.ID),
        CreatedAt: parseObject(item.CreatedAt),
        UpdatedAt: parseObject(item.UpdatedAt),
        FirstName: parseObject(item.FirstName),
        MiddleName: parseObject(item.MiddleName),
        LastName: parseObject(item.LastName),
        Email: parseObject(item.Email),
        PhoneNumber: parseObject(item.PhoneNumber),
      }}
      onSubmit={(values, { setSubmitting }) =>
        updateContact({
          variables: {
            id: item.ID,
            contact: {
              ...values,
              ID: undefined,
              CreatedAt: undefined,
              UpdatedAt: undefined,
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
                <h1>Edit {item.ID}</h1>
                <Widget>
                  <Label>ID</Label>
                  <Input
                    value={values.ID}
                    disabled
                    name="ID"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>CreatedAt</Label>
                  <Input
                    value={values.CreatedAt}
                    disabled
                    name="CreatedAt"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>UpdatedAt</Label>
                  <Input
                    value={values.UpdatedAt}
                    disabled
                    name="UpdatedAt"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>FirstName</Label>
                  <Input
                    value={values.FirstName}
                    name="FirstName"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>MiddleName</Label>
                  <Input
                    value={values.MiddleName}
                    name="MiddleName"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>LastName</Label>
                  <Input
                    value={values.LastName}
                    name="LastName"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>Email</Label>
                  <Input
                    value={values.Email}
                    name="Email"
                    onChange={handleChange}
                    onBlur={handleBlur}
                  />
                </Widget>
                <Widget>
                  <Label>PhoneNumber</Label>
                  <Input
                    value={values.PhoneNumber}
                    name="PhoneNumber"
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

EditContact.propTypes = {
  id: PropTypes.string,
}

export default EditContact
