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
        />
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>CreatedAt</th>
              <th>UpdatedAt</th>
              <th>FirstName</th>
              <th>MiddleName</th>
              <th>LastName</th>
              <th>Email</th>
              <th>PhoneNumber</th>
            </tr>
          </thead>
          {(items || []).map(item => (
            <tr>
              <td>
                <Link to={`/app/admin/contact/${item.ID}`}>{item.ID}</Link>
              </td>
              <td>{parseObject(item.CreatedAt)}</td>
              <td>{parseObject(item.UpdatedAt)}</td>
              <td>{parseObject(item.FirstName)}</td>
              <td>{parseObject(item.MiddleName)}</td>
              <td>{parseObject(item.LastName)}</td>
              <td>{parseObject(item.Email)}</td>
              <td>{parseObject(item.PhoneNumber)}</td>
            </tr>
          ))}
        </table>
      </Card>
    </Content>
  )
}

ListContact.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListContact
