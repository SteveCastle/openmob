import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Button from '@openmob/bluebird/src/components/buttons/Button'
import DataTable from '@openmob/bluebird/src/components/tables/DataTable'
import TableHeader from '@openmob/bluebird/src/components/tables/TableHeader'
import TableHeaderCell from '@openmob/bluebird/src/components/tables/TableHeaderCell'
import TableRow from '@openmob/bluebird/src/components/tables/TableRow'
import TableCell from '@openmob/bluebird/src/components/tables/TableCell'

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
      <Content>
        <Button
          label="Create the first Field"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List Field</h1>
        <Button
          label="Create a new Field"
          onClick={() => navigate('create')}
          block
          variant="primary"
        />
        <DataTable>
          <TableHeader>
            <TableRow>
              <TableHeaderCell>ID</TableHeaderCell>
              <TableHeaderCell>CreatedAt</TableHeaderCell>
              <TableHeaderCell>UpdatedAt</TableHeaderCell>
              <TableHeaderCell>FieldType</TableHeaderCell>
              <TableHeaderCell>StringValue</TableHeaderCell>
              <TableHeaderCell>IntValue</TableHeaderCell>
              <TableHeaderCell>FloatValue</TableHeaderCell>
              <TableHeaderCell>BooleanValue</TableHeaderCell>
              <TableHeaderCell>DateTimeValue</TableHeaderCell>
              <TableHeaderCell>Component</TableHeaderCell>
            </TableRow>
          </TableHeader>
          {(items || []).map(item => (
            <TableRow>
              <TableCell>
                <Link to={`/app/admin/field/${item.ID}`}>{item.ID}</Link>
              </TableCell>
              <TableCell>{parseObject(item.CreatedAt)}</TableCell>
              <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
              <TableCell>{parseObject(item.FieldType)}</TableCell>
              <TableCell>{parseObject(item.StringValue)}</TableCell>
              <TableCell>{parseObject(item.IntValue)}</TableCell>
              <TableCell>{parseObject(item.FloatValue)}</TableCell>
              <TableCell>{parseObject(item.BooleanValue)}</TableCell>
              <TableCell>{parseObject(item.DateTimeValue)}</TableCell>
              <TableCell>{parseObject(item.Component)}</TableCell>
            </TableRow>
          ))}
        </DataTable>
      </Card>
    </Content>
  )
}

ListField.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListField
