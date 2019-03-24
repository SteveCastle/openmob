import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_LAYOUTTYPE = gql`
  {
    listLayoutType(limit: 20) {
      ID
    }
  }
`

function ListLayoutType({ navigate }) {
  const {
    data: { listLayoutType: items = [] },
    error,
    loading,
  } = useQuery(LIST_LAYOUTTYPE)

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
          label="Create the first LayoutType"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List LayoutType</h1>
      <Button
        label="Create a new LayoutType"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/layouttype/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListLayoutType.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListLayoutType
