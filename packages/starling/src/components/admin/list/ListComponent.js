import React from 'react';
import PropTypes from 'prop-types';
import { useQuery } from 'react-apollo-hooks';
import { Link } from '@reach/router';
import gql from 'graphql-tag';
import Spinner from '@openmob/bluebird/src/components/loaders/Spinner';
import Content from '@openmob/bluebird/src/components/layout/Content';
import Card from '@openmob/bluebird/src/components/cards/Card';
import Button from '@openmob/bluebird/src/components/buttons/Button';
import DataTable from '@openmob/bluebird/src/components/tables/DataTable';
import TableHeader from '@openmob/bluebird/src/components/tables/TableHeader';
import TableHeaderCell from '@openmob/bluebird/src/components/tables/TableHeaderCell';
import TableRow from '@openmob/bluebird/src/components/tables/TableRow';
import TableCell from '@openmob/bluebird/src/components/tables/TableCell';
import parseObject from '../../../common/helpers';

const LIST_COMPONENT = gql`
  {
    listComponent(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      ComponentType {
        ID
      }
      ComponentImplementation {
        ID
      }
      LayoutColumn {
        ID
      }
      Weight
    }
  }
`;

function ListComponent({ navigate = () => {} }) {
  const {
    data: { listComponent: items = [] },
    error,
    loading,
  } = useQuery(LIST_COMPONENT);

  if (loading) {
    return <Spinner />;
  }

  if (error) {
    return <div>Error! {error.message}</div>;
  }

  if (items === null || items.length === 0) {
    return (
      <Content>
        <Button
          label="Create the first Component"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    );
  }

  return (
    <Content top>
      <Card width={9 / 10}>
        <h1>List Component</h1>
        <Button
          label="Create a new Component"
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
              <TableHeaderCell>ComponentType</TableHeaderCell>
              <TableHeaderCell>ComponentImplementation</TableHeaderCell>
              <TableHeaderCell>LayoutColumn</TableHeaderCell>
              <TableHeaderCell>Weight</TableHeaderCell>
            </TableRow>
          </TableHeader>
          <tbody>
            {(items || []).map(item => (
              <TableRow key={item.ID}>
                <TableCell>
                  <Link to={`/app/admin/component/${parseObject(item.ID)}`}>
                    {parseObject(item.ID)}
                  </Link>
                </TableCell>
                <TableCell>{parseObject(item.CreatedAt)}</TableCell>
                <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
                <TableCell>
                  <Link
                    to={`/app/admin/component-type/${parseObject(
                      item.ComponentType
                    )}`}
                  >
                    {parseObject(item.ComponentType)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link
                    to={`/app/admin/component-implementation/${parseObject(
                      item.ComponentImplementation
                    )}`}
                  >
                    {parseObject(item.ComponentImplementation)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link
                    to={`/app/admin/layout-column/${parseObject(
                      item.LayoutColumn
                    )}`}
                  >
                    {parseObject(item.LayoutColumn)}
                  </Link>
                </TableCell>
                <TableCell>{parseObject(item.Weight)}</TableCell>
              </TableRow>
            ))}
          </tbody>
        </DataTable>
      </Card>
    </Content>
  );
}

ListComponent.propTypes = {
  navigate: PropTypes.func,
};

export default ListComponent;
