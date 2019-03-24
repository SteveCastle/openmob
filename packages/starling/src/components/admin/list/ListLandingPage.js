import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Button from '@openmob/bluebird/src/components/buttons/Button'

const LIST_LANDINGPAGE = gql`
  {
    listLandingPage(limit: 20) {
      ID
    }
  }
`

function ListLandingPage({ navigate }) {
  const {
    data: { listLandingPage: items = [] },
    error,
    loading,
  } = useQuery(LIST_LANDINGPAGE)

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
          label="Create the first LandingPage"
          onClick={() => navigate('create')}
        />
      </div>
    )
  }

  return (
    <div>
      <h1>List LandingPage</h1>
      <Button
        label="Create a new LandingPage"
        onClick={() => navigate('create')}
      />
      {(items || []).map(item => (
        <li>
          <Link to={`/app/admin/landingpage/${item.ID}`}>{item.ID}</Link>
        </li>
      ))}
    </div>
  )
}

ListLandingPage.propTypes = {
  navigate: PropTypes.func.isRequired,
}

export default ListLandingPage
