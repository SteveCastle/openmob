import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_EMAILADDRESS = gql`
  {
    listEmailAddress(limit: 20) {
      ID
    }
  }
`

function ListEmailAddress({ navigate }) {
  const {
    data: { listEmailAddress: items = [] },
    error,
    loading,
  } = useQuery(LIST_EMAILADDRESS)

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
          label="Create the first EmailAddress"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List EmailAddress</h1>
      <Button
        label="Create a new EmailAddress"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/emailaddress/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListEmailAddress.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListEmailAddress
