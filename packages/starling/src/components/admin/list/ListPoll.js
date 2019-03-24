import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_POLL = gql`
  {
    listPoll(limit: 20) {
      ID
    }
  }
`

function ListPoll({ navigate }) {
  const {
    data: { listPoll: items = [] },
    error,
    loading,
  } = useQuery(LIST_POLL)

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
          label="Create the first Poll"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Poll</h1>
      <Button label="Create a new Poll" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/poll/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListPoll.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPoll
