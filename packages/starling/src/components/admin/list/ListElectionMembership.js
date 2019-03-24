import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_ELECTIONMEMBERSHIP = gql`
  {
    listElectionMembership(limit: 20) {
      ID
    }
  }
`

function ListElectionMembership({ navigate }) {
  const {
    data: { listElectionMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_ELECTIONMEMBERSHIP)

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
          label="Create the first ElectionMembership"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List ElectionMembership</h1>
      <Button
        label="Create a new ElectionMembership"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/electionmembership/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListElectionMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListElectionMembership
