import React from 'react';
import PropTypes from 'prop-types';
import { useQuery } from 'react-apollo-hooks';
import { Link } from '@reach/router';
import gql from 'graphql-tag';
import Spinner from '@openmob/bluebird/src/components/loaders/Spinner';

import Content from '@openmob/bluebird/src/components/layout/Content';
import CardGrid from '@openmob/bluebird/src/components/lists/CardGrid';
import CardGridItem from '@openmob/bluebird/src/components/lists/CardGridItem';
import Button from '@openmob/bluebird/src/components/buttons/Button';
import Header from '@openmob/bluebird/src/components/type/Header';
import parseObject from '../../common/helpers';

const LIST_CAUSE = gql`
  {
    listCause(limit: 20) {
      ID
      CreatedAt {
        seconds
      }
      UpdatedAt {
        seconds
      }
      Title
      Slug
      Summary
      Photo {
        URI
      }
    }
  }
`;

function MyCauses({ navigate = () => {} }) {
  const {
    data: { listCause: items = [] },
    error,
    loading,
  } = useQuery(LIST_CAUSE);

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
          label="Create the first Cause"
          onClick={() => navigate('create')}
          variant="primary"
        />
      </Content>
    );
  }

  return (
    <Content top left>
      <Header>My Causes</Header>
      <CardGrid>
        {(items || []).map(item => (
          <CardGridItem
            title={item.Title}
            caption={item.Summary}
            image={item.Photo.URI}
            onClick={() => navigate(`/app/cause/${parseObject(item.ID)}`)}
          />
        ))}
        <CardGridItem title="+" onClick={() => navigate('new')} />
      </CardGrid>
    </Content>
  );
}

MyCauses.propTypes = {
  navigate: PropTypes.func,
};

export default MyCauses;
