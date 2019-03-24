import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_ELECTION = gql`
  {
    listElection(limit: 20) {
      ID
    }
  }
`

function ListElection({ navigate }) {
  const {
    data: { listElection: items = [] },
    error,
    loading,
  } = useQuery(LIST_ELECTION)

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
          label="Create the first Election"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Election</h1>
      <Button
        label="Create a new Election"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/election/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListElection.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListElection
