import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_CONTACT = gql`
  {
    listContact(limit: 20) {
      ID
    }
  }
`

function ListContact({ navigate }) {
  const {
    data: { listContact: items = [] },
    error,
    loading,
  } = useQuery(LIST_CONTACT)

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
          label="Create the first Contact"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Contact</h1>
      <Button label="Create a new Contact" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/contact/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListContact.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListContact
