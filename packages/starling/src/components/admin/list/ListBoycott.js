import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_BOYCOTT = gql`
  {
    listBoycott(limit: 20) {
      ID
    }
  }
`

function ListBoycott({ navigate }) {
  const {
    data: { listBoycott: items = [] },
    error,
    loading,
  } = useQuery(LIST_BOYCOTT)

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
          label="Create the first Boycott"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Boycott</h1>
      <Button label="Create a new Boycott" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/boycott/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListBoycott.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListBoycott
