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

const LIST_CONTACT = gql`
  {
    listContact(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      FirstName
      MiddleName
      LastName
      Email
      PhoneNumber
    }
  }
`

function ListContact({ navigate }) {
  const {
    data: { listContact: items = [] },
    error,
    loading,
  } = useQuery(LIST_CONTACT)

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
          label="Create the first Contact"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List Contact</h1>
        <Button
          label="Create a new Contact"
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
              <TableHeaderCell>FirstName</TableHeaderCell>
              <TableHeaderCell>MiddleName</TableHeaderCell>
              <TableHeaderCell>LastName</TableHeaderCell>
              <TableHeaderCell>Email</TableHeaderCell>
              <TableHeaderCell>PhoneNumber</TableHeaderCell>
            </TableRow>
          </TableHeader>
          {(items || []).map(item => (
            <TableRow>
              <TableCell>
                <Link to={`/app/admin/contact/${item.ID}`}>{item.ID}</Link>
              </TableCell>
              <TableCell>{parseObject(item.CreatedAt)}</TableCell>
              <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
              <TableCell>{parseObject(item.FirstName)}</TableCell>
              <TableCell>{parseObject(item.MiddleName)}</TableCell>
              <TableCell>{parseObject(item.LastName)}</TableCell>
              <TableCell>{parseObject(item.Email)}</TableCell>
              <TableCell>{parseObject(item.PhoneNumber)}</TableCell>
            </TableRow>
          ))}
        </DataTable>
      </Card>
    </Content>
  )
}

ListContact.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListContact
