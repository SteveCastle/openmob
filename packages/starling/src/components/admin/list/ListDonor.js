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

const LIST_DONOR = gql`
  {
    listDonor(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      CustomerOrder {
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

function ListDonor({ navigate }) {
  const {
    data: { listDonor: items = [] },
    error,
    loading,
  } = useQuery(LIST_DONOR)

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
          label="Create the first Donor"
          onClick={() => navigate('create')}
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List Donor</h1>
        <Button label="Create a new Donor" onClick={() => navigate('create')} />
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>CreatedAt</th>
              <th>UpdatedAt</th>
              <th>CustomerOrder</th>
              <th>Contact</th>
              <th>Cause</th>
            </tr>
          </thead>
          {(items || []).map(item => (
            <tr>
              <td>
                <Link to={`/app/admin/donor/${item.ID}`}>{item.ID}</Link>
              </td>
              <td>{parseObject(item.CreatedAt)}</td>
              <td>{parseObject(item.UpdatedAt)}</td>
              <td>{parseObject(item.CustomerOrder)}</td>
              <td>{parseObject(item.Contact)}</td>
              <td>{parseObject(item.Cause)}</td>
            </tr>
          ))}
        </table>
      </Card>
    </Content>
  )
}

ListDonor.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListDonor
