import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_VOLUNTEEROPPORTUNITYMEMBERSHIP = gql`
  {
    listVolunteerOpportunityMembership(limit: 20) {
      ID
    }
  }
`

function ListVolunteerOpportunityMembership({ navigate }) {
  const {
    data: { listVolunteerOpportunityMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_VOLUNTEEROPPORTUNITYMEMBERSHIP)

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
          label="Create the first VolunteerOpportunityMembership"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List VolunteerOpportunityMembership</h1>
      <Button
        label="Create a new VolunteerOpportunityMembership"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/volunteeropportunitymembership/${item.ID}`}>
            {item.ID}
          </Link>
        </li>
      ))}
    </div>
  )
}

ListVolunteerOpportunityMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListVolunteerOpportunityMembership
