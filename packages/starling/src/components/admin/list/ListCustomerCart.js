import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_CUSTOMERCART = gql`
  {
    listCustomerCart(limit: 20) {
      ID
    }
  }
`

function ListCustomerCart({ navigate }) {
  const {
    data: { listCustomerCart: items = [] },
    error,
    loading,
  } = useQuery(LIST_CUSTOMERCART)

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
          label="Create the first CustomerCart"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List CustomerCart</h1>
      <Button
        label="Create a new CustomerCart"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/customercart/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListCustomerCart.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListCustomerCart
