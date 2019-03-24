import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_VOLUNTEER = gql`
  {
    listVolunteer(limit: 20) {
      ID
    }
  }
`

function ListVolunteer({ navigate }) {
  const {
    data: { listVolunteer: items = [] },
    error,
    loading,
  } = useQuery(LIST_VOLUNTEER)

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
          label="Create the first Volunteer"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Volunteer</h1>
      <Button
        label="Create a new Volunteer"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/volunteer/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListVolunteer.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListVolunteer
