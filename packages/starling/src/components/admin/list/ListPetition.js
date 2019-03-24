import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_PETITION = gql`
  {
    listPetition(limit: 20) {
      ID
    }
  }
`

function ListPetition({ navigate }) {
  const {
    data: { listPetition: items = [] },
    error,
    loading,
  } = useQuery(LIST_PETITION)

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
          label="Create the first Petition"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Petition</h1>
      <Button
        label="Create a new Petition"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/petition/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListPetition.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPetition
