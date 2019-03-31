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

const LIST_FIELDTYPE = gql`
  {
    listFieldType(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Title
      DataType
      PropName
      StringValueDefault
      IntValueDefault
      FloatValueDefault
      BooleanValueDefault
      DateTimeValueDefault {
        seconds
      }
      ComponentType {
        ID
      }
    }
  }
`

function ListFieldType({ navigate }) {
  const {
    data: { listFieldType: items = [] },
    error,
    loading,
  } = useQuery(LIST_FIELDTYPE)

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
          label="Create the first FieldType"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List FieldType</h1>
        <Button
          label="Create a new FieldType"
          onClick={() => navigate('create')}
          block
          variant="primary"
        />
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>CreatedAt</th>
              <th>UpdatedAt</th>
              <th>Title</th>
              <th>DataType</th>
              <th>PropName</th>
              <th>StringValueDefault</th>
              <th>IntValueDefault</th>
              <th>FloatValueDefault</th>
              <th>BooleanValueDefault</th>
              <th>DateTimeValueDefault</th>
              <th>ComponentType</th>
            </tr>
          </thead>
          {(items || []).map(item => (
            <tr>
              <td>
                <Link to={`/app/admin/field-type/${item.ID}`}>{item.ID}</Link>
              </td>
              <td>{parseObject(item.CreatedAt)}</td>
              <td>{parseObject(item.UpdatedAt)}</td>
              <td>{parseObject(item.Title)}</td>
              <td>{parseObject(item.DataType)}</td>
              <td>{parseObject(item.PropName)}</td>
              <td>{parseObject(item.StringValueDefault)}</td>
              <td>{parseObject(item.IntValueDefault)}</td>
              <td>{parseObject(item.FloatValueDefault)}</td>
              <td>{parseObject(item.BooleanValueDefault)}</td>
              <td>{parseObject(item.DateTimeValueDefault)}</td>
              <td>{parseObject(item.ComponentType)}</td>
            </tr>
          ))}
        </table>
      </Card>
    </Content>
  )
}

ListFieldType.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListFieldType
