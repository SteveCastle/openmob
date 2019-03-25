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

const GET_FIELD = gql`
  query getFieldById($id: ID!) {
    getField(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      FieldType {
        ID
      }
      StringValue
      IntValue
      FloatValue
      BooleanValue
      DateTimeValue {
        seconds
      }
      Component {
        ID
      }
    }
  }
`

function EditField({ id }) {
  const {
    data: { getField: item = {} },
    error,
    loading,
  } = useQuery(GET_FIELD, {
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
        <label>FieldType</label>
        <Input placeholder={parseObject(item.FieldType)} />
      </div>
      <div>
        <label>StringValue</label>
        <Input placeholder={parseObject(item.StringValue)} />
      </div>
      <div>
        <label>IntValue</label>
        <Input placeholder={parseObject(item.IntValue)} />
      </div>
      <div>
        <label>FloatValue</label>
        <Input placeholder={parseObject(item.FloatValue)} />
      </div>
      <div>
        <label>BooleanValue</label>
        <Input placeholder={parseObject(item.BooleanValue)} />
      </div>
      <div>
        <label>DateTimeValue</label>
        <Input placeholder={parseObject(item.DateTimeValue)} />
      </div>
      <div>
        <label>Component</label>
        <Input placeholder={parseObject(item.Component)} />
      </div>

      <Button label="Edit" />
    </div>
  )
}

EditField.propTypes = {
  id: PropTypes.string,
}

export default EditField
