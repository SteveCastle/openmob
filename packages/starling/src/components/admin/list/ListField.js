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

const LIST_FIELD = gql`
  {
    listField(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      FieldType {
        ID
      }
      StringValue
      IntValue
      FloatValue
      BooleanValue
      DateTimeValue {
        seconds
      }
      Component {
        ID
      }
    }
  }
`

function ListField({ navigate }) {
  const {
    data: { listField: items = [] },
    error,
    loading,
  } = useQuery(LIST_FIELD)

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
          label="Create the first Field"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Field</h1>
      <Button label="Create a new Field" onClick={() => navigate('create')} />
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>CreatedAt</th>
            <th>UpdatedAt</th>
            <th>FieldType</th>
            <th>StringValue</th>
            <th>IntValue</th>
            <th>FloatValue</th>
            <th>BooleanValue</th>
            <th>DateTimeValue</th>
            <th>Component</th>
          </tr>
        </thead>
        {(items || []).map(item => (
          <tr>
            <td>
              <Link to={`/app/admin/field/${item.ID}`}>{item.ID}</Link>
            </td>
            <td>{parseObject(item.CreatedAt)}</td>
            <td>{parseObject(item.UpdatedAt)}</td>
            <td>{parseObject(item.FieldType)}</td>
            <td>{parseObject(item.StringValue)}</td>
            <td>{parseObject(item.IntValue)}</td>
            <td>{parseObject(item.FloatValue)}</td>
            <td>{parseObject(item.BooleanValue)}</td>
            <td>{parseObject(item.DateTimeValue)}</td>
            <td>{parseObject(item.Component)}</td>
          </tr>
        ))}
      </table>
    </div>
  )
}

ListField.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListField
