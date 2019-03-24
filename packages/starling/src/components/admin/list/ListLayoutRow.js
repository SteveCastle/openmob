import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_LAYOUTROW = gql`
  {
    listLayoutRow(limit: 20) {
      ID
    }
  }
`

function ListLayoutRow({ navigate }) {
  const {
    data: { listLayoutRow: items = [] },
    error,
    loading,
  } = useQuery(LIST_LAYOUTROW)

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
          label="Create the first LayoutRow"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List LayoutRow</h1>
      <Button
        label="Create a new LayoutRow"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/layoutrow/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListLayoutRow.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListLayoutRow
