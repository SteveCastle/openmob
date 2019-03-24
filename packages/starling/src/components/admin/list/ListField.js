import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_FIELD = gql`
  {
    listField(limit: 20) {
      ID
    }
  }
`

function ListField({ navigate }) {
  const {
    data: { listField: items = [] },
    error,
    loading,
  } = useQuery(LIST_FIELD)

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
          label="Create the first Field"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Field</h1>
      <Button label="Create a new Field" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/field/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListField.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListField
