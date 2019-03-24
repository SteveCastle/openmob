import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_POLLMEMBERSHIP = gql`
  {
    listPollMembership(limit: 20) {
      ID
    }
  }
`

function ListPollMembership({ navigate }) {
  const {
    data: { listPollMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_POLLMEMBERSHIP)

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
          label="Create the first PollMembership"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List PollMembership</h1>
      <Button
        label="Create a new PollMembership"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/pollmembership/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListPollMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPollMembership
