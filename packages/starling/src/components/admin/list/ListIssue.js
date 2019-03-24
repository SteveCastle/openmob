import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_ISSUE = gql`
  {
    listIssue(limit: 20) {
      ID
    }
  }
`

function ListIssue({ navigate }) {
  const {
    data: { listIssue: items = [] },
    error,
    loading,
  } = useQuery(LIST_ISSUE)

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
          label="Create the first Issue"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Issue</h1>
      <Button label="Create a new Issue" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/issue/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListIssue.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListIssue
