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

const GET_PETITIONSIGNER = gql`
  query getPetitionSignerById($id: ID!) {
    getPetitionSigner(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Petition {
        ID
      }
      Contact {
        ID
      }
      Cause {
        ID
      }
    }
  }
`

function EditPetitionSigner({ id }) {
  const {
    data: { getPetitionSigner: item = {} },
    error,
    loading,
  } = useQuery(GET_PETITIONSIGNER, {
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
        <label>Petition</label>
        <Input placeholder={parseObject(item.Petition)} />
      </div>
      <div>
        <label>Contact</label>
        <Input placeholder={parseObject(item.Contact)} />
      </div>
      <div>
        <label>Cause</label>
        <Input placeholder={parseObject(item.Cause)} />
      </div>

      <Button label="Edit" />
    </div>
  )
}

EditPetitionSigner.propTypes = {
  id: PropTypes.string,
}

export default EditPetitionSigner
