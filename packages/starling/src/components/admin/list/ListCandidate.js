import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_CANDIDATE = gql`
  {
    listCandidate(limit: 20) {
      ID
    }
  }
`

function ListCandidate({ navigate }) {
  const {
    data: { listCandidate: items = [] },
    error,
    loading,
  } = useQuery(LIST_CANDIDATE)

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
          label="Create the first Candidate"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Candidate</h1>
      <Button
        label="Create a new Candidate"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/candidate/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListCandidate.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListCandidate
