import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_PRODUCTTYPE = gql`
  {
    listProductType(limit: 20) {
      ID
    }
  }
`

function ListProductType({ navigate }) {
  const {
    data: { listProductType: items = [] },
    error,
    loading,
  } = useQuery(LIST_PRODUCTTYPE)

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
          label="Create the first ProductType"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List ProductType</h1>
      <Button
        label="Create a new ProductType"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/producttype/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListProductType.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListProductType
