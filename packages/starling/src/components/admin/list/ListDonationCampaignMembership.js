import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_DONATIONCAMPAIGNMEMBERSHIP = gql`
  {
    listDonationCampaignMembership(limit: 20) {
      ID
    }
  }
`

function ListDonationCampaignMembership({ navigate }) {
  const {
    data: { listDonationCampaignMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_DONATIONCAMPAIGNMEMBERSHIP)

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
          label="Create the first DonationCampaignMembership"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List DonationCampaignMembership</h1>
      <Button
        label="Create a new DonationCampaignMembership"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/donationcampaignmembership/${item.ID}`}>
            {item.ID}
          </Link>
        </li>
      ))}
    </div>
  )
}

ListDonationCampaignMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListDonationCampaignMembership
