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

const LIST_PHOTO = gql`
  {
    listPhoto(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      URI
      Width
      Height
    }
  }
`;

function ListPhoto({ navigate = () => {} }) {
  const {
    data: { listPhoto: items = [] },
    error,
    loading,
  } = useQuery(LIST_PHOTO);

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
          label="Create the first Photo"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    );
  }

  return (
    <Content top>
      <Card width={9 / 10}>
        <h1>List Photo</h1>
        <Button
          label="Create a new Photo"
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
              <TableHeaderCell>URI</TableHeaderCell>
              <TableHeaderCell>Width</TableHeaderCell>
              <TableHeaderCell>Height</TableHeaderCell>
            </TableRow>
          </TableHeader>
          <tbody>
            {(items || []).map(item => (
              <TableRow key={item.ID}>
                <TableCell>
                  <Link to={`/app/admin/photo/${parseObject(item.ID)}`}>
                    {parseObject(item.ID)}
                  </Link>
                </TableCell>
                <TableCell>{parseObject(item.CreatedAt)}</TableCell>
                <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
                <TableCell>{parseObject(item.URI)}</TableCell>
                <TableCell>{parseObject(item.Width)}</TableCell>
                <TableCell>{parseObject(item.Height)}</TableCell>
              </TableRow>
            ))}
          </tbody>
        </DataTable>
      </Card>
    </Content>
  );
}

ListPhoto.propTypes = {
  navigate: PropTypes.func,
};

export default ListPhoto;
