import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_ACCOUNT = gql`
  {
    listAccount(limit: 20) {
      ID
    }
  }
`

function ListAccount({ navigate }) {
  const {
    data: { listAccount: items = [] },
    error,
    loading,
  } = useQuery(LIST_ACCOUNT)

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
          label="Create the first Account"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Account</h1>
      <Button label="Create a new Account" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/account/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListAccount.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListAccount
