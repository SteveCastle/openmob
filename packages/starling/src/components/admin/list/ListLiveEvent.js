import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_LIVEEVENT = gql`
  {
    listLiveEvent(limit: 20) {
      ID
    }
  }
`

function ListLiveEvent({ navigate }) {
  const {
    data: { listLiveEvent: items = [] },
    error,
    loading,
  } = useQuery(LIST_LIVEEVENT)

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
          label="Create the first LiveEvent"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List LiveEvent</h1>
      <Button
        label="Create a new LiveEvent"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/liveevent/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListLiveEvent.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListLiveEvent
