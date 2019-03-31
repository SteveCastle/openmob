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
        <DataTable>
          <TableHeader>
            <TableRow>
              <TableHeaderCell>ID</TableHeaderCell>
              <TableHeaderCell>CreatedAt</TableHeaderCell>
              <TableHeaderCell>UpdatedAt</TableHeaderCell>
              <TableHeaderCell>Title</TableHeaderCell>
              <TableHeaderCell>DataType</TableHeaderCell>
              <TableHeaderCell>PropName</TableHeaderCell>
              <TableHeaderCell>StringValueDefault</TableHeaderCell>
              <TableHeaderCell>IntValueDefault</TableHeaderCell>
              <TableHeaderCell>FloatValueDefault</TableHeaderCell>
              <TableHeaderCell>BooleanValueDefault</TableHeaderCell>
              <TableHeaderCell>DateTimeValueDefault</TableHeaderCell>
              <TableHeaderCell>ComponentType</TableHeaderCell>
            </TableRow>
          </TableHeader>
          {(items || []).map(item => (
            <TableRow>
              <TableCell>
                <Link to={`/app/admin/field-type/${item.ID}`}>{item.ID}</Link>
              </TableCell>
              <TableCell>{parseObject(item.CreatedAt)}</TableCell>
              <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
              <TableCell>{parseObject(item.Title)}</TableCell>
              <TableCell>{parseObject(item.DataType)}</TableCell>
              <TableCell>{parseObject(item.PropName)}</TableCell>
              <TableCell>{parseObject(item.StringValueDefault)}</TableCell>
              <TableCell>{parseObject(item.IntValueDefault)}</TableCell>
              <TableCell>{parseObject(item.FloatValueDefault)}</TableCell>
              <TableCell>{parseObject(item.BooleanValueDefault)}</TableCell>
              <TableCell>{parseObject(item.DateTimeValueDefault)}</TableCell>
              <TableCell>{parseObject(item.ComponentType)}</TableCell>
            </TableRow>
          ))}
        </DataTable>
      </Card>
    </Content>
  )
}

ListFieldType.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListFieldType
