import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_NOTE = gql`
  {
    listNote(limit: 20) {
      ID
    }
  }
`

function ListNote({ navigate }) {
  const {
    data: { listNote: items = [] },
    error,
    loading,
  } = useQuery(LIST_NOTE)

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
          label="Create the first Note"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Note</h1>
      <Button label="Create a new Note" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/note/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListNote.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListNote
