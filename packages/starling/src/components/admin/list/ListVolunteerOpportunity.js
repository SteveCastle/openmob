import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const isObject = a => !!a && a.constructor === Object
const getValue = obj =>
  Object.entries(obj).reduce(entry =>
    entry[0] === 'seconds' || entry[0] === 'ID' ? entry[1] : null
  )
const parseObject = obj => (isObject(obj) ? getValue(obj) : obj)

const LIST_VOLUNTEEROPPORTUNITY = gql`
  {
    listVolunteerOpportunity(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Title
      VolunteerOpportunityType {
        ID
      }
    }
  }
`

function ListVolunteerOpportunity({ navigate }) {
  const {
    data: { listVolunteerOpportunity: items = [] },
    error,
    loading,
  } = useQuery(LIST_VOLUNTEEROPPORTUNITY)

  if (loading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <div>Error! {error.message}</div>
  }

  if (items === null || items.length === 0) {
    return (
      <div>
        <Button
          label="Create the first VolunteerOpportunity"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List VolunteerOpportunity</h1>
      <Button
        label="Create a new VolunteerOpportunity"
        onClick={() => navigate('create')}
      />
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>CreatedAt</th>
            <th>UpdatedAt</th>
            <th>Title</th>
            <th>VolunteerOpportunityType</th>
          </tr>
        </thead>
        {(items || []).map(item => (
          <tr>
            <td>
              <Link to={`/app/admin/volunteer-opportunity/${item.ID}`}>
                {item.ID}
              </Link>
            </td>
            <td>{parseObject(item.CreatedAt)}</td>
            <td>{parseObject(item.UpdatedAt)}</td>
            <td>{parseObject(item.Title)}</td>
            <td>{parseObject(item.VolunteerOpportunityType)}</td>
          </tr>
        ))}
      </table>
    </div>
  )
}

ListVolunteerOpportunity.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListVolunteerOpportunity
