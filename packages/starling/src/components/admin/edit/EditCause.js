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

const GET_CAUSE = gql`
  query getCauseById($id: ID!) {
    getCause(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Title
      Slug
      Summary
      HomePage {
        ID
      }
      Photo {
        ID
      }
    }
  }
`

function EditCause({ id }) {
  const {
    data: { getCause: item = {} },
    error,
    loading,
  } = useQuery(GET_CAUSE, {
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
        <label>Title</label>
        <Input placeholder={parseObject(item.Title)} />
      </div>
      <div>
        <label>Slug</label>
        <Input placeholder={parseObject(item.Slug)} />
      </div>
      <div>
        <label>Summary</label>
        <Input placeholder={parseObject(item.Summary)} />
      </div>
      <div>
        <label>HomePage</label>
        <Input placeholder={parseObject(item.HomePage)} />
      </div>
      <div>
        <label>Photo</label>
        <Input placeholder={parseObject(item.Photo)} />
      </div>

      <Button label="Edit" />
    </div>
  )
}

EditCause.propTypes = {
  id: PropTypes.string,
}

export default EditCause
