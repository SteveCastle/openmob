import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const isObject = a => !!a && a.constructor === Object
const getValue = obj =>
  Object.entries(obj).reduce(entry =>
    entry[0] === 'seconds' || entry[0] === 'ID' ? entry[1] : null
  )
const parseObject = obj => (isObject(obj) ? getValue(obj) : obj)

const LIST_DONATIONCAMPAIGN = gql`
  {
    listDonationCampaign(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Title
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
      <Content>
        <Button
          label="Create the first DonationCampaign"
          onClick={() => navigate('create')}
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List DonationCampaign</h1>
        <Button
          label="Create a new DonationCampaign"
          onClick={() => navigate('create')}
          block
        />
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>CreatedAt</th>
              <th>UpdatedAt</th>
              <th>Title</th>
            </tr>
          </thead>
          {(items || []).map(item => (
            <tr>
              <td>
                <Link to={`/app/admin/donation-campaign/${item.ID}`}>
                  {item.ID}
                </Link>
              </td>
              <td>{parseObject(item.CreatedAt)}</td>
              <td>{parseObject(item.UpdatedAt)}</td>
              <td>{parseObject(item.Title)}</td>
            </tr>
          ))}
        </table>
      </Card>
    </Content>
  )
}

ListDonationCampaign.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListDonationCampaign
