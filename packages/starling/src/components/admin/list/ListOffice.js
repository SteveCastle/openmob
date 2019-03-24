import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_OFFICE = gql`
  {
    listOffice(limit: 20) {
      ID
    }
  }
`

function ListOffice({ navigate }) {
  const {
    data: { listOffice: items = [] },
    error,
    loading,
  } = useQuery(LIST_OFFICE)

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
          label="Create the first Office"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Office</h1>
      <Button label="Create a new Office" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/office/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListOffice.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListOffice
