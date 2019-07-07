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

const LIST_FOLLOWER = gql`
  {
    listFollower(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Contact {
        ID
      }
      Cause {
        ID
      }
    }
  }
`;

function ListFollower({ navigate = () => {} }) {
  const {
    data: { listFollower: items = [] },
    error,
    loading,
  } = useQuery(LIST_FOLLOWER);

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
          label="Create the first Follower"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    );
  }

  return (
    <Content top>
      <Card width={9 / 10}>
        <h1>List Follower</h1>
        <Button
          label="Create a new Follower"
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
              <TableHeaderCell>Contact</TableHeaderCell>
              <TableHeaderCell>Cause</TableHeaderCell>
            </TableRow>
          </TableHeader>
          <tbody>
            {(items || []).map(item => (
              <TableRow key={item.ID}>
                <TableCell>
                  <Link to={`/app/admin/follower/${parseObject(item.ID)}`}>
                    {parseObject(item.ID)}
                  </Link>
                </TableCell>
                <TableCell>{parseObject(item.CreatedAt)}</TableCell>
                <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
                <TableCell>
                  <Link to={`/app/admin/contact/${parseObject(item.Contact)}`}>
                    {parseObject(item.Contact)}
                  </Link>
                </TableCell>
                <TableCell>
                  <Link to={`/app/admin/cause/${parseObject(item.Cause)}`}>
                    {parseObject(item.Cause)}
                  </Link>
                </TableCell>
              </TableRow>
            ))}
          </tbody>
        </DataTable>
      </Card>
    </Content>
  );
}

ListFollower.propTypes = {
  navigate: PropTypes.func,
};

export default ListFollower;
