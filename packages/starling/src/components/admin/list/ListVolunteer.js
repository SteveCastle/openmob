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

const LIST_VOLUNTEER = gql`
  {
    listVolunteer(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      VolunteerOpportunity {
        ID
      }
      Contact {
        ID
      }
      Cause {
        ID
      }
    }
  }
`

function ListVolunteer({ navigate = () => {} }) {
  const {
    data: { listVolunteer: items = [] },
    error,
    loading,
  } = useQuery(LIST_VOLUNTEER)

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
          label="Create the first Volunteer"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List Volunteer</h1>
        <Button
          label="Create a new Volunteer"
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
              <TableHeaderCell>VolunteerOpportunity</TableHeaderCell>
              <TableHeaderCell>Contact</TableHeaderCell>
              <TableHeaderCell>Cause</TableHeaderCell>
            </TableRow>
          </TableHeader>
          <tbody>
            {(items || []).map(item => (
              <TableRow key={item.ID}>
                <TableCell>
                  <Link to={`/app/admin/volunteer/${parseObject(item.ID)}`}>
                    {parseObject(item.ID)}
                  </Link>
                </TableCell>
                <TableCell>{parseObject(item.CreatedAt)}</TableCell>
                <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
                <TableCell>
                  <Link
                    to={`/app/admin/volunteer-opportunity/${parseObject(
                      item.VolunteerOpportunity
                    )}`}
                  >
                    {parseObject(item.VolunteerOpportunity)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link to={`/app/admin/contact/${parseObject(item.Contact)}`}>
                    {parseObject(item.Contact)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link to={`/app/admin/cause/${parseObject(item.Cause)}`}>
                    {parseObject(item.Cause)}
                  </Link>
                </TableCell>
              </TableRow>
            ))}
          </tbody>
        </DataTable>
      </Card>
    </Content>
  )
}

ListVolunteer.propTypes = {
  navigate: PropTypes.func,
}

export default ListVolunteer
