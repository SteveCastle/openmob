import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_AGENTMEMBERSHIP = gql`
  {
    listAgentMembership(limit: 20) {
      ID
    }
  }
`

function ListAgentMembership({ navigate }) {
  const {
    data: { listAgentMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_AGENTMEMBERSHIP)

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
          label="Create the first AgentMembership"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List AgentMembership</h1>
      <Button
        label="Create a new AgentMembership"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/agentmembership/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListAgentMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListAgentMembership
