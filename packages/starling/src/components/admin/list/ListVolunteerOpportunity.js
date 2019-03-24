import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_VOLUNTEEROPPORTUNITY = gql`
  {
    listVolunteerOpportunity(limit: 20) {
      ID
    }
  }
`

function ListVolunteerOpportunity({ navigate }) {
  const {
    data: { listVolunteerOpportunity: items = [] },
    error,
    loading,
  } = useQuery(LIST_VOLUNTEEROPPORTUNITY)

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
          label="Create the first VolunteerOpportunity"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List VolunteerOpportunity</h1>
      <Button
        label="Create a new VolunteerOpportunity"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/volunteeropportunity/${item.ID}`}>
            {item.ID}
          </Link>
        </li>
      ))}
    </div>
  )
}

ListVolunteerOpportunity.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListVolunteerOpportunity
