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

const LIST_COMPONENTTYPEFIELDS = gql`
  {
    listComponentTypeFields(limit: 20) {
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
      FieldType {
        ID
      }
      Weight
      Required
    }
  }
`;

function ListComponentTypeFields({ navigate = () => {} }) {
  const {
    data: { listComponentTypeFields: items = [] },
    error,
    loading,
  } = useQuery(LIST_COMPONENTTYPEFIELDS);

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
          label="Create the first ComponentTypeFields"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    );
  }

  return (
    <Content top>
      <Card width={9 / 10}>
        <h1>List ComponentTypeFields</h1>
        <Button
          label="Create a new ComponentTypeFields"
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
              <TableHeaderCell>FieldType</TableHeaderCell>
              <TableHeaderCell>Weight</TableHeaderCell>
              <TableHeaderCell>Required</TableHeaderCell>
            </TableRow>
          </TableHeader>
          <tbody>
            {(items || []).map(item => (
              <TableRow key={item.ID}>
                <TableCell>
                  <Link
                    to={`/app/admin/component-type-fields/${parseObject(
                      item.ID
                    )}`}
                  >
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
                    to={`/app/admin/field-type/${parseObject(item.FieldType)}`}
                  >
                    {parseObject(item.FieldType)}
                  </Link>
                </TableCell>
                <TableCell>{parseObject(item.Weight)}</TableCell>
                <TableCell>{parseObject(item.Required)}</TableCell>
              </TableRow>
            ))}
          </tbody>
        </DataTable>
      </Card>
    </Content>
  );
}

ListComponentTypeFields.propTypes = {
  navigate: PropTypes.func,
};

export default ListComponentTypeFields;
