import React from 'react';
import PropTypes from 'prop-types';
import { useQuery } from 'react-apollo-hooks';
import gql from 'graphql-tag';
import Content from '@openmob/bluebird/src/components/layout/Content';
import CardGrid from '@openmob/bluebird/src/components/lists/CardGrid';
import CardGridItem from '@openmob/bluebird/src/components/lists/CardGridItem';
import Header from '@openmob/bluebird/src/components/type/Header';
import parseObject from '../../common/helpers';

const LIST_PAGE = gql`
  query LandingPageQuery($id: ID!) {
    getCause(ID: $id) {
      ID
      HomePage {
        ID
        Title
      }
      LandingPages {
        ID
        Title
      }
    }
  }
`;

function MyPages({ navigate = () => {}, causeID }) {
  const {
    data: { getCause: cause = {} },
    error,
    loading,
  } = useQuery(LIST_PAGE, {
    variables: {
      id: causeID,
    },
  });
  console.log('cause', cause);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error! {error.message}</div>;
  }

  return (
    <Content top left>
      <Header>My Pages</Header>
      <CardGrid>
        <CardGridItem
          title={parseObject(cause.HomePage.Title)}
          onClick={() =>
            navigate(
              `/app/cause/${parseObject(cause.ID)}/pages/homepage/${parseObject(
                cause.HomePage.ID
              )}`
            )
          }
        />
        {(cause.LandingPages || []).map(item => (
          <CardGridItem
            title={parseObject(item.Title)}
            description={parseObject(item.Summary)}
            onClick={() => navigate(`/app/cause/${parseObject(item.ID)}`)}
          />
        ))}
        <CardGridItem title="+" onClick={() => navigate('new')} />
      </CardGrid>
    </Content>
  );
}

MyPages.propTypes = {
  navigate: PropTypes.func,
};

export default MyPages;
