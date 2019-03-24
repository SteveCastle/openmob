import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_LAYOUT = gql`
  {
    listLayout(limit: 20) {
      ID
    }
  }
`

function ListLayout({ navigate }) {
  const {
    data: { listLayout: items = [] },
    error,
    loading,
  } = useQuery(LIST_LAYOUT)

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
          label="Create the first Layout"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Layout</h1>
      <Button label="Create a new Layout" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/layout/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListLayout.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListLayout
