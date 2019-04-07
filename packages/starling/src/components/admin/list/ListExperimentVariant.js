import React from 'react'
import PropTypes from 'prop-types'
import { useQuery } from 'react-apollo-hooks'
import { Link } from '@reach/router'
import gql from 'graphql-tag'
import Content from '@openmob/bluebird/src/components/layout/Content'
import Card from '@openmob/bluebird/src/components/cards/Card'
import Button from '@openmob/bluebird/src/components/buttons/Button'
import DataTable from '@openmob/bluebird/src/components/tables/DataTable'
import TableHeader from '@openmob/bluebird/src/components/tables/TableHeader'
import TableHeaderCell from '@openmob/bluebird/src/components/tables/TableHeaderCell'
import TableRow from '@openmob/bluebird/src/components/tables/TableRow'
import TableCell from '@openmob/bluebird/src/components/tables/TableCell'
import parseObject from '../../../common/helpers'

const LIST_EXPERIMENTVARIANT = gql`
  {
    listExperimentVariant(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Title
      VariantType
      Experiment {
        ID
      }
      LandingPage {
        ID
      }
      Field {
        ID
      }
      Component {
        ID
      }
    }
  }
`

function ListExperimentVariant({ navigate = () => {} }) {
  const {
    data: { listExperimentVariant: items = [] },
    error,
    loading,
  } = useQuery(LIST_EXPERIMENTVARIANT)

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
          label="Create the first ExperimentVariant"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    )
  }

  return (
    <Content>
      <Card>
        <h1>List ExperimentVariant</h1>
        <Button
          label="Create a new ExperimentVariant"
          onClick={() => navigate('create')}
          block
          variant="primary"
        />
        <DataTable>
          <TableHeader>
            <TableRow>
              <TableHeaderCell>ID</TableHeaderCell>
              <TableHeaderCell>CreatedAt</TableHeaderCell>
              <TableHeaderCell>UpdatedAt</TableHeaderCell>
              <TableHeaderCell>Title</TableHeaderCell>
              <TableHeaderCell>VariantType</TableHeaderCell>
              <TableHeaderCell>Experiment</TableHeaderCell>
              <TableHeaderCell>LandingPage</TableHeaderCell>
              <TableHeaderCell>Field</TableHeaderCell>
              <TableHeaderCell>Component</TableHeaderCell>
            </TableRow>
          </TableHeader>
          <tbody>
            {(items || []).map(item => (
              <TableRow key={item.ID}>
                <TableCell>
                  <Link
                    to={`/app/admin/experiment-variant/${parseObject(item.ID)}`}
                  >
                    {parseObject(item.ID)}
                  </Link>
                </TableCell>
                <TableCell>{parseObject(item.CreatedAt)}</TableCell>
                <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
                <TableCell>{parseObject(item.Title)}</TableCell>
                <TableCell>{parseObject(item.VariantType)}</TableCell>
                <TableCell>
                  <Link
                    to={`/app/admin/experiment/${parseObject(item.Experiment)}`}
                  >
                    {parseObject(item.Experiment)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link
                    to={`/app/admin/landing-page/${parseObject(
                      item.LandingPage
                    )}`}
                  >
                    {parseObject(item.LandingPage)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link to={`/app/admin/field/${parseObject(item.Field)}`}>
                    {parseObject(item.Field)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link
                    to={`/app/admin/component/${parseObject(item.Component)}`}
                  >
                    {parseObject(item.Component)}
                  </Link>
                </TableCell>
              </TableRow>
            ))}
          </tbody>
        </DataTable>
      </Card>
    </Content>
  )
}

ListExperimentVariant.propTypes = {
  navigate: PropTypes.func,
}

export default ListExperimentVariant
