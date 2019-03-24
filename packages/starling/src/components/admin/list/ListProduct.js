import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_PRODUCT = gql`
  {
    listProduct(limit: 20) {
      ID
    }
  }
`

function ListProduct({ navigate }) {
  const {
    data: { listProduct: items = [] },
    error,
    loading,
  } = useQuery(LIST_PRODUCT)

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
          label="Create the first Product"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Product</h1>
      <Button label="Create a new Product" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/product/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListProduct.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListProduct
