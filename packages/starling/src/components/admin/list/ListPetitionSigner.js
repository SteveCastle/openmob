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

const LIST_PETITIONSIGNER = gql`
  {
    listPetitionSigner(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Petition {
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

function ListPetitionSigner({ navigate }) {
  const {
    data: { listPetitionSigner: items = [] },
    error,
    loading,
  } = useQuery(LIST_PETITIONSIGNER)

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
          label="Create the first PetitionSigner"
          onClick={() => navigate('create')}
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List PetitionSigner</h1>
        <Button
          label="Create a new PetitionSigner"
          onClick={() => navigate('create')}
          block
        />
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>CreatedAt</th>
              <th>UpdatedAt</th>
              <th>Petition</th>
              <th>Contact</th>
              <th>Cause</th>
            </tr>
          </thead>
          {(items || []).map(item => (
            <tr>
              <td>
                <Link to={`/app/admin/petition-signer/${item.ID}`}>
                  {item.ID}
                </Link>
              </td>
              <td>{parseObject(item.CreatedAt)}</td>
              <td>{parseObject(item.UpdatedAt)}</td>
              <td>{parseObject(item.Petition)}</td>
              <td>{parseObject(item.Contact)}</td>
              <td>{parseObject(item.Cause)}</td>
            </tr>
          ))}
        </table>
      </Card>
    </Content>
  )
}

ListPetitionSigner.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListPetitionSigner
