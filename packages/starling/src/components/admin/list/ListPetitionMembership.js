import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_PETITIONMEMBERSHIP = gql`
  {
    listPetitionMembership(limit: 20) {
      ID
    }
  }
`

function ListPetitionMembership({ navigate }) {
  const {
    data: { listPetitionMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_PETITIONMEMBERSHIP)

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
          label="Create the first PetitionMembership"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List PetitionMembership</h1>
      <Button
        label="Create a new PetitionMembership"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/petitionmembership/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListPetitionMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPetitionMembership
