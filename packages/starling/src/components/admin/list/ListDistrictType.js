import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_DISTRICTTYPE = gql`
  {
    listDistrictType(limit: 20) {
      ID
    }
  }
`

function ListDistrictType({ navigate }) {
  const {
    data: { listDistrictType: items = [] },
    error,
    loading,
  } = useQuery(LIST_DISTRICTTYPE)

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
          label="Create the first DistrictType"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List DistrictType</h1>
      <Button
        label="Create a new DistrictType"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/districttype/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListDistrictType.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListDistrictType
