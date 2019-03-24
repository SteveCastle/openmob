import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_POLLITEM = gql`
  {
    listPollItem(limit: 20) {
      ID
    }
  }
`

function ListPollItem({ navigate }) {
  const {
    data: { listPollItem: items = [] },
    error,
    loading,
  } = useQuery(LIST_POLLITEM)

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
          label="Create the first PollItem"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List PollItem</h1>
      <Button
        label="Create a new PollItem"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/pollitem/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListPollItem.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPollItem
