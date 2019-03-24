import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_COMPONENT = gql`
  {
    listComponent(limit: 20) {
      ID
    }
  }
`

function ListComponent({ navigate }) {
  const {
    data: { listComponent: items = [] },
    error,
    loading,
  } = useQuery(LIST_COMPONENT)

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
          label="Create the first Component"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Component</h1>
      <Button
        label="Create a new Component"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/component/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListComponent.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListComponent
