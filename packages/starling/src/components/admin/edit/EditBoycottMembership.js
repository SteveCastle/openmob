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

const GET_BOYCOTTMEMBERSHIP = gql`
  query getBoycottMembershipById($id: ID!) {
    getBoycottMembership(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Cause {
        ID
      }
      Boycott {
        ID
      }
    }
  }
`

function EditBoycottMembership({ id }) {
  const {
    data: { getBoycottMembership: item = {} },
    error,
    loading,
  } = useQuery(GET_BOYCOTTMEMBERSHIP, {
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
        <label>Cause</label>
        <Input placeholder={parseObject(item.Cause)} />
      </div>
      <div>
        <label>Boycott</label>
        <Input placeholder={parseObject(item.Boycott)} />
      </div>

      <Button label="Edit" />
    </div>
  )
}

EditBoycottMembership.propTypes = {
  id: PropTypes.string,
}

export default EditBoycottMembership
