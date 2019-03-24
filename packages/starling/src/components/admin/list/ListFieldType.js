import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_FIELDTYPE = gql`
  {
    listFieldType(limit: 20) {
      ID
    }
  }
`

function ListFieldType({ navigate }) {
  const {
    data: { listFieldType: items = [] },
    error,
    loading,
  } = useQuery(LIST_FIELDTYPE)

  if (loading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <div>Error! {error.message}</div>
  }

  if (items === null || items.length === 0) {
    return (
      <div>
        <Button
          label="Create the first FieldType"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List FieldType</h1>
      <Button
        label="Create a new FieldType"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/fieldtype/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListFieldType.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListFieldType
