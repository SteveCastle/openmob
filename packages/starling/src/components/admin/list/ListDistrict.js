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
import parseObject from '../../../common/helpers'

const LIST_DISTRICT = gql`
  {
    listDistrict(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Geom
      Title
      DistrictType {
        ID
      }
    }
  }
`

function ListDistrict({ navigate = () => {} }) {
  const {
    data: { listDistrict: items = [] },
    error,
    loading,
  } = useQuery(LIST_DISTRICT)

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
          label="Create the first District"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List District</h1>
        <Button
          label="Create a new District"
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
              <TableHeaderCell>Geom</TableHeaderCell>
              <TableHeaderCell>Title</TableHeaderCell>
              <TableHeaderCell>DistrictType</TableHeaderCell>
            </TableRow>
          </TableHeader>
          <tbody>
            {(items || []).map(item => (
              <TableRow key={item.ID}>
                <TableCell>
                  <Link to={`/app/admin/district/${item.ID}`}>{item.ID}</Link>
                </TableCell>
                <TableCell>{parseObject(item.CreatedAt)}</TableCell>
                <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
                <TableCell>{parseObject(item.Geom)}</TableCell>
                <TableCell>{parseObject(item.Title)}</TableCell>
                <TableCell>{parseObject(item.DistrictType)}</TableCell>
              </TableRow>
            ))}
          </tbody>
        </DataTable>
      </Card>
    </Content>
  )
}

ListDistrict.propTypes = {
  navigate: PropTypes.func,
}

export default ListDistrict
