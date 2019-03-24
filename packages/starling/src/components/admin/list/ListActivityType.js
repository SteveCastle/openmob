import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_ACTIVITYTYPE = gql`
  {
    listActivityType(limit: 20) {
      ID
    }
  }
`

function ListActivityType({ navigate }) {
  const {
    data: { listActivityType: items = [] },
    error,
    loading,
  } = useQuery(LIST_ACTIVITYTYPE)

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
          label="Create the first ActivityType"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List ActivityType</h1>
      <Button
        label="Create a new ActivityType"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/activitytype/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListActivityType.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListActivityType
