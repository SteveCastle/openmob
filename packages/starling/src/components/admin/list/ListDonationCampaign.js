import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_DONATIONCAMPAIGN = gql`
  {
    listDonationCampaign(limit: 20) {
      ID
    }
  }
`

function ListDonationCampaign({ navigate }) {
  const {
    data: { listDonationCampaign: items = [] },
    error,
    loading,
  } = useQuery(LIST_DONATIONCAMPAIGN)

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
          label="Create the first DonationCampaign"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List DonationCampaign</h1>
      <Button
        label="Create a new DonationCampaign"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/donationcampaign/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListDonationCampaign.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListDonationCampaign
