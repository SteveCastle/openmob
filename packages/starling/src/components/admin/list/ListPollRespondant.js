import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_POLLRESPONDANT = gql`
  {
    listPollRespondant(limit: 20) {
      ID
    }
  }
`

function ListPollRespondant({ navigate }) {
  const {
    data: { listPollRespondant: items = [] },
    error,
    loading,
  } = useQuery(LIST_POLLRESPONDANT)

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
          label="Create the first PollRespondant"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List PollRespondant</h1>
      <Button
        label="Create a new PollRespondant"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/pollrespondant/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListPollRespondant.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPollRespondant
