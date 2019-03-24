import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_PHONENUMBER = gql`
  {
    listPhoneNumber(limit: 20) {
      ID
    }
  }
`

function ListPhoneNumber({ navigate }) {
  const {
    data: { listPhoneNumber: items = [] },
    error,
    loading,
  } = useQuery(LIST_PHONENUMBER)

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
          label="Create the first PhoneNumber"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List PhoneNumber</h1>
      <Button
        label="Create a new PhoneNumber"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/phonenumber/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListPhoneNumber.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPhoneNumber
