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

const LIST_PETITIONMEMBERSHIP = gql`
  {
    listPetitionMembership(limit: 20) {
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
      Petition {
        ID
      }
    }
  }
`

function ListPetitionMembership({ navigate }) {
  const {
    data: { listPetitionMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_PETITIONMEMBERSHIP)

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
          label="Create the first PetitionMembership"
          onClick={() => navigate('create')}
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List PetitionMembership</h1>
        <Button
          label="Create a new PetitionMembership"
          onClick={() => navigate('create')}
          block
        />
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>CreatedAt</th>
              <th>UpdatedAt</th>
              <th>Cause</th>
              <th>Petition</th>
            </tr>
          </thead>
          {(items || []).map(item => (
            <tr>
              <td>
                <Link to={`/app/admin/petition-membership/${item.ID}`}>
                  {item.ID}
                </Link>
              </td>
              <td>{parseObject(item.CreatedAt)}</td>
              <td>{parseObject(item.UpdatedAt)}</td>
              <td>{parseObject(item.Cause)}</td>
              <td>{parseObject(item.Petition)}</td>
            </tr>
          ))}
        </table>
      </Card>
    </Content>
  )
}

ListPetitionMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPetitionMembership
