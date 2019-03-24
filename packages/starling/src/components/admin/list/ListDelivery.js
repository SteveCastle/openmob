import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_DELIVERY = gql`
  {
    listDelivery(limit: 20) {
      ID
    }
  }
`

function ListDelivery({ navigate }) {
  const {
    data: { listDelivery: items = [] },
    error,
    loading,
  } = useQuery(LIST_DELIVERY)

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
          label="Create the first Delivery"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Delivery</h1>
      <Button
        label="Create a new Delivery"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/delivery/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListDelivery.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListDelivery
