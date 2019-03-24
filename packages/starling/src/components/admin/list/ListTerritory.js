import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_TERRITORY = gql`
  {
    listTerritory(limit: 20) {
      ID
    }
  }
`

function ListTerritory({ navigate }) {
  const {
    data: { listTerritory: items = [] },
    error,
    loading,
  } = useQuery(LIST_TERRITORY)

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
          label="Create the first Territory"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Territory</h1>
      <Button
        label="Create a new Territory"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/territory/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListTerritory.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListTerritory
