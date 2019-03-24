import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_VOLUNTEEROPPORTUNITYTYPE = gql`
  {
    listVolunteerOpportunityType(limit: 20) {
      ID
    }
  }
`

function ListVolunteerOpportunityType({ navigate }) {
  const {
    data: { listVolunteerOpportunityType: items = [] },
    error,
    loading,
  } = useQuery(LIST_VOLUNTEEROPPORTUNITYTYPE)

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
          label="Create the first VolunteerOpportunityType"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List VolunteerOpportunityType</h1>
      <Button
        label="Create a new VolunteerOpportunityType"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/volunteeropportunitytype/${item.ID}`}>
            {item.ID}
          </Link>
        </li>
      ))}
    </div>
  )
}

ListVolunteerOpportunityType.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListVolunteerOpportunityType
