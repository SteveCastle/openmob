import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_OWNERMEMBERSHIP = gql`
  {
    listOwnerMembership(limit: 20) {
      ID
    }
  }
`

function ListOwnerMembership({ navigate }) {
  const {
    data: { listOwnerMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_OWNERMEMBERSHIP)

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
          label="Create the first OwnerMembership"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List OwnerMembership</h1>
      <Button
        label="Create a new OwnerMembership"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/ownermembership/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListOwnerMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListOwnerMembership
