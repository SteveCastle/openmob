import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_CONTACTMEMBERSHIP = gql`
  {
    listContactMembership(limit: 20) {
      ID
    }
  }
`

function ListContactMembership({ navigate }) {
  const {
    data: { listContactMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_CONTACTMEMBERSHIP)

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
          label="Create the first ContactMembership"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List ContactMembership</h1>
      <Button
        label="Create a new ContactMembership"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/contactmembership/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListContactMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListContactMembership
