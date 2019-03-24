import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_ACTIVITY = gql`
  {
    listActivity(limit: 20) {
      ID
    }
  }
`

function ListActivity({ navigate }) {
  const {
    data: { listActivity: items = [] },
    error,
    loading,
  } = useQuery(LIST_ACTIVITY)

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
          label="Create the first Activity"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Activity</h1>
      <Button
        label="Create a new Activity"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/activity/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListActivity.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListActivity
