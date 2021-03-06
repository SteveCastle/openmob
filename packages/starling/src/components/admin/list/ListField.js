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

const LIST_FIELD = gql`
  {
    listField(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      FieldType {
        ID
      }
      StringValue
      IntValue
      FloatValue
      BooleanValue
      DateTimeValue {
        seconds
      }
      DataPath
      Component {
        ID
      }
    }
  }
`;

function ListField({ navigate = () => {} }) {
  const {
    data: { listField: items = [] },
    error,
    loading,
  } = useQuery(LIST_FIELD);

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
          label="Create the first Field"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    );
  }

  return (
    <Content top>
      <Card width={9 / 10}>
        <h1>List Field</h1>
        <Button
          label="Create a new Field"
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
              <TableHeaderCell>FieldType</TableHeaderCell>
              <TableHeaderCell>StringValue</TableHeaderCell>
              <TableHeaderCell>IntValue</TableHeaderCell>
              <TableHeaderCell>FloatValue</TableHeaderCell>
              <TableHeaderCell>BooleanValue</TableHeaderCell>
              <TableHeaderCell>DateTimeValue</TableHeaderCell>
              <TableHeaderCell>DataPath</TableHeaderCell>
              <TableHeaderCell>Component</TableHeaderCell>
            </TableRow>
          </TableHeader>
          <tbody>
            {(items || []).map(item => (
              <TableRow key={item.ID}>
                <TableCell>
                  <Link to={`/app/admin/field/${parseObject(item.ID)}`}>
                    {parseObject(item.ID)}
                  </Link>
                </TableCell>
                <TableCell>{parseObject(item.CreatedAt)}</TableCell>
                <TableCell>{parseObject(item.UpdatedAt)}</TableCell>
                <TableCell>
                  <Link
                    to={`/app/admin/field-type/${parseObject(item.FieldType)}`}
                  >
                    {parseObject(item.FieldType)}
                  </Link>
                </TableCell>
                <TableCell>{parseObject(item.StringValue)}</TableCell>
                <TableCell>{parseObject(item.IntValue)}</TableCell>
                <TableCell>{parseObject(item.FloatValue)}</TableCell>
                <TableCell>{parseObject(item.BooleanValue)}</TableCell>
                <TableCell>{parseObject(item.DateTimeValue)}</TableCell>
                <TableCell>{parseObject(item.DataPath)}</TableCell>
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
  );
}

ListField.propTypes = {
  navigate: PropTypes.func,
};

export default ListField;
