import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_AGENT = gql`
  {
    listAgent(limit: 20) {
      ID
    }
  }
`

function ListAgent({ navigate }) {
  const {
    data: { listAgent: items = [] },
    error,
    loading,
  } = useQuery(LIST_AGENT)

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
          label="Create the first Agent"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Agent</h1>
      <Button label="Create a new Agent" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/agent/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListAgent.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListAgent
