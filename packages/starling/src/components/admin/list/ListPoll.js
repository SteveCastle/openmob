import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const isObject = a => !!a && a.constructor === Object
const getValue = obj =>
  Object.entries(obj).reduce(entry =>
    entry[0] === 'seconds' || entry[0] === 'ID' ? entry[1] : null
  )
const parseObject = obj => (isObject(obj) ? getValue(obj) : obj)

const LIST_POLL = gql`
  {
    listPoll(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Title
    }
  }
`

function ListPoll({ navigate }) {
  const {
    data: { listPoll: items = [] },
    error,
    loading,
  } = useQuery(LIST_POLL)

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
          label="Create the first Poll"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Poll</h1>
      <Button label="Create a new Poll" onClick={() => navigate('create')} />
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>CreatedAt</th>
            <th>UpdatedAt</th>
            <th>Title</th>
          </tr>
        </thead>
        {(items || []).map(item => (
          <tr>
            <td>
              <Link to={`/app/admin/poll/${item.ID}`}>{item.ID}</Link>
            </td>
            <td>{parseObject(item.CreatedAt)}</td>
            <td>{parseObject(item.UpdatedAt)}</td>
            <td>{parseObject(item.Title)}</td>
          </tr>
        ))}
      </table>
    </div>
  )
}

ListPoll.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPoll
