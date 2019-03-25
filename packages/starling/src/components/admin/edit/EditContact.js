import React from 'react'
import { useQuery } from 'react-apollo-hooks'
import gql from 'graphql-tag'
import PropTypes from 'prop-types'
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
    <div>
      <h1>Edit {item.ID}</h1>
      <div>
        <label>ID</label>
        <Input placeholder={parseObject(item.ID)} />
      </div>
      <div>
        <label>CreatedAt</label>
        <Input placeholder={parseObject(item.CreatedAt)} />
      </div>
      <div>
        <label>UpdatedAt</label>
        <Input placeholder={parseObject(item.UpdatedAt)} />
      </div>
      <div>
        <label>FirstName</label>
        <Input placeholder={parseObject(item.FirstName)} />
      </div>
      <div>
        <label>MiddleName</label>
        <Input placeholder={parseObject(item.MiddleName)} />
      </div>
      <div>
        <label>LastName</label>
        <Input placeholder={parseObject(item.LastName)} />
      </div>
      <div>
        <label>Email</label>
        <Input placeholder={parseObject(item.Email)} />
      </div>
      <div>
        <label>PhoneNumber</label>
        <Input placeholder={parseObject(item.PhoneNumber)} />
      </div>

      <Button label="Edit" />
    </div>
  )
}

EditContact.propTypes = {
  id: PropTypes.string,
}

export default EditContact
