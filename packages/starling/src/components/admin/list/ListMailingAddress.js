import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_MAILINGADDRESS = gql`
  {
    listMailingAddress(limit: 20) {
      ID
    }
  }
`

function ListMailingAddress({ navigate }) {
  const {
    data: { listMailingAddress: items = [] },
    error,
    loading,
  } = useQuery(LIST_MAILINGADDRESS)

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
          label="Create the first MailingAddress"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List MailingAddress</h1>
      <Button
        label="Create a new MailingAddress"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/mailingaddress/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListMailingAddress.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListMailingAddress
