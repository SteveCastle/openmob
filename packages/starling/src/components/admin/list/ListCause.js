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

const LIST_CAUSE = gql`
  {
    listCause(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Title
      Slug
      Summary
      HomePage {
        ID
      }
      Photo {
        ID
      }
    }
  }
`

function ListCause({ navigate }) {
  const {
    data: { listCause: items = [] },
    error,
    loading,
  } = useQuery(LIST_CAUSE)

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
          label="Create the first Cause"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Cause</h1>
      <Button label="Create a new Cause" onClick={() => navigate('create')} />
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>CreatedAt</th>
            <th>UpdatedAt</th>
            <th>Title</th>
            <th>Slug</th>
            <th>Summary</th>
            <th>HomePage</th>
            <th>Photo</th>
          </tr>
        </thead>
        {(items || []).map(item => (
          <tr>
            <td>
              <Link to={`/app/admin/cause/${item.ID}`}>{item.ID}</Link>
            </td>
            <td>{parseObject(item.CreatedAt)}</td>
            <td>{parseObject(item.UpdatedAt)}</td>
            <td>{parseObject(item.Title)}</td>
            <td>{parseObject(item.Slug)}</td>
            <td>{parseObject(item.Summary)}</td>
            <td>{parseObject(item.HomePage)}</td>
            <td>{parseObject(item.Photo)}</td>
          </tr>
        ))}
      </table>
    </div>
  )
}

ListCause.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListCause
