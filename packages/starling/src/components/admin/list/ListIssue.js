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

const LIST_ISSUE = gql`
  {
    listIssue(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Title
      Election {
        ID
      }
    }
  }
`

function ListIssue({ navigate = () => {} }) {
  const {
    data: { listIssue: items = [] },
    error,
    loading,
  } = useQuery(LIST_ISSUE)

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
          label="Create the first Issue"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List Issue</h1>
        <Button
          label="Create a new Issue"
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
              <TableHeaderCell>Election</TableHeaderCell>
            </TableRow>
          </TableHeader>
          <tbody>
            {(items || []).map(item => (
              <TableRow key={item.ID}>
                <TableCell>
                  <Link to={`/app/admin/issue/${item.ID}`}>{item.ID}</Link>
                </TableCell>
                <TableCell>{parseObject(item.CreatedAt)}</TableCell>
                <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
                <TableCell>{parseObject(item.Title)}</TableCell>
                <TableCell>{parseObject(item.Election)}</TableCell>
              </TableRow>
            ))}
          </tbody>
        </DataTable>
      </Card>
    </Content>
  )
}

ListIssue.propTypes = {
  navigate: PropTypes.func,
}

export default ListIssue