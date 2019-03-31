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

const LIST_MAILINGADDRESS = gql`
  {
    listMailingAddress(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      StreetAddress
      City
      State
      ZipCode
    }
  }
`

function ListMailingAddress({ navigate }) {
  const {
    data: { listMailingAddress: items = [] },
    error,
    loading,
  } = useQuery(LIST_MAILINGADDRESS)

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
          label="Create the first MailingAddress"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List MailingAddress</h1>
        <Button
          label="Create a new MailingAddress"
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
              <TableHeaderCell>StreetAddress</TableHeaderCell>
              <TableHeaderCell>City</TableHeaderCell>
              <TableHeaderCell>State</TableHeaderCell>
              <TableHeaderCell>ZipCode</TableHeaderCell>
            </TableRow>
          </TableHeader>
          {(items || []).map(item => (
            <TableRow>
              <TableCell>
                <Link to={`/app/admin/mailing-address/${item.ID}`}>
                  {item.ID}
                </Link>
              </TableCell>
              <TableCell>{parseObject(item.CreatedAt)}</TableCell>
              <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
              <TableCell>{parseObject(item.StreetAddress)}</TableCell>
              <TableCell>{parseObject(item.City)}</TableCell>
              <TableCell>{parseObject(item.State)}</TableCell>
              <TableCell>{parseObject(item.ZipCode)}</TableCell>
            </TableRow>
          ))}
        </DataTable>
      </Card>
    </Content>
  )
}

ListMailingAddress.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListMailingAddress
