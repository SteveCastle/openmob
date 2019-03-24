import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_PURCHASER = gql`
  {
    listPurchaser(limit: 20) {
      ID
    }
  }
`

function ListPurchaser({ navigate }) {
  const {
    data: { listPurchaser: items = [] },
    error,
    loading,
  } = useQuery(LIST_PURCHASER)

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
          label="Create the first Purchaser"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Purchaser</h1>
      <Button
        label="Create a new Purchaser"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/purchaser/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListPurchaser.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPurchaser
