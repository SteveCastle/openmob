import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_CUSTOMERORDER = gql`
  {
    listCustomerOrder(limit: 20) {
      ID
    }
  }
`

function ListCustomerOrder({ navigate }) {
  const {
    data: { listCustomerOrder: items = [] },
    error,
    loading,
  } = useQuery(LIST_CUSTOMERORDER)

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
          label="Create the first CustomerOrder"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List CustomerOrder</h1>
      <Button
        label="Create a new CustomerOrder"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/customerorder/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListCustomerOrder.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListCustomerOrder
