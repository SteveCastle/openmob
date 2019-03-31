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

const LIST_VOLUNTEEROPPORTUNITYMEMBERSHIP = gql`
  {
    listVolunteerOpportunityMembership(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Cause {
        ID
      }
      VolunteerOpportunity {
        ID
      }
    }
  }
`

function ListVolunteerOpportunityMembership({ navigate }) {
  const {
    data: { listVolunteerOpportunityMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_VOLUNTEEROPPORTUNITYMEMBERSHIP)

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
          label="Create the first VolunteerOpportunityMembership"
          onClick={() => navigate('create')}
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List VolunteerOpportunityMembership</h1>
        <Button
          label="Create a new VolunteerOpportunityMembership"
          onClick={() => navigate('create')}
        />
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>CreatedAt</th>
              <th>UpdatedAt</th>
              <th>Cause</th>
              <th>VolunteerOpportunity</th>
            </tr>
          </thead>
          {(items || []).map(item => (
            <tr>
              <td>
                <Link
                  to={`/app/admin/volunteer-opportunity-membership/${item.ID}`}
                >
                  {item.ID}
                </Link>
              </td>
              <td>{parseObject(item.CreatedAt)}</td>
              <td>{parseObject(item.UpdatedAt)}</td>
              <td>{parseObject(item.Cause)}</td>
              <td>{parseObject(item.VolunteerOpportunity)}</td>
            </tr>
          ))}
        </table>
      </Card>
    </Content>
  )
}

ListVolunteerOpportunityMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListVolunteerOpportunityMembership
