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

const LIST_LAYOUTCOLUMN = gql`
  {
    listLayoutColumn(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      LayoutRow {
        ID
      }
      Width
    }
  }
`

function ListLayoutColumn({ navigate }) {
  const {
    data: { listLayoutColumn: items = [] },
    error,
    loading,
  } = useQuery(LIST_LAYOUTCOLUMN)

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
          label="Create the first LayoutColumn"
          onClick={() => navigate('create')}
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List LayoutColumn</h1>
        <Button
          label="Create a new LayoutColumn"
          onClick={() => navigate('create')}
        />
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>CreatedAt</th>
              <th>UpdatedAt</th>
              <th>LayoutRow</th>
              <th>Width</th>
            </tr>
          </thead>
          {(items || []).map(item => (
            <tr>
              <td>
                <Link to={`/app/admin/layout-column/${item.ID}`}>
                  {item.ID}
                </Link>
              </td>
              <td>{parseObject(item.CreatedAt)}</td>
              <td>{parseObject(item.UpdatedAt)}</td>
              <td>{parseObject(item.LayoutRow)}</td>
              <td>{parseObject(item.Width)}</td>
            </tr>
          ))}
        </table>
      </Card>
    </Content>
  )
}

ListLayoutColumn.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListLayoutColumn
