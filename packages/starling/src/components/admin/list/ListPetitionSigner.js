import React from 'react';
import PropTypes from 'prop-types';
import { useQuery } from 'react-apollo-hooks';
import { Link } from '@reach/router';
import gql from 'graphql-tag';
import Content from '@openmob/bluebird/src/components/layout/Content';
import Card from '@openmob/bluebird/src/components/cards/Card';
import Button from '@openmob/bluebird/src/components/buttons/Button';
import DataTable from '@openmob/bluebird/src/components/tables/DataTable';
import TableHeader from '@openmob/bluebird/src/components/tables/TableHeader';
import TableHeaderCell from '@openmob/bluebird/src/components/tables/TableHeaderCell';
import TableRow from '@openmob/bluebird/src/components/tables/TableRow';
import TableCell from '@openmob/bluebird/src/components/tables/TableCell';
import parseObject from '../../../common/helpers';

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
`;

function ListPetitionSigner({ navigate = () => {} }) {
  const {
    data: { listPetitionSigner: items = [] },
    error,
    loading,
  } = useQuery(LIST_PETITIONSIGNER);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error! {error.message}</div>;
  }

  if (items === null || items.length === 0) {
    return (
      <Content>
        <Button
          label="Create the first PetitionSigner"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    );
  }

  return (
    <Content>
      <Card>
        <h1>List PetitionSigner</h1>
        <Button
          label="Create a new PetitionSigner"
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
              <TableHeaderCell>Petition</TableHeaderCell>
              <TableHeaderCell>Contact</TableHeaderCell>
              <TableHeaderCell>Cause</TableHeaderCell>
            </TableRow>
          </TableHeader>
          <tbody>
            {(items || []).map(item => (
              <TableRow key={item.ID}>
                <TableCell>
                  <Link
                    to={`/app/admin/petition-signer/${parseObject(item.ID)}`}
                  >
                    {parseObject(item.ID)}
                  </Link>
                </TableCell>
                <TableCell>{parseObject(item.CreatedAt)}</TableCell>
                <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
                <TableCell>
                  <Link
                    to={`/app/admin/petition/${parseObject(item.Petition)}`}
                  >
                    {parseObject(item.Petition)}
                  </Link>
                </TableCell>
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

ListPetitionSigner.propTypes = {
  navigate: PropTypes.func,
};

export default ListPetitionSigner;
