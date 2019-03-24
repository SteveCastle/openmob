import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_PAYMENT = gql`
  {
    listPayment(limit: 20) {
      ID
    }
  }
`

function ListPayment({ navigate }) {
  const {
    data: { listPayment: items = [] },
    error,
    loading,
  } = useQuery(LIST_PAYMENT)

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
          label="Create the first Payment"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Payment</h1>
      <Button label="Create a new Payment" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/payment/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListPayment.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPayment
