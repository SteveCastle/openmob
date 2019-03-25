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

const GET_MAILINGADDRESS = gql`
  query getMailingAddressById($id: ID!) {
    getMailingAddress(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      StreetAddress
      City
      State
      ZipCode
    }
  }
`

function EditMailingAddress({ id }) {
  const {
    data: { getMailingAddress: item = {} },
    error,
    loading,
  } = useQuery(GET_MAILINGADDRESS, {
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
        <label>StreetAddress</label>
        <Input placeholder={parseObject(item.StreetAddress)} />
      </div>
      <div>
        <label>City</label>
        <Input placeholder={parseObject(item.City)} />
      </div>
      <div>
        <label>State</label>
        <Input placeholder={parseObject(item.State)} />
      </div>
      <div>
        <label>ZipCode</label>
        <Input placeholder={parseObject(item.ZipCode)} />
      </div>

      <Button label="Edit" />
    </div>
  )
}

EditMailingAddress.propTypes = {
  id: PropTypes.string,
}

export default EditMailingAddress
