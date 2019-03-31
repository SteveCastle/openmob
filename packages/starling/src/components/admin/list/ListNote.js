import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const isObject = a => !!a && a.constructor === Object
const getValue = obj =>
  Object.entries(obj).reduce(entry =>
    entry[0] === 'seconds' || entry[0] === 'ID' ? entry[1] : null
  )
const parseObject = obj => (isObject(obj) ? getValue(obj) : obj)

const LIST_NOTE = gql`
  {
    listNote(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Contact {
        ID
      }
      Cause {
        ID
      }
      Body
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
      <Content>
        <Button
          label="Create the first Note"
          onClick={() => navigate('create')}
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List Note</h1>
        <Button
          label="Create a new Note"
          onClick={() => navigate('create')}
          block
        />
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>CreatedAt</th>
              <th>UpdatedAt</th>
              <th>Contact</th>
              <th>Cause</th>
              <th>Body</th>
            </tr>
          </thead>
          {(items || []).map(item => (
            <tr>
              <td>
                <Link to={`/app/admin/note/${item.ID}`}>{item.ID}</Link>
              </td>
              <td>{parseObject(item.CreatedAt)}</td>
              <td>{parseObject(item.UpdatedAt)}</td>
              <td>{parseObject(item.Contact)}</td>
              <td>{parseObject(item.Cause)}</td>
              <td>{parseObject(item.Body)}</td>
            </tr>
          ))}
        </table>
      </Card>
    </Content>
  )
}

ListNote.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListNote
