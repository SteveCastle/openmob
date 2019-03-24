import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_PETITIONSIGNER = gql`
  {
    listPetitionSigner(limit: 20) {
      ID
    }
  }
`

function ListPetitionSigner({ navigate }) {
  const {
    data: { listPetitionSigner: items = [] },
    error,
    loading,
  } = useQuery(LIST_PETITIONSIGNER)

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
          label="Create the first PetitionSigner"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List PetitionSigner</h1>
      <Button
        label="Create a new PetitionSigner"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/petitionsigner/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListPetitionSigner.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPetitionSigner
