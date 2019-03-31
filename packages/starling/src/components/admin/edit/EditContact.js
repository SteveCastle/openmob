import React from 'react'
import { useQuery } from 'react-apollo-hooks'
import gql from 'graphql-tag'
import PropTypes from 'prop-types'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Form from '@openmob/bluebird/src/components/forms/Form'
import Widget from '@openmob/bluebird/src/components/forms/Widget'
import Label from '@openmob/bluebird/src/components/forms/Label'
import Input from '@openmob/bluebird/src/components/forms/Input'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const isObject = a => !!a && a.constructor === Object
const getValue = obj =>
  Object.entries(obj).reduce(entry =>
    entry[0] === 'seconds' || entry[0] === 'ID' ? entry[1] : null
  )
const parseObject = obj => (isObject(obj) ? getValue(obj) : obj)

const GET_CONTACT = gql`
  query getContactById($id: ID!) {
    getContact(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      FirstName
      MiddleName
      LastName
      Email
      PhoneNumber
    }
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

  if (loading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <div>Error! {error.message}</div>
  }

  return (
    <Content>
      <Card>
        <Form>
          <h1>Edit {item.ID}</h1>
          <Widget>
            <Label>ID</Label>
            <Input placeholder={parseObject(item.ID)} />
          </Widget>
          <Widget>
            <Label>CreatedAt</Label>
            <Input placeholder={parseObject(item.CreatedAt)} />
          </Widget>
          <Widget>
            <Label>UpdatedAt</Label>
            <Input placeholder={parseObject(item.UpdatedAt)} />
          </Widget>
          <Widget>
            <Label>FirstName</Label>
            <Input placeholder={parseObject(item.FirstName)} />
          </Widget>
          <Widget>
            <Label>MiddleName</Label>
            <Input placeholder={parseObject(item.MiddleName)} />
          </Widget>
          <Widget>
            <Label>LastName</Label>
            <Input placeholder={parseObject(item.LastName)} />
          </Widget>
          <Widget>
            <Label>Email</Label>
            <Input placeholder={parseObject(item.Email)} />
          </Widget>
          <Widget>
            <Label>PhoneNumber</Label>
            <Input placeholder={parseObject(item.PhoneNumber)} />
          </Widget>

          <Button label="Edit" block />
        </Form>
      </Card>
    </Content>
  )
}

EditContact.propTypes = {
  id: PropTypes.string,
}

export default EditContact
