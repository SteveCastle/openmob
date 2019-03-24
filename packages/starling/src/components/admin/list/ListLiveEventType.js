import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_LIVEEVENTTYPE = gql`
  {
    listLiveEventType(limit: 20) {
      ID
    }
  }
`

function ListLiveEventType({ navigate }) {
  const {
    data: { listLiveEventType: items = [] },
    error,
    loading,
  } = useQuery(LIST_LIVEEVENTTYPE)

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
          label="Create the first LiveEventType"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List LiveEventType</h1>
      <Button
        label="Create a new LiveEventType"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/liveeventtype/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListLiveEventType.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListLiveEventType
