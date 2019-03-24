import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_BOYCOTTMEMBERSHIP = gql`
  {
    listBoycottMembership(limit: 20) {
      ID
    }
  }
`

function ListBoycottMembership({ navigate }) {
  const {
    data: { listBoycottMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_BOYCOTTMEMBERSHIP)

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
          label="Create the first BoycottMembership"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List BoycottMembership</h1>
      <Button
        label="Create a new BoycottMembership"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/boycottmembership/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListBoycottMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListBoycottMembership
