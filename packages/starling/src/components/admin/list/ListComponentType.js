import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_COMPONENTTYPE = gql`
  {
    listComponentType(limit: 20) {
      ID
    }
  }
`

function ListComponentType({ navigate }) {
  const {
    data: { listComponentType: items = [] },
    error,
    loading,
  } = useQuery(LIST_COMPONENTTYPE)

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
          label="Create the first ComponentType"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List ComponentType</h1>
      <Button
        label="Create a new ComponentType"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/componenttype/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListComponentType.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListComponentType
