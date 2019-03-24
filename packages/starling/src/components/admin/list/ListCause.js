import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_CAUSE = gql`
  {
    listCause(limit: 20) {
      ID
    }
  }
`

function ListCause({ navigate }) {
  const {
    data: { listCause: items = [] },
    error,
    loading,
  } = useQuery(LIST_CAUSE)

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
          label="Create the first Cause"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Cause</h1>
      <Button label="Create a new Cause" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/cause/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListCause.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListCause
