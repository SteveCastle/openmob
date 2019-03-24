import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_LAYOUTCOLUMN = gql`
  {
    listLayoutColumn(limit: 20) {
      ID
    }
  }
`

function ListLayoutColumn({ navigate }) {
  const {
    data: { listLayoutColumn: items = [] },
    error,
    loading,
  } = useQuery(LIST_LAYOUTCOLUMN)

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
          label="Create the first LayoutColumn"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List LayoutColumn</h1>
      <Button
        label="Create a new LayoutColumn"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/layoutcolumn/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListLayoutColumn.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListLayoutColumn
