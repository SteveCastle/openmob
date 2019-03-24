import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_COMPANY = gql`
  {
    listCompany(limit: 20) {
      ID
    }
  }
`

function ListCompany({ navigate }) {
  const {
    data: { listCompany: items = [] },
    error,
    loading,
  } = useQuery(LIST_COMPANY)

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
          label="Create the first Company"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List Company</h1>
      <Button label="Create a new Company" onClick={() => navigate('create')} />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/company/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListCompany.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListCompany
