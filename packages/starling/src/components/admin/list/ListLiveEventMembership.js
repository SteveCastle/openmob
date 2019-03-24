import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_LIVEEVENTMEMBERSHIP = gql`
  {
    listLiveEventMembership(limit: 20) {
      ID
    }
  }
`

function ListLiveEventMembership({ navigate }) {
  const {
    data: { listLiveEventMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_LIVEEVENTMEMBERSHIP)

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
          label="Create the first LiveEventMembership"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List LiveEventMembership</h1>
      <Button
        label="Create a new LiveEventMembership"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/liveeventmembership/${item.ID}`}>
            {item.ID}
          </Link>
        </li>
      ))}
    </div>
  )
}

ListLiveEventMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListLiveEventMembership
