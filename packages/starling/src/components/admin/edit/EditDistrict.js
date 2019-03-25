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

const GET_DISTRICT = gql`
  query getDistrictById($id: ID!) {
    getDistrict(ID: $id) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Geom
      Title
      DistrictType {
        ID
      }
    }
  }
`

function EditDistrict({ id }) {
  const {
    data: { getDistrict: item = {} },
    error,
    loading,
  } = useQuery(GET_DISTRICT, {
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
        <label>Geom</label>
        <Input placeholder={parseObject(item.Geom)} />
      </div>
      <div>
        <label>Title</label>
        <Input placeholder={parseObject(item.Title)} />
      </div>
      <div>
        <label>DistrictType</label>
        <Input placeholder={parseObject(item.DistrictType)} />
      </div>

      <Button label="Edit" />
    </div>
  )
}

EditDistrict.propTypes = {
  id: PropTypes.string,
}

export default EditDistrict
