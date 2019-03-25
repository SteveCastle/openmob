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

const LIST_COMPONENT = gql`
  {
    listComponent(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      ComponentType {
        ID
      }
      ComponentImplementation {
        ID
      }
      LayoutColumn {
        ID
      }
    }
  }
`

function ListComponent({ navigate }) {
  const {
    data: { listComponent: items = [] },
    error,
    loading,
  } = useQuery(LIST_COMPONENT)

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
          label="Create the first Component"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Component</h1>
      <Button
        label="Create a new Component"
        onClick={() => navigate('create')}
      />
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>CreatedAt</th>
            <th>UpdatedAt</th>
            <th>ComponentType</th>
            <th>ComponentImplementation</th>
            <th>LayoutColumn</th>
          </tr>
        </thead>
        {(items || []).map(item => (
          <tr>
            <td>
              <Link to={`/app/admin/component/${item.ID}`}>{item.ID}</Link>
            </td>
            <td>{parseObject(item.CreatedAt)}</td>
            <td>{parseObject(item.UpdatedAt)}</td>
            <td>{parseObject(item.ComponentType)}</td>
            <td>{parseObject(item.ComponentImplementation)}</td>
            <td>{parseObject(item.LayoutColumn)}</td>
          </tr>
        ))}
      </table>
    </div>
  )
}

ListComponent.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListComponent
