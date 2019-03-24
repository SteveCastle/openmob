import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_ACL = gql`
  {
    listACL(limit: 20) {
      ID
    }
  }
`

function ListACL({ navigate }) {
  const {
    data: { listACL: items = [] },
    error,
    loading,
  } = useQuery(LIST_ACL)

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
          label="Create the first ACL"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List ACL</h1>
      <Button label="Create a new ACL" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/acl/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListACL.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListACL
