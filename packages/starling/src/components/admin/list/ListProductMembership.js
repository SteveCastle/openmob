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

const LIST_PRODUCTMEMBERSHIP = gql`
  {
    listProductMembership(limit: 20) {
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
      Product {
        ID
      }
    }
  }
`

function ListProductMembership({ navigate }) {
  const {
    data: { listProductMembership: items = [] },
    error,
    loading,
  } = useQuery(LIST_PRODUCTMEMBERSHIP)

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
          label="Create the first ProductMembership"
          onClick={() => navigate('create')}
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List ProductMembership</h1>
        <Button
          label="Create a new ProductMembership"
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
              <th>Product</th>
            </tr>
          </thead>
          {(items || []).map(item => (
            <tr>
              <td>
                <Link to={`/app/admin/product-membership/${item.ID}`}>
                  {item.ID}
                </Link>
              </td>
              <td>{parseObject(item.CreatedAt)}</td>
              <td>{parseObject(item.UpdatedAt)}</td>
              <td>{parseObject(item.Cause)}</td>
              <td>{parseObject(item.Product)}</td>
            </tr>
          ))}
        </table>
      </Card>
    </Content>
  )
}

ListProductMembership.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListProductMembership
